package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"

	core "github.com/mt1976/templatebuiler/core"
	logs "github.com/mt1976/templatebuiler/logs"
)

type enrichments struct {
	ObjectName      string
	ObjectNameLower string
	ObjectGlyph     string
	EndpointRoot    string
	QueryString     string
	SourceName      string
	Version         string
	Date            string
	Time            string
	Who             string
	Host            string
	FieldsList      []fields
	UserFrendlyName string
	SQLTableName    string
	SQLTableID      string
	SearchKey       string
	SourceType      string
	MessageList     []messages
	Path            string
	ProjectRepo     string
}

type fields struct {
	FieldName string
	Type      string
	Default   string
	FieldSQL  string
}

type messages struct {
	Message string
}

func main() {

	tmpHostname, _ := os.Hostname()
	logs.Break()
	logs.Header("Template Generator")
	logs.Break()

	logs.Information("Initialising...", "")

	core.Initialise()

	logs.Success("Initialised")
	logs.Break()

	logs.Header("Application Information")
	logs.Break()

	logs.Header("Application")
	logs.Information("Name", core.ApplicationProperties["appname"])
	logs.Information("Host Name", tmpHostname)
	release := fmt.Sprintf("%s [r%s-%s]", core.ApplicationProperties["releaseid"], core.ApplicationProperties["releaselevel"], core.ApplicationProperties["releasenumber"])
	logs.Information("Server Release", release)
	logs.Information("Server Date", time.Now().Format(core.DATEFORMATUSER))

	logs.Information("Licence", core.ApplicationProperties["licname"])
	logs.Information("Lic URL", core.ApplicationProperties["liclink"])
	logs.Header("Runtime")
	logs.Information("GO Version", runtime.Version())
	logs.Information("Operating System", runtime.GOOS)
	pwd, _ := os.Getwd()
	logs.Information("Working Directory", pwd)
	logs.Header("Application Database (MSSQL)")
	logs.Information("Server", core.ApplicationPropertiesDB["server"])
	logs.Information("Database", core.ApplicationPropertiesDB["database"])
	logs.Information("Schema", core.ApplicationPropertiesDB["schema"])
	logs.Information("Parent Schema", core.ApplicationPropertiesDB["parentschema"])

	logs.Header("Siena Database (MSSQL)")
	logs.Information("Server", core.SienaPropertiesDB["server"])
	logs.Information("Database", core.SienaPropertiesDB["database"])
	logs.Information("Schema", core.SienaPropertiesDB["schema"])
	logs.Information("Parent Schema", core.SienaPropertiesDB["parentschema"])

	logs.Header("Connectivity")
	logs.Information("Build Queue", core.SienaProperties["static_in"])
	logs.Information("Artifacts", core.SienaProperties["static_out"])

	//scheduler.RunJobLSE("")
	//scheduler.RunJobFII("")

	logs.Header("READY STEADY GO!!!")
	logs.Information("Initialisation", "Vrooom, Vrooooom, Vroooooooo..."+logs.Bike+logs.Bike+logs.Bike+logs.Bike)
	logs.Break()
	paths := getFiles(core.SienaProperties["static_in"])
	logs.Success("Found Files in " + core.SienaProperties["static_in"])
	logs.Information("Found ", fmt.Sprintf("%d", len(paths))+" files in "+core.SienaProperties["static_in"])
	// loop through files from Paths
	noFiles := len(paths)

	logs.Break()
	logs.Header("Loading Templates")

	// applicationTemplate := loadTemplate("application.template")
	// daoTemplate := loadTemplate("dao.template")
	// datamodelTemplate := loadTemplate("datamodel.template")
	// jobTemplate := loadTemplate("job.template")
	// menuTemplate := loadTemplate("menu.template")
	// headerInfoTemplate := loadTemplate("header.nfo")

	// logs.Information("Template", applicationTemplate)
	// logs.Information("Template", daoTemplate)
	// logs.Information("Template", datamodelTemplate)
	// logs.Information("Template", jobTemplate)
	// logs.Information("Template", menuTemplate)
	// logs.Information("Template", headerInfoTemplate)
	logs.Break()
	//clear terminal screen
	//	logs.Clear()

	///

	for i := 0; i < noFiles; i++ {

		// if last four character in paths[i] are ".cfg" then proceed otherwise skip this item

		fileExtension := paths[i][len(paths[i])-4:]
		fmt.Println(fileExtension) // gives "lo"

		if fileExtension == ".cfg" {
			logs.Information("Processing", paths[i])
			logs.Information("Populate", "Replacement Values")

			props := core.Config_Get(paths[i])
			enrichment := enrichments{ObjectName: props["objectname"]}
			//capitalize first character of enrichment.ObjectName
			enrichment.ObjectName = strings.Title(enrichment.ObjectName)

			enrichment.ObjectNameLower = strings.ToLower(enrichment.ObjectName)
			enrichment.Version = release
			enrichment.Time = time.Now().Format(core.TIMEFORMATUSER)
			enrichment.Date = time.Now().Format(core.DATEFORMATUSER)
			enrichment.Host = tmpHostname
			enrichment.Who = username()
			enrichment.UserFrendlyName = props["userfrendlyname"]
			enrichment.SQLTableName = props["tablename"]
			enrichment.SQLTableID = props["sqlid"]
			enrichment.QueryString = props["querystring"]
			//enrichment.SearchKey = props["sqlid"]
			enrichment.SourceType = props["whichdb"]
			enrichment.Path = pwd
			enrichment.ObjectGlyph = props["objectglyph"]
			enrichment.ProjectRepo = props["projectrepo"] + "/"

			csvPath := pwd + "/data/in/" + enrichment.ObjectName + ".csv"

			enrichment = ReadCsvFile(csvPath, enrichment)

			processTemplate("dao.template", paths[i], "dao", enrichment)
			enrichment.MessageList = append(enrichment.MessageList, messages{Message: "* dao"})

			processTemplate("application.template", paths[i], "application", enrichment)
			enrichment.MessageList = append(enrichment.MessageList, messages{Message: "* application"})

			processTemplate("datamodel.template", paths[i], "datamodel", enrichment)
			enrichment.MessageList = append(enrichment.MessageList, messages{Message: "* datamodel"})

			processTemplate("job.template", paths[i], "job", enrichment)
			enrichment.MessageList = append(enrichment.MessageList, messages{Message: "* job"})

			processTemplate("menu.template", paths[i], "menu", enrichment)
			enrichment.MessageList = append(enrichment.MessageList, messages{Message: "* menu"})

			processTemplate("header.nfo", paths[i], "results", enrichment)
			enrichment.MessageList = append(enrichment.MessageList, messages{Message: "* results"})

		}

	}
	logs.Success("Templating Complete")
	logs.Break()

}

