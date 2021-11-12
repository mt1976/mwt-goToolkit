package main

import (
	"fmt"
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

	core "github.com/mt1976/templatebuiler/core"
	logs "github.com/mt1976/templatebuiler/logs"
)

type replacements struct {
	ObjectName      string
	ObjectNameLower string
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
	TableName       string
	SQLID           string
	SearchKey       string
	SourceType      string
}

type fields struct {
	FieldName string
	Type      string
	Default   string
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
	logs.Information("Populate", "Replacement Values")
	//clear terminal screen
	logs.Clear()
	for i := 0; i < noFiles; i++ {
		logs.Information("Processing File", paths[i])
		//testit(paths[i])
		props := core.Config_Get(paths[i])
		replacements := replacements{ObjectName: props["objectname"]}
		replacements.ObjectNameLower = strings.ToLower(replacements.ObjectName)
		replacements.Version = release
		replacements.Time = time.Now().Format(core.TIMEFORMATUSER)
		replacements.Date = time.Now().Format(core.DATEFORMATUSER)
		replacements.Host = tmpHostname
		replacements.Who = username()
		replacements.UserFrendlyName = props["userfrendlyname"]
		replacements.TableName = props["tablename"]
		replacements.SQLID = props["sqlid"]
		replacements.QueryString = props["querystring"]
		replacements.SearchKey = props["sqlid"]
		replacements.SourceType = props["whichdb"]

		replacements.FieldsList = append(replacements.FieldsList, fields{FieldName: "Field1", Type: "String"})
		replacements.FieldsList = append(replacements.FieldsList, fields{FieldName: "Field2", Type: "Float"})
		replacements.FieldsList = append(replacements.FieldsList, fields{FieldName: "ID", Type: "Int"})
		replacements.FieldsList = append(replacements.FieldsList, fields{FieldName: "Code", Type: "String"})
		replacements.FieldsList = append(replacements.FieldsList, fields{FieldName: "Name", Type: "Time"})
		replacements.FieldsList = append(replacements.FieldsList, fields{FieldName: "_who", Type: "String"})

		//spew.Dump(replacements)
		fp := pwd + "/templates/" + "dao.template"
		logs.Information("Template File", fp)

		t, err := template.ParseFiles(fp)
		if err != nil {
		}

		f, err := os.Create(paths[i] + "output" + "dao")
		if err != nil {
			log.Println("create file: ", err)
			return
		}

		err2 := t.Execute(f, replacements)
		if err2 != nil {
		}
		f.Close()

		logs.Break()
		logs.Break()

	}
	logs.Success("Templates Loaded")
	logs.Break()

}

func loadTemplate(name string) string {
	applicationTemplate, _ := core.ReadDataFile(name, "/templates")
	logs.Success("Template '" + name + "' Loaded")
	return applicationTemplate
}

func testit(path string) {
	logs.Information("Processing File", path)
	//read the file
	file, err := os.Open(path)
	if err != nil {
		logs.Fatal("Can't Open File", err)
	}
	defer file.Close()
	//read the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		logs.Fatal("Cant Get Data", err)
	}
	//print the file
	logs.Information("File Contents", string(data))

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
