package core

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	logs "github.com/mt1976/mwt-goToolkit/logs"
)

func extractCreate(in string) string {
	//result := in
	findCREATE := "CREATE"
	selectPos := strings.Index(in, findCREATE)
	//log.Println(findCREATE, selectPos)
	newString := in[selectPos:]
	findGO := "GO"
	goPos := strings.Index(newString, findGO)
	outString := newString[0:goPos]
	//log.Println(outString)
	return outString
}

func PokeDatabase(DB *sql.DB) error {
	errordb := DB.Ping()
	if errordb != nil {
		//log.Println(errordb.Error())
		logs.Error("Poke Error", errordb)
	}
	return errordb
}

func connect(mssqlConfig map[string]string) (*sql.DB, error) {

	server := mssqlConfig["server"]
	user := mssqlConfig["user"]
	password := mssqlConfig["password"]
	port := mssqlConfig["port"]
	database := mssqlConfig["database"]
	//	instance := mssqlConfig["instance"]
	if database != "master" {
		//log.Println("Information   : Attemping connection to " + server + " " + database)
		logs.Message("Connecting", "Attemping connection to "+server+" "+database)
	}
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", server, user, password, port, database)
	//connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;", server, user, password, port)

	//log.Println("Connection    : " + connString)
	dbInstance, err := sql.Open("mssql", connString)
	if err != nil {
		//log.Fatal("ERROR!        : Connection attempt with "+database+connString+" failed:", err.Error())
		logs.Error("Connection attempt with "+database+connString+" failed:", err)
	}
	if database != "master" {
		logs.Success("Connected to " + server + " " + database)
		//	log.Println("Information   : Connected to " + server + " " + database)
	}
	keepalive, _ := time.ParseDuration("-1h")
	dbInstance.SetConnMaxLifetime(keepalive)
	//log.Printf("Test Connection %d %d", dbInstance.Stats().OpenConnections, dbInstance.Stats().Idle)
	stmt, err := dbInstance.Prepare("select @@version")
	if err != nil {
		logs.Error("Connection attempt with "+database+connString+" failed:", err)
	}
	row := stmt.QueryRow()
	var result string
	err = row.Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	return dbInstance, err
}

// DataStoreConnect connects application to its datastore database
func GlobalsDatabaseConnect(mssqlConfig map[string]string) (*sql.DB, error) {
	// Connect to SQL Server DB
	//mssqlConfig := getProperties(config)

	var returnDB *sql.DB
	database := mssqlConfig["database"]
	instance := mssqlConfig["instance"]

	logs.Message("Connecting", mssqlConfig["server"]+" "+database+" "+instance)

	mssqlConfig["database"] = "master"

	dbInstance, errConnect := connect(mssqlConfig)
	if errConnect != nil {
		logs.Information("Connection attempt with "+database+" failed:", errConnect.Error())
		//log.Panic(errConnect.Error())
	}

	mssqlConfig["database"] = database

	dbName := database
	if len(instance) != 0 {
		dbName = database + "-" + instance
	}

	checkDBstmt := "SELECT create_date FROM sys.databases WHERE name = '" + dbName + "'"

	stmt2, err2 := dbInstance.Prepare(checkDBstmt)
	//log.Println(checkDBstmt)
	//spew.Dump(stmt2)
	if err2 != nil {
		log.Panic(err2.Error())
	}
	dbCheck := stmt2.QueryRow()
	var result2 string

	err2 = dbCheck.Scan(&result2)
	if err2 != nil {

		logs.Warning("Database " + dbName + " does not exists. GENERATING")
		CreateDatabase(dbInstance, PropertiesDB, dbName)

		PropertiesDB["database"] = dbName
		//Connect to the specific instance
		newDBInstance, errReCon := connect(PropertiesDB)
		if errReCon != nil {
			log.Panicln(errReCon.Error())
		}

		if len(instance) != 0 {
			PropertiesDB["database"] = database + "-" + instance
		}
		CreateDatabaseObjects(newDBInstance, PropertiesDB, "/config/database/appdb/tables", true)
		CreateDatabaseObjects(newDBInstance, PropertiesDB, "/config/database/appdb/views", true)
		returnDB = newDBInstance
	} else {
		logs.Success("Database " + dbName + " exists Created: " + result2)
		if len(mssqlConfig["instance"]) != 0 {
			mssqlConfig["database"] = database + "-" + mssqlConfig["instance"]
		}
		newDBInstance, errReCon := connect(mssqlConfig)
		if errReCon != nil {
			log.Panicln(errReCon.Error())
		}
		returnDB = newDBInstance
	}
	//log.Printf("%d %s", returnDB.Stats().OpenConnections, mssqlConfig["database"])
	//fmt.Printf("%s\n", result)
	logs.Success("Connected to " + mssqlConfig["server"] + " " + mssqlConfig["database"])
	return returnDB, errConnect
}

func CreateDatabase(dbInstance *sql.DB, mssqlConfig map[string]string, dbName string) {

	createDBSQL := "CREATE DATABASE [" + dbName + "]"
	//log.Println(createDBSQL)
	_, errCreateDBSQL := dbInstance.Exec(createDBSQL)
	if errCreateDBSQL != nil {
		log.Panic(errCreateDBSQL)
	}

	log.Println("Information   : Database " + dbName + " created " + Tick)

}

