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
	FriendlyName       string
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
	CanView            bool
	CanEdit            bool
	CanSave            bool
	CanNew             bool
	CanDelete          bool
	CanList            bool
	PropertiesName     string
	UsesAdaptor        bool
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

const (
	go_template   = ".go_template"
	html_template = ".html_template"
	json_template = ".json_template"
	nfo_template  = ".nfo_template"
)

func main() {

	logs.Break()
	logs.Header("Template Generator")
	logs.Break()

	core.Initialise()

	headerBumpf()

	logs.Break()
	logs.Header("Searching for work...")
	logs.Break()

	paths := getFiles(data_in())

	noFiles := len(paths)

	logs.Success("Found Files in " + data_in())
	logs.Information("Found ", fmt.Sprintf("%d %s %s", noFiles, " files in ", data_in()))

	// loop through files from Paths

	logs.Break()

	for i := 0; i < noFiles; i++ {
		// if last four character in paths[i] are ".cfg" then proceed otherwise skip this item
		fileExtension := paths[i][len(paths[i])-4:]
		//fmt.Println(fileExtension) // gives "lo"
		if fileExtension == ".cfg" {
			processConfigFile(paths[i])
		}
	}
	logs.Break()
	logs.Success("Templating Complete")
	logs.Break()
}

func processConfigFile(configFile string) {
	logs.Information("Processing", configFile)
	//	logs.Information("Populate", "Replacement Values")

	props := core.Config_Get(configFile)
	e := setupEnrichment(props)

	csvPath := getPWD() + data_in() + "/" + e.ObjectName + ".csv"

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
		processTemplate("dao"+go_template, configFile, "dao", e)
		e.MessageList = append(e.MessageList, messages{Message: "* dao"})
	} else {
		logs.Information("Skipping", "dao")
	}

	if strings.ToUpper(props["create_application"]) == "Y" {
		processTemplate("application"+go_template, configFile, "application", e)
		e.MessageList = append(e.MessageList, messages{Message: "* application"})
	} else {
		logs.Information("Skipping", "application")
	}

	if strings.ToUpper(props["create_datamodel"]) == "Y" {
		processTemplate("datamodel"+go_template, configFile, "datamodel", e)
		e.MessageList = append(e.MessageList, messages{Message: "* datamodel"})
	} else {
		logs.Information("Skipping", "datamodel")
	}

	if strings.ToUpper(props["create_job"]) == "Y" {
		processTemplate("job"+go_template, configFile, "jobs", e)
		e.MessageList = append(e.MessageList, messages{Message: "* job"})
	} else {
		logs.Information("Skipping", "job")
	}

	if strings.ToUpper(props["create_menu"]) == "Y" {
		processTemplate("menu"+json_template, configFile, "design/menu", e)
		e.MessageList = append(e.MessageList, messages{Message: "* menu"})
	} else {
		logs.Information("Skipping", "menu")
	}

	if strings.ToUpper(props["create_html"]) == "Y" {
		if e.CanList {
			processHTMLTemplate("list", configFile, "html", e)
			e.MessageList = append(e.MessageList, messages{Message: "* html -> list"})
		}
		if e.CanView {
			processHTMLTemplate("view", configFile, "html", e)
			e.MessageList = append(e.MessageList, messages{Message: "* html -> view"})
		}
		if e.CanEdit {
			processHTMLTemplate("edit", configFile, "html", e)
			e.MessageList = append(e.MessageList, messages{Message: "* html -> edit"})
		}
		if e.CanNew {
			processHTMLTemplate("new", configFile, "html", e)
			e.MessageList = append(e.MessageList, messages{Message: "* html -> new"})
		}

	} else {
		logs.Information("Skipping", "html")
	}

	processTemplate("catalog"+nfo_template, configFile, "design/catalog", e)
	e.MessageList = append(e.MessageList, messages{Message: "* catalog"})
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
	dir = getPWD() + dir + "/"
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

