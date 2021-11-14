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
	"github.com/mt1976/templatebuiler/das"
	logs "github.com/mt1976/templatebuiler/logs"
)

type enrichments struct {
	ObjectName         string
	ObjectNameLower    string
	ObjectCamelCase    string
	ObjectGlyph        string
	EndpointRoot       string
	QueryString        string
	QueryField         string
	SourceName         string
	Version            string
	Date               string
	Time               string
	Who                string
	Host               string
	FieldsList         []fields
	UserFrendlyName    string
	SQLTableName       string
	SQLSearchID        string
	SearchKey          string
	SourceType         string
	MessageList        []messages
	Path               string
	ProjectRepo        string
	UUID               string
	Title              string
	PageTitle          string
	UserMenu           string
	MenuHeader         string
	RangeUserMenuStart string
	RangeEnd           string
	MenuHREF           string
	MenuOnClick        string
	MenuGlyph          string
	MenuTextClass      string
	MenuText           string
	ItemsOnPageWc      string
	ItemList           string
	RangeItemList      string
}

type fields struct {
	FieldName     string
	Type          string
	Default       string
	FieldSQL      string
	Formatted     string
	TemplateField string
	Disabled      string
	Hidden        string
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
	logs.Information("User", username())
	logs.Header("Connectivity")
	logs.Information("Default Input", core.ApplicationProperties["static_in"])

	logs.Information("Default Output", core.ApplicationProperties["static_out"])

	logs.Break()
	logs.Header("Searching for work...")
	logs.Break()
	paths := getFiles(core.ApplicationProperties["static_in"])
	logs.Success("Found Files in " + core.ApplicationProperties["static_in"])
	logs.Information("Found ", fmt.Sprintf("%d", len(paths))+" files in "+core.ApplicationProperties["static_in"])
	// loop through files from Paths
	noFiles := len(paths)

	logs.Break()

	for i := 0; i < noFiles; i++ {

		// if last four character in paths[i] are ".cfg" then proceed otherwise skip this item

		fileExtension := paths[i][len(paths[i])-4:]
		//fmt.Println(fileExtension) // gives "lo"

		if fileExtension == ".cfg" {
			logs.Information("Processing", paths[i])
			//	logs.Information("Populate", "Replacement Values")

			props := core.Config_Get(paths[i])
			e := enrichments{ObjectName: props["objectname"]}
			//capitalize first character of enrichment.ObjectName
			e.ObjectName = strings.Title(e.ObjectName)
			e.ObjectCamelCase = strings.ToLower(e.ObjectName[:1]) + e.ObjectName[1:]
			e.ObjectNameLower = strings.ToLower(e.ObjectName)
			e.Version = release
			e.Time = time.Now().Format(core.TIMEFORMATUSER)
			e.Date = time.Now().Format(core.DATEFORMATUSER)
			e.Host = tmpHostname
			e.Who = username()

			e.UserFrendlyName = props["userfrendlyname"]
			if e.UserFrendlyName == "" {
				e.UserFrendlyName = e.ObjectNameLower
			}
			e.SQLTableName = props["sqltablename"]
			e.SQLSearchID = props["sqlsearchid"]
			e.QueryString = props["querystring"]
			e.QueryField = "{{." + props["queryfield"] + "}}"
			if props["endpointroot"] == "" {
				e.EndpointRoot = e.ObjectNameLower
			} else {
				e.EndpointRoot = props["endpointroot"]
			}

			e.Path = pwd
			e.ObjectGlyph = props["objectglyph"]
			e.ProjectRepo = props["projectrepo"] + "/"
			e.UUID = genUUID()

			e.Title = wrap("Title")
			e.PageTitle = wrap("PageTitle")
			e.UserMenu = wrap("UserMenu")
			e.MenuHeader = "{{ (index .UserMenu 0).MenuHeaderText}}"
			e.RangeUserMenuStart = "{{range .UserMenu}}"
			e.RangeEnd = "{{end}}"
			e.MenuHREF = wrap("MenuHREF")
			e.MenuOnClick = wrap("MenuOnClick")
			e.MenuGlyph = wrap("MenuGlyph")
			e.MenuTextClass = wrap("MenuTextClass")
			e.MenuText = wrap("MenuText")
			e.ItemsOnPageWc = wrap("ItemsOnPage")
			e.ItemList = wrap("ItemList")
			e.RangeItemList = "{{range .ItemList}}"

			csvPath := pwd + core.ApplicationProperties["static_in"] + "/" + e.ObjectName + ".csv"

			if props["use"] == "db" {
				// Do nothing for now
				logs.Information("Getting List of fields from DB", props["server"]+" "+props["database"]+" "+props["tablename"])
				e = getFieldsFromDB(e, props)
			} else {
				logs.Information("Getting List of fields from CSV", csvPath)
				e = ReadCsvFile(csvPath, e)
			}

			logs.Break()
			logs.Header("Generating Files")
			logs.Break()

			if strings.ToUpper(props["create_dao"]) == "Y" {
				processTemplate("dao.template", paths[i], "dao", e)
				e.MessageList = append(e.MessageList, messages{Message: "* dao"})
			} else {
				logs.Information("Skipping", "dao")
			}

			if strings.ToUpper(props["create_application"]) == "Y" {
				processTemplate("application.template", paths[i], "application", e)
				e.MessageList = append(e.MessageList, messages{Message: "* application"})
			} else {
				logs.Information("Skipping", "application")
			}

			if strings.ToUpper(props["create_datamodel"]) == "Y" {
				processTemplate("datamodel.template", paths[i], "datamodel", e)
				e.MessageList = append(e.MessageList, messages{Message: "* datamodel"})
			} else {
				logs.Information("Skipping", "datamodel")
			}

			if strings.ToUpper(props["create_job"]) == "Y" {
				processTemplate("job.template", paths[i], "job", e)
				e.MessageList = append(e.MessageList, messages{Message: "* job"})
			} else {
				logs.Information("Skipping", "job")
			}

			if strings.ToUpper(props["create_menu"]) == "Y" {
				processTemplate("menu.template", paths[i], "menu", e)
				e.MessageList = append(e.MessageList, messages{Message: "* menu"})
			} else {
				logs.Information("Skipping", "menu")
			}

			if strings.ToUpper(props["create_html"]) == "Y" {
				processHTMLTemplate("htmlList.template", paths[i], "html", e, "List")
				e.MessageList = append(e.MessageList, messages{Message: "* html - List"})
				processHTMLTemplate("htmlView.template", paths[i], "html", e, "View")
				e.MessageList = append(e.MessageList, messages{Message: "* html - View"})
				processHTMLTemplate("htmlEdit.template", paths[i], "html", e, "Edit")
				e.MessageList = append(e.MessageList, messages{Message: "* html - Edit"})
				processHTMLTemplate("htmlNew.template", paths[i], "html", e, "New")
				e.MessageList = append(e.MessageList, messages{Message: "* html - New"})
			} else {
				logs.Information("Skipping", "html")
			}

			processTemplate("header.nfo", paths[i], "results", e)
			e.MessageList = append(e.MessageList, messages{Message: "* results"})

		}

	}
	logs.Break()
	logs.Success("Templating Complete")
	logs.Break()

}