func addField(en enrichments, fn string, tp string, df string) []fields {

	origfn := fn

	//if first charachter of fieldName is _ then replace _ with SYS
	if string(fn[0]) == "_" {
		fn = "SYS" + fn
		fn = strings.Replace(fn, "_", "", -1)
	}

	en.FieldsList = append(en.FieldsList, fields{FieldName: fn, Type: tp, Default: df, FieldSQL: origfn})
	info := fmt.Sprintf("Added Field: %s %s %s %s", fn, tp, df, origfn)
	logs.Information(info, "")
	return en.FieldsList
}

//Get list of files from a folder
func getFiles(dir string) []string {

	pwd, _ := os.Getwd()
	dir = pwd + dir + "/"
	logs.Information("In Queue Path", dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var paths []string
	for _, file := range files {
		if !file.IsDir() {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}
	return paths
}

func username() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.Username
}

func genUUID() string {
	id := uuid.New()
	//fmt.Printf("github.com/google/uuid:         %s\n", id.String())
	return id.String()
}

func processTemplate(w string, p string, c string, e enrichments) {
	logs.Information("Processing File", w)

	//spew.Dump(replacements)
	fp := e.Path + "/templates/" + w
	logs.Information("Template File", fp)

	t, err := template.ParseFiles(fp)
	if err != nil {
		logs.Error("Load Template", err)
	}

	f, err := os.Create(e.Path + "/" + core.SienaProperties["static_out"] + "/" + c + "/" + e.ObjectNameLower + genUUID() + ".gotmp")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err2 := t.Execute(f, e)
	if err2 != nil {
		logs.Error("Process Template", err2)
	}
	f.Close()

}

func ReadCsvFile(filePath string, e enrichments) enrichments {
	// Load a csv file.
	logs.Information("Read CSV", filePath)
	f, _ := os.Open(filePath)
	//logs.Information("File Open", filePath)
	// Create a new reader.
	r := csv.NewReader(f)
	//logs.Information("New Reader", filePath)
	for {
		record, err := r.Read()
		//fmt.Printf("record: %v\n", record)
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			logs.Fatal("Read CSV", err)
			panic(err)
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		// fmt.Println(record)
		// fmt.Println(len(record))
		// for value := range record {
		// 	fmt.Printf("  %v\n", record[value])
		// }
		e.FieldsList = addField(e, record[0], record[1], record[2])

	}
	return e
}