func getUsername() string {
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

func processTemplate(w string, p string, destFolder string, e enrichments) {
	logs.Information("Processing Template", w)

	extn := ".go_tmp"
	if core.ApplicationProperties["deliverto"] != "" {
		extn = ".go"
	}

	if destFolder == "design/catalog" {
		extn = ".nfo"
	}

	if destFolder == "design/menu" {
		extn = ".json"
	}

	//spew.Dump(replacements)
	fp := e.Path + "/templates/" + w

	t, err := template.ParseFiles(fp)
	if err != nil {
		logs.Error("Load Template", err)
	}

	f, err := os.Create(data_out() + "/" + destFolder + "/" + e.ObjectCamelCase + extn)
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

func processHTMLTemplate(w string, p string, destFolder string, e enrichments) {
	logs.Information("Processing Template", w+html_template)

	fileNamePrefix := strings.ToUpper(w[:1]) + w[1:]

	//spew.Dump(replacements)
	fp := e.Path + "/templates/" + w + html_template

	t, err := template.ParseFiles(fp)
	if err != nil {
		logs.Error("Load Template", err)
	}

	f, err := os.Create(data_out() + "/" + destFolder + "/" + e.ObjectName + "_" + fileNamePrefix + ".html")
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

func genReleaseName() string {
	return fmt.Sprintf("%s [r%s-%s]",
		core.ApplicationProperties["releaseid"],
		core.ApplicationProperties["releaselevel"],
		core.ApplicationProperties["releasenumber"])
}

func getHostName() string {
	host, _ := os.Hostname()
	return host
}

func getPWD() string {
	thisPwd, _ := os.Getwd()
	return thisPwd
}

func data_out() string {
	do := ""
	if core.ApplicationProperties["deliverto"] != "" {
		do = core.ApplicationProperties["deliverto"]
	} else {
		do = getPWD() + core.ApplicationProperties["data_out"]
	}
	return do
}

func data_in() string {
	return core.ApplicationProperties["data_in"]
}

func headerBumpf() {

	logs.Break()

	logs.Header("Application Information")
	logs.Break()

	logs.Header("Application")
	logs.Information("Name", core.ApplicationProperties["appname"])
	logs.Information("Host Name", getHostName())

	logs.Information("Server Release", genReleaseName())
	logs.Information("Server Date", time.Now().Format(core.DATEFORMATUSER))

	logs.Information("Licence", core.ApplicationProperties["licname"])
	logs.Information("Lic URL", core.ApplicationProperties["liclink"])
	logs.Header("Runtime")
	logs.Information("GO Version", runtime.Version())
	logs.Information("Operating System", runtime.GOOS)

	logs.Information("Working Directory", getPWD())
	logs.Information("User", getUsername())
	logs.Header("Connectivity")
	logs.Information("Default Input", data_in())
	logs.Information("Default Output", data_out())
}

func setupEnrichment(props map[string]string) enrichments {
	e := enrichments{ObjectName: props["objectname"]}
	//capitalize first character of enrichment.ObjectName
	e.ObjectName = strings.Title(e.ObjectName)
	e.ObjectCamelCase = strings.ToLower(e.ObjectName[:1]) + e.ObjectName[1:]
	e.ObjectNameLower = strings.ToLower(e.ObjectName)
	e.Version = genReleaseName()
	e.Time = time.Now().Format(core.TIMEFORMATUSER)
	e.Date = time.Now().Format(core.DATEFORMATUSER)
	e.Host = getHostName()
	e.Who = getUsername()

	e.FriendlyName = props["friendlyname"]
	if e.FriendlyName == "" {
		e.FriendlyName = e.ObjectCamelCase
	}
	e.SQLTableName = props["sqltablename"]
	e.SQLSearchID = strings.TrimSpace(props["sqlsearchid"])
	e.QueryString = props["querystring"]
	e.QueryField = "{{." + props["queryfield"] + "}}"
	if props["endpointroot"] == "" {
		e.EndpointRoot = e.ObjectNameLower
	} else {
		e.EndpointRoot = props["endpointroot"]
	}

	e.Path = getPWD()
	e.ObjectGlyph = props["objectglyph"]
	e.ProjectRepo = props["projectrepo"] + "/"
	e.UUID = genUUID()

	e.PropertiesName = ""
	e.UsesAdaptor = false
	if props["propertiesoverride"] == "" {
		e.PropertiesName = "Application"
	} else {
		e.PropertiesName = props["propertiesoverride"]
		e.UsesAdaptor = true
	}

	e = setupTemplateEnrichment(e, props)

	e = setupPermissions(e, props)

	return e
}

func setupTemplateEnrichment(e enrichments, props map[string]string) enrichments {
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
	return e
}

func setupPermissions(e enrichments, props map[string]string) enrichments {
	e.CanView = true
	e.CanEdit = true
	e.CanDelete = true
	e.CanNew = true
	e.CanSave = true
	e.CanList = true

	if strings.ToUpper(props["can_view"]) == "N" {
		e.CanView = false
	}
	if strings.ToUpper(props["can_edit"]) == "N" {
		e.CanEdit = false
	}
	if strings.ToUpper(props["can_delete"]) == "N" {
		e.CanDelete = false
	}
	if strings.ToUpper(props["can_new"]) == "N" {
		e.CanNew = false
	}
	if strings.ToUpper(props["can_save"]) == "N" {
		e.CanSave = false
		e.CanEdit = false
		e.CanNew = false
	}
	if strings.ToUpper(props["can_list"]) == "N" {
		e.CanList = false
	}

	//spew.Dump(e)

	return e
}