func GlobalsDatabasePoke(dbInstance *sql.DB, mssqlConfig map[string]string) *sql.DB {
	logs.Poke(fmt.Sprintf("Server '%s' Database '%s' Schema '%s'", mssqlConfig["server"], mssqlConfig["database"], mssqlConfig["schema"]), "")
	err := dbInstance.Ping()
	if err != nil {
		logs.Error(fmt.Sprintf("Reconnecting  : Server '%s' Database '%s' Schema '%s'", mssqlConfig["server"], mssqlConfig["database"], mssqlConfig["schema"]), err)
		// Try to reconnect
		dbInstance, err = GlobalsDatabaseConnect(mssqlConfig)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return dbInstance
}

func CreateDatabaseObjects(DB *sql.DB, dbConfig map[string]string, sourcePath string, initialiseDB bool) {
	log.Println("Information   : Creating Database Objects in " + dbConfig["database"] + " from " + sourcePath)

	errdb := DB.Ping()
	if errdb != nil {
		logs.Warning("Unable to Find DB")
	}
	//spew.Dump(DB)
	//spew.Dump(DB.Stats())
	// Get a full list of all views
	_, requiredViews, _ := GetDataList(sourcePath)
	createTemplate, err := ReadDataFile("templateCreate.sql", "/config/database/templates")
	if err != nil {
		log.Println(err.Error())
	}
	dropTemplate, err := ReadDataFile("templateDrop.sql", "/config/database/templates")
	if err != nil {
		log.Println(err.Error())
	}

	schemaTemplate, err := ReadDataFile("templateCreateSchema.sql", "/config/database/templates")
	if err != nil {
		log.Println(err.Error())
	}

	schemaName := dbConfig["schema"]
	databaseName := dbConfig["database"]
	parentschema := dbConfig["parentschema"]

	schemaTemplate = ReplaceWildcard(schemaTemplate, "!SQL.SCHEMA", schemaName)

	//	fmt.Println("***************************************")
	//	log.Println(schemaTemplate)
	//	fmt.Println("***************************************")
	_, err4 := DB.Exec(schemaTemplate)
	if err4 != nil {
		log.Println(err.Error())
	}

	for _, view := range requiredViews {
		thisDrop := dropTemplate
		thisCreate := createTemplate

		fileID := view
		objectName := strings.ReplaceAll(fileID, ".sql", "")
		//	fmt.Println("view name:", objectName)

		thisDrop = ReplaceWildcard(thisDrop, "!SQL.DB", databaseName)
		thisDrop = ReplaceWildcard(thisDrop, "!SQL.SCHEMA", schemaName)
		thisDrop = ReplaceWildcard(thisDrop, "!SQL.VIEW", objectName)

		sqlBody, err := ReadDataFile(fileID, sourcePath)
		if err != nil {
			log.Println(err.Error())
		}

		sqlBody = extractCreate(sqlBody)

		//log.Println("***", sqlBody, "***")

		thisCreate = ReplaceWildcard(thisCreate, "!SQL.DB", databaseName)
		thisCreate = ReplaceWildcard(thisCreate, "!SQL.SCHEMA", schemaName)
		thisCreate = ReplaceWildcard(thisCreate, "!SQL.VIEW", objectName)
		thisCreate = ReplaceWildcard(thisCreate, "!SQL.SOURCE", parentschema)

		sqlBody = ReplaceWildcard(sqlBody, "!SQL.DB", databaseName)
		sqlBody = ReplaceWildcard(sqlBody, "!SQL.SCHEMA", schemaName)
		sqlBody = ReplaceWildcard(sqlBody, "!SQL.VIEW", objectName)
		sqlBody = ReplaceWildcard(sqlBody, "!SQL.SOURCE", parentschema)

		thisCreate = ReplaceWildcard(thisCreate, "!SQL.BODY", sqlBody)
		//fmt.Println("index:", i, fileID, thisDrop, thisCreate)
		//	fmt.Println("***************************************")
		//	fmt.Println(fileID, objectName)
		//	fmt.Println("***************************************")
		//	fmt.Println(thisDrop)
		//fmt.Println("***************************************")

		log.Println("Generating    :", objectName)
		log.Println(sqlBody)
		if !initialiseDB {
			//log.Println(">>>", thisDrop, "<<<")
			_, erra := DB.Exec(thisDrop)
			//spew.Dump(info)
			//spew.Dump(erra)
			if erra != nil {
				log.Panic("ARSE", err.Error())
			}
		}

		//fmt.Println("***************************************")
		//fmt.Println(info)
		//fmt.Println("***************************************")
		//fmt.Println(sqlBody)
		//fmt.Println("***************************************")
		//fmt.Println(thisCreate)
		//fmt.Println("***************************************")

		//spew.Dump(thisCreate)
		//log.Println(">>>", sqlBody, "<<<")
		_, err2 := DB.Exec(sqlBody)

		if err2 != nil {
			log.Println(err2.Error())
		} else {
			log.Println("Generated     :", objectName, Tick)
		}
		//spew.Dump(err2)

		//log.Println("***************************************")
	}
}