func wrap(in string) string {
	return "{{." + in + "}}"
}

func addField(en enrichments, fn string, tp string, df string) []fields {

	origfn := fn

	//if first charachter of fieldName is _ then replace _ with SYS

	noinput := ""
	hidden := ""

	if string(fn[0]) == "_" {
		//Convert fn to Title Case
		fn = strings.Replace(fn, "_", "", -1)
		fn = strings.ToUpper(fn[:1]) + fn[1:]
		//fmt.Println(fn)
		fn = "SYS" + fn
		noinput = "disabled=\"disabled\""
		hidden = "hidden"
	}
	fn = strings.ToUpper(fn[:1]) + fn[1:]

	info := fmt.Sprintf("| %-25s | %-10s | %10s | %-25s |", fn, tp, df, origfn)
	tplField := "{{." + fn + "}}"
	en.FieldsList = append(en.FieldsList, fields{FieldName: fn,
		Type:          tp,
		Default:       df,
		FieldSQL:      origfn,
		Formatted:     info,
		TemplateField: tplField,
		Disabled:      noinput,
		Hidden:        hidden})
	logs.Information(info, "")

	return en.FieldsList
}

//Get list of files from a folder
func getFiles(dir string) []string {
	logs.Information("Searching...", "")
	pwd, _ := os.Getwd()
	dir = pwd + dir + "/"
	//logs.Information("In Queue Path", dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var paths []string
	for _, file := range files {
		if !file.IsDir() {
			paths = append(paths, filepath.Join(dir, file.Name()))
			logs.Information("Found File", file.Name())
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
	logs.Information("Processing Template", w)

	//spew.Dump(replacements)
	fp := e.Path + "/templates/" + w

	t, err := template.ParseFiles(fp)
	if err != nil {
		logs.Error("Load Template", err)
	}

	f, err := os.Create(e.Path + "/" + core.ApplicationProperties["static_out"] + "/" + c + "/" + e.ObjectCamelCase + ".go_tmp")
	if err != nil {
		logs.Error("Create file: ", err)
		return
	}

	err2 := t.Execute(f, e)
	if err2 != nil {
		logs.Error("Process Template", err2)
	}
	f.Close()
	logs.Success(f.Name() + " Generated")
}

func processHTMLTemplate(w string, p string, c string, e enrichments, pt string) {
	logs.Information("Processing Template", w)

	//spew.Dump(replacements)
	fp := e.Path + "/templates/" + w

	t, err := template.ParseFiles(fp)
	if err != nil {
		logs.Error("Load Template", err)
	}

	f, err := os.Create(e.Path + "/" + core.ApplicationProperties["static_out"] + "/" + c + "/" + e.ObjectName + "_" + pt + ".html")
	if err != nil {
		logs.Error("Create file: ", err)
		return
	}

	err2 := t.Execute(f, e)
	if err2 != nil {
		logs.Error("Process Template", err2)
	}
	f.Close()
	logs.Success(f.Name() + " Generated")
}

func tableHeader() {
	logs.Break()
	logs.Header("Table Information")
	logs.Break()
	info := fmt.Sprintf("| %-25s | %-10s | %-10s | %-25s |", "Field Name", "Type", "Default", "SQL Field Name")
	logs.Information(info, "")
	logs.Break()
}

func ReadCsvFile(filePath string, e enrichments) enrichments {
	// Load a csv file.
	//logs.Information("Read CSV", filePath)
	f, _ := os.Open(filePath)
	//logs.Information("File Open", filePath)
	// Create a new reader.
	r := csv.NewReader(f)
	//logs.Information("New Reader", filePath)
	tableHeader()
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

		e.FieldsList = addField(e, record[0], record[1], record[2])

	}
	logs.Break()
	return e
}

func getFieldsFromDB(e enrichments, p map[string]string) enrichments {
	// Open Database Connection
	db, err := core.GlobalsDatabaseConnect(p)
	if err != nil {
		logs.Error("Database Connection", err)
	}
	//fmt.Printf("db: %v\n", db)

	tsql := fmt.Sprintf("USE %s EXEC sp_columns '%s'", p["database"], p["sqltablename"])

	logs.Information("SQL", tsql)
	results, noFields, err := das.Query(db, tsql)
	//fmt.Printf("results: %v\n", results)
	//fmt.Printf("noFields: %v\n", noFields)
	//spew.Dump(results)
	if noFields == 0 {
		logs.Error("No Fields Found", err)
	}
	tableHeader()
	for _, row := range results {
		colName := row["COLUMN_NAME"].(string)
		colType := row["TYPE_NAME"].(string)
		//fmt.Printf("colName: %v %v\n", colName, colType)
		colDefault := ""
		switch colType {
		case "varchar", "nvarchar", "char", "nchar", "text", "ntext":
			colDefault = ""
			colType = "String"
		case "int", "bigint", "smallint", "tinyint", "bit":
			colDefault = "0"
			colType = "Int"
		case "decimal", "numeric":
			colDefault = "0.00"
			colType = "Float"
		case "datetime", "smalldatetime", "date", "time", "datetime2", "datetimeoffset":
			colDefault = ""
			colType = "Time"
		case "float", "real", "money", "smallmoney":
			colDefault = "0.00"
			colType = "Float"
		case "int identity":
			colDefault = "0"
			colType = "Int"
		default:
			colType = "String"
			colDefault = ""
		}
		e.FieldsList = addField(e, colName, colType, colDefault)
	}
	return e
}
