package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	core "github.com/mt1976/mwt-goToolkit/core"
	"github.com/mt1976/mwt-goToolkit/das"
	"github.com/mt1976/mwt-goToolkit/logs"
)

func main() {

	logs.Break()
	logs.Header("Template Generator")
	logs.Break()

	core.Initialise()

	displayApplicationHeader()

	logs.Break()
	logs.Activity("Searching for work...", data_in())
	logs.Break()

	pwd, _ := os.Getwd()

	clItem := ""
	//log.Println(os.Args[1:], len(os.Args[1:]))
	//log.Println(os.Args[len(os.Args)-1], len(os.Args[len(os.Args)-1]))
	if len(os.Args[len(os.Args)-1]) > 1 {
		clItem = pwd + data_in() + "/" + os.Args[len(os.Args)-1]
	}
	//log.Println(clItem)

	var paths []string
	if clItem == "" {
		// Get list of files from a folder
		logs.Activity("Searching...", data_in())
		paths = seekTableDefinitions(data_in())
		logs.Success("Object Definition Files in " + data_in())
	} else {
		logs.Information("Object Definition File", clItem+".cfg")
		paths = append(paths, clItem+".cfg")
	}

	noFiles := len(paths)

	logs.Information("Object Definition File(s) Found ", fmt.Sprintf("%d %s %s", noFiles, " files in ", data_in()))

	// loop through files from Paths

	logs.Break()

	for i := 0; i < noFiles; i++ {
		// if last four character in paths[i] are ".cfg" then proceed otherwise skip this item
		fileExtension := paths[i][len(paths[i])-4:]
		//fmt.Println(fileExtension) // gives "lo"
		if fileExtension == ".cfg" {
			processObjectDefinition(paths[i])
		}
	}
	logs.Break()
	logs.Success("Templating Complete")
	logs.Break()
}

// Get list of files from a folder
func seekTableDefinitions(dir string) []string {
	logs.Information("Searching...", "")
	dir = getPWD() + dir + "/"
	//logs.Information("In Queue Path", dir)
	files, err := os.ReadDir(dir)
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

func processObjectDefinition(configFile string) {
	logs.Processing(configFile)
	//	logs.Information("Populate", "Replacement Values")
	//logs.Information("sausage", "")
	props := core.Config_Get(configFile)

	//fmt.Printf("props: %v\n", props)

	e := setupObjectEnrichment(props)

	csvPath := getPWD() + data_in() + "/" + e.ObjectName + ".csv"
	enriPath := getPWD() + data_in() + "/" + e.ObjectName + ".enri"
	logs.Information("CSV  Path", csvPath)
	logs.Information("Enri Path", enriPath)

	if props["use"] == "db" {
		// Do nothing for now
		logs.Information("Getting List of fields from DB", props["server"]+" "+props["database"]+" "+props["tablename"])
		e = getFieldDefinitions_DB(e, props)
		e.SourceType = "Application"
	} else {
		e.SourceType = "Application"
		if getProperty("HasFetchAdaptor", props) {
			e.SourceType = "External"
		}
		if getProperty("HasStoreAdaptor", props) {
			e.SourceType = "External"
		}
		logs.Information("Getting List of fields from CSV", csvPath)
		e = getFieldDefinitions_CSV(csvPath, e)
	}

	if getProperty("hasenrichments", props) {
		//logs.Break()
		logs.Information("Getting Enrichment Fields from enri", enriPath)
		e = mergeEnrichmentDefinitions(enriPath, e)
	}

	// for i := 0; i < len(e.FieldsList); i++ {
	// 	logs.Information(e.FieldsList[i].FieldName, strconv.Itoa(i))
	// }
	logs.Break()
	logs.Header("Generating Artifacts")
	logs.Break()

	e = generateCodeArtifact("application", props, configFile, e)

	e = generateCodeArtifact("adaptor", props, configFile, e)

	e = generateCodeArtifact("validation", props, configFile, e)

	if e.CanAPI {
		e = generateCodeArtifact("api", props, configFile, e)
	}

	e = generateCodeArtifact("dao", props, configFile, e)

	e = generateCodeArtifact("datamodel", props, configFile, e)

	e = generateCodeArtifact("job", props, configFile, e)

	e = generateCodeArtifact("menu", props, configFile, e)

	e = generateHTMLArtifacts("html", props, configFile, e)

	e = generateCodeArtifact("catalog", props, configFile, e)

	e = generateCodeArtifact("monitor", props, configFile, e)

	//spew.Dump(e)
}

func generateCodeArtifact(a string, props map[string]string, configFile string, e ObjectDefinition) ObjectDefinition {

	if getProperty("create_"+a, props) || a == "catalog" {
		e = processCodeArtifact(a, configFile, a, e)
	} else {
		logs.Skipping(a)
	}
	return e
}

func processCodeArtifact(w string, p string, destFolder string, e ObjectDefinition) ObjectDefinition {
	//logs.Processing(w)

	in_extn := ".go_template"
	out_extn := "_core.go"

	if destFolder == "catalog" {
		destFolder = "design/catalog"
		out_extn = ".md"
		in_extn = ".nfo_template"
	}

	// if destFolder == "application" {
	// 	out_extn = "_core" + out_extn
	// }

	if destFolder == "api" {
		destFolder = "application"
		out_extn = "_api.go"
		in_extn = ".go_template"
	}

	if destFolder == "menu" {
		destFolder = "design/menu"
		out_extn = ".json"
		in_extn = ".json_template"
	}

	if destFolder == "html" {
		in_extn = ".html_template"
	}

	if destFolder == "monitor" {
		destFolder = "application"
		out_extn = "_monitor_impl.go_template"
		in_extn = ".go_template"
	}

	if destFolder == "job" {
		destFolder = "jobs"
		out_extn = "_core.go"
		in_extn = ".go_template"
	}

	if destFolder == "adaptor" {
		destFolder = "dao"
		out_extn = "_adaptor.go_template"
		in_extn = ".go_template"
	}

	if destFolder == "validation" {
		destFolder = "dao"
		out_extn = "_validation.go_template"
		in_extn = ".go_template"
	}

	if core.Properties["deliverto"] == "" {
		out_extn = out_extn + "_tmp"
		if destFolder == "design/catalog" {
			out_extn = ".md"
		}
	}
	//spew.Dump(replacements)
	fp := e.Path + "/templates/" + w + in_extn

	t, err := template.ParseFiles(fp)
	if err != nil {
		logs.Error("Load Template :", err)
	}
	dfp := "/" + destFolder
	dest := dfp + "/" + e.ObjectCamelCase + out_extn
	//fullname := data_out() + dest
	//fmt.Printf("dfp: %v\n", dfp)
	//fmt.Printf("dest: %v\n", dest)
	//fmt.Printf("fullname: %v\n", fullname)

	fullPath := data_out() + dfp //+ "/" + e.ObjectPackage
	fullname := fullPath + "/" + e.ObjectCamelCase + out_extn

	//fmt.Printf("dfp: %v\n", dfp)
	//fmt.Printf("dest: %v\n", dest)
	//fmt.Printf("fullname: %v\n", fullname)
	//fmt.Printf("fullPath: %v\n", fullPath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		logs.Created(fullPath)
		os.MkdirAll(fullPath, 0700)
	}

	f, err := os.Create(fullname)
	if err != nil {
		logs.Error("Create file : ", err)
		return e
	}

	//	e.MessageList = append(e.MessageList, messages{Message: "* " + w + " (" + dest + ")"})

	err2 := t.Execute(f, e)
	if err2 != nil {
		logs.Error("Process Template", err2)
	}
	f.Close()
	logs.Created(f.Name())

	e = logArtifact(w, dest, e, "code", f.Name())

	return e
}

func logArtifact(inName string, fileName string, e ObjectDefinition, inType string, filePath string) ObjectDefinition {
	art := artifact{Name: inName, Path: fileName, Type: inType, FilePath: filePath}
	e.Artifacts = append(e.Artifacts, art)
	return e
}

func generateHTMLArtifacts(a string, props map[string]string, configFile string, e ObjectDefinition) ObjectDefinition {
	destinationFolder := "html/base"
	if strings.ToUpper(props["create_html"]) == "Y" {
		if e.CanList {
			e = generateHTMLArtifact("list", configFile, destinationFolder, e)
		} else {
			logs.Skipping("Listing is not enabled for this object")
		}

		if e.CanView {
			e = generateHTMLArtifact("view", configFile, destinationFolder, e)
		} else {
			logs.Skipping("Viewing is not enabled for this object")
		}

		if e.CanEdit {
			e = generateHTMLArtifact("edit", configFile, destinationFolder, e)
		} else {
			logs.Skipping("Editing is not enabled for this object")
		}

		if e.CanNew {
			e = generateHTMLArtifact("new", configFile, destinationFolder, e)
		} else {
			logs.Skipping("Creating is not enabled for this object")
		}

		if !e.CanExport {
			logs.Skipping("Exporting is not enabled for this object")
		}

	} else {
		logs.Skipping("html")
	}
	return e
}

func generateHTMLArtifact(w string, p string, destFolder string, e ObjectDefinition) ObjectDefinition {
	//logs.Processing(w + html_template)

	userAction := strings.ToUpper(w[:1]) + w[1:]

	//spew.Dump(replacements)
	fp := e.Path + "/templates/" + w + html_template

	t, err := template.ParseFiles(fp)
	if err != nil {
		logs.Error("Load Template", err)
	}

	dfp := "/" + destFolder + "/" + e.ObjectName
	dest := "/" + e.ObjectName + userAction + ".html"

	//dest := dfp + "/" + e.ObjectCamelCase

	pathname := data_out() + dfp
	fullname := pathname + dest

	//fmt.Printf("dfp: %v\n", dfp)

	//fmt.Printf("dest: %v\n", dest)
	//fmt.Printf("pathname: %v\n", pathname)
	//fmt.Printf("fullname: %v\n", fullname)

	if _, err := os.Stat(fullname); os.IsNotExist(err) {

		logs.Created(pathname)

		os.MkdirAll(pathname, 0700)

	}

	f, err := os.Create(fullname)

	if err != nil {

		logs.Error("Create file  : ", err)

		return e

	}

	// f, err := os.Create(data_out() + dest)
	// if err != nil {
	// 	logs.Error("Create file: ", err)
	// 	return e
	// }
	//spew.Dump(e)
	err2 := t.Execute(f, e)
	if err2 != nil {
		logs.Error("Process Template", err2)
	}
	f.Close()
	//e.MessageList = append(e.MessageList, messages{Message: "* html -> " + userAction + " (" + dest + ")"})
	logs.Created(f.Name())
	e = logArtifact(w, dest, e, "html", f.Name())
	return e
}

func getFieldDefinitions_CSV(filePath string, e ObjectDefinition) ObjectDefinition {
	// Load a csv file.
	//logs.Information("Read CSV", filePath)
	f, _ := os.Open(filePath)
	//logs.Information("File Open", filePath)
	// Create a new reader.
	r := csv.NewReader(f)
	//logs.Information("New Reader", filePath)
	//displayTableHeader("Table")

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

		colMand := false
		if record[3] == "true" {
			colMand = true
		}
		noInput := false
		if record[4] == "true" {
			noInput = true
		}
		if record[0] == "Name" && record[1] == "Type" {
			continue
		}
		e.FieldsList = addField(e, record[0], record[1], record[2], colMand, noInput)

	}
	logs.Break()
	return e
}

func getFieldDefinitions_DB(e ObjectDefinition, p map[string]string) ObjectDefinition {
	// Open Database Connection
	db, err := core.GlobalsDatabaseConnect(p)
	if err != nil {
		logs.Error("Database Connection", err)
	}
	//fmt.Printf("db: %v\n", db)

	//tsql := fmt.Sprintf("USE %s EXEC sp_columns '%s'", p["database"], p["sqltablename"])
	tsql := fmt.Sprintf("EXEC sp_columns '%s', @table_owner = '%s'", p["sqltablename"], p["schema"])
	logs.Query(tsql)
	results, noFields, err := das.Query(db, tsql)
	//fmt.Printf("results: %v\n", results)
	//fmt.Printf("noFields: %v\n", noFields)
	//spew.Dump(results)
	if noFields == 0 {
		logs.Error("No Fields Found", err)
	} else {
		logs.Success("Fields Found =" + strconv.Itoa(noFields))
	}
	//displayTableHeader("Table")
	for _, row := range results {
		colName := row["COLUMN_NAME"].(string)
		colType := row["TYPE_NAME"].(string)
		colMand := false
		if row["IS_NULLABLE"].(string) == "NO" {
			colMand = true
		}

		//logs.Message("found", fmt.Sprintf("name %v type %v nullable %v mandatory %t\n", colName, colType, row["IS_NULLABLE"], colMand))
		if colName == "ID" {
			colMand = true
		}
		colDefault := ""
		switch colType {
		case "varchar", "nvarchar", "char", "nchar", "text", "ntext":
			colDefault = ""
			colType = "String"
		case "int", "bigint", "smallint", "tinyint", "int64":
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
		case "bigint identity":
			colDefault = "0"
			colType = "Int"
		case "bit":
			colDefault = "True"
			colType = "Bool"
		default:
			colType = "String"
			colDefault = ""
		}
		e.FieldsList = addField(e, colName, colType, colDefault, colMand, false)
	}
	return e
}

func setupObjectEnrichment(props map[string]string) ObjectDefinition {

	e := ObjectDefinition{ObjectName: props["objectname"]}
	//capitalize first character of enrichment.ObjectName
	logs.Information("Object Name", e.ObjectName)
	logs.Accessing("Object Name " + e.ObjectName)
	logs.Processing("Object Name " + e.ObjectName)
	logs.Servicing("Object Name " + e.ObjectName)
	logs.System("Object Name " + e.ObjectName)
	logs.Default("Object Name", e.ObjectName)
	logs.Created("Object Name " + e.ObjectName)

	//caser := cases.Title(language.English)
	//e.ObjectName = caser.String(e.ObjectName)
	e.ObjectCamelCase = strings.ToLower(e.ObjectName[:1]) + e.ObjectName[1:]
	e.ObjectNameLower = strings.ToLower(e.ObjectName)
	e.Version = genReleaseName()
	e.Time = time.Now().Format(core.TIMEFORMATUSER)
	e.Date = time.Now().Format(core.DATEFORMATUSER)
	e.Host = getHostName()
	e.Who = getUsername()
	e.DoesLookup = false
	e.DoesListLookup = false

	e.FriendlyName = props["friendlyname"]
	if e.FriendlyName == "" {
		e.FriendlyName = e.ObjectCamelCase
	}
	e.SQLTableName = props["sqltablename"]
	e.SQLSearchID = strings.TrimSpace(props["sqlsearchid"])
	e.SearchKey = props["searchkey"]
	if e.SearchKey == "" {
		e.SearchKey = props["queryfield"]
	}
	e.QueryString = props["querystring"]
	e.QueryField = "{{." + props["queryfield"] + "}}"
	e.QueryFieldID = props["queryfield"]
	if props["endpointroot"] == "" {
		e.EndpointRoot = e.ObjectName
	} else {
		e.EndpointRoot = props["endpointroot"]
	}
	e.EndpointRoot = strings.ToUpper(e.EndpointRoot[:1]) + e.EndpointRoot[1:]

	e.ObjectPackage = e.ObjectName
	if props["package"] != "" {
		e.ObjectPackage = strings.ToUpper(props["package"][:1]) + props["package"][1:]
	}
	logs.Default("Object Package ", e.ObjectPackage)

	//fmt.Printf("props[\"package\"]: %v\n", props["package"])

	e.Path = getPWD()
	e.ObjectGlyph = props["objectglyph"]
	e.ObjectTextClass = props["textclass"]
	e.ProjectRepo = props["projectrepo"] + "/"
	e.UUID = genUUID()

	e.PropertiesName = ""
	e.HasStoreAdaptor = getProperty("hasstoreadaptor", props)
	// if props["hasstoreadaptor"] == "y" {
	// 	e.HasStoreAdaptor = true
	// }
	e.HasFetchAdaptor = getProperty("hasfetchadaptor", props)

	// e.HasFetchAdaptor = false
	// if props["hasfetchadaptor"] == "y" {
	// 	e.HasFetchAdaptor = true
	// }

	// e.TemplateAudit = ""
	// if props["hasaudit"] == "y" {
	// 	e.TemplateAudit = wrapTemplate("audit")
	// }
	e.TemplateAudit = ""
	e.HasAudit = getProperty("hasaudit", props)
	if e.HasAudit {
		e.TemplateAudit = wrapTemplate("audit")
	}

	e.IsSpecial = false

	if props["propertiesoverride"] == "" {
		e.PropertiesName = "Application"
		//e.TemplateAudit = wrapTemplate("audit")
	} else {
		e.HasStoreAdaptor = true
		if props["propertiesoverride"] == "special" {
			e.PropertiesName = "Application"
			e.IsSpecial = true
		} else {
			e.PropertiesName = props["propertiesoverride"]
		}
	}

	e.HasPostPutAction = false
	if props["haspostputaction"] == "y" {
		e.HasPostPutAction = true
	}

	e = setupTemplateEnrichment(e, props)

	e = setupPermissions(e, props)

	e.ProvidesReverseLookup = false
	e.ReverseLookup = ""
	if props["reverselookup"] != "" {
		e.ProvidesReverseLookup = true
		e.ReverseLookup = props["reverselookup"]
	}

	// e.IsSpecial = false
	// if strings.ToUpper(props["isspecial"]) == "Y" {
	// 	e.IsSpecial = true
	// }

	e.ProvidesLookup = false
	e.ProvidesLookup = getProperty("provideslookup", props)

	if e.ProvidesLookup {
		//e.ProvidesLookup = true
		e.LookupID = props["lookupid"]
		e.LookupName = props["lookupname"]
	}

	e.TemplateHeader = wrapTemplate("header")
	e.TemplateUserFooter = wrapTemplate("userfooter")
	e.TemplatePageFooter = wrapTemplate("pagefooter")
	e.TemplateScripts = wrapTemplate("scripts")
	e.TemplateBody = wrapTemplate("bodydefinition")
	e.TemplateListControls = wrapTemplate("tablecontrols")
	e.TemplateExportControls = wrapTemplate("exportcontrols")

	e.MonitorPath = props["monitorpath"]
	e.HasMonitor = getProperty("hasmonitor", props)

	e.CanOverrideID = getProperty("canoverrideid", props)

	e.WrapContext = wrapVariable("Context")
	e.HasCrossval = getProperty("crossvalidate", props)
	//spew.Dump(e)
	//fmt.Printf("e: %v\n", e)
	return e
}

func setupTemplateEnrichment(e ObjectDefinition, props map[string]string) ObjectDefinition {
	e.Title = wrapVariable("Title")
	e.PageTitle = wrapVariable("PageTitle")
	e.UserMenu = wrapVariable("UserMenu")
	e.MenuHeader = "{{ (index .UserMenu 0).MenuHeaderText}}"
	e.RangeUserMenuStart = "{{range .UserMenu}}"
	e.RangeEnd = "{{end}}"
	e.MenuHREF = wrapVariable("MenuHREF")
	e.MenuOnClick = wrapVariable("MenuOnClick")
	e.MenuGlyph = wrapVariable("MenuGlyph")
	e.MenuTextClass = wrapVariable("MenuTextClass")
	e.MenuText = wrapVariable("MenuText")
	e.ItemsOnPageWc = wrapVariable("ItemsOnPage")
	e.ItemList = wrapVariable("ItemList")
	e.RangeItemList = "{{range .ItemList}}"
	return e
}

func setupPermissions(e ObjectDefinition, props map[string]string) ObjectDefinition {

	e.CanView = getProperty("can_view", props)
	e.CanEdit = getProperty("can_edit", props)
	e.CanDelete = getProperty("can_delete", props)
	e.CanNew = getProperty("can_new", props)
	e.CanSave = getProperty("can_save", props)
	if !e.CanSave {
		e.CanSave = false
		e.CanEdit = false
		e.CanNew = false
	}
	e.CanList = getProperty("can_list", props)
	e.CanExport = getProperty("can_export", props)
	e.CanAPI = getProperty("can_api", props)
	e.CanDo = getProperty("can_do", props)
	e.CanSoftDelete = getProperty("can_softdelete", props)
	//if e.CanSoftDelete {
	//	e.CanDelete = false
	//}
	//spew.Dump(e)
	//spew.Dump(e)

	return e
}

func addField(en ObjectDefinition, fn string, tp string, df string, mand bool, noInput bool) []FieldProperties {

	en.FieldsList = addComplexField(en, fn, tp, df, mand, true, false, "", "", "", "", noInput, false, false, false, false, false, false)

	return en.FieldsList
}

func addComplexField(en ObjectDefinition, fn string, tp string, df string, mand bool, baseField bool, isLookup bool, lkObject string, lkKeyField string, lkValueField string, lkRange string, noinp bool, isExtra bool, isOverride bool, isListLookup bool, isFetch bool, isHidden bool, isFiltered bool) []FieldProperties {

	// log parameters

	//log.Println("addComplexField:"+fn+" "+tp+" "+df+" "+strconv.FormatBool(mand)+" "+strconv.FormatBool(baseField)+" "+strconv.FormatBool(isLookup)+" "+lkObject+" "+lkKeyField+" "+lkValueField+" "+lkRange+" "+strconv.FormatBool(noinp), strconv.FormatBool(isExtra), strconv.FormatBool(isOverride))

	origfn := fn

	//if first charachter of fieldName is _ then replace _ with SYS

	noinput := ""
	hidden := ""
	userField := true
	//logs.Processing("fn: " + fn)
	if isAudit(fn) {
		//logs.Processing("isAudit: " + fn)
		//Convert fn to Title Case
		fn = strings.Replace(fn, "_", "", -1)
		fn = strings.ToUpper(fn[:1]) + fn[1:]
		//fmt.Println(fn)
		fn = "SYS" + fn
		noinput = "hidden"
		hidden = "hidden"
		userField = false
	}

	if isHidden {
		hidden = "hidden"
	}

	if noinp {
		noinput = html_disabled
	}
	fn = strings.ToUpper(fn[:1]) + fn[1:]

	//info := fmt.Sprintf(tableRow, fn, tp, df, tf(mand), tf(baseField), tf(isExtra), tf(isOverride), lkval, lkObject, lkKeyField, lkValueField)
	tplField := "{{." + fn + "}}"

	hasAPI := false
	if isExtra {
		hasAPI = true
	}

	isKey := false
	//fmt.Printf("fn: %v\n", fn)
	//fmt.Printf("en.SearchKey: %v\n", en.SearchKey)
	//spew.Dump(en)
	if fn == en.SearchKey {
		isKey = true
	}

	en.FieldsList = append(en.FieldsList, FieldProperties{FieldName: fn,
		Type:                     tp,
		Default:                  df,
		FieldSQL:                 origfn,
		Formatted:                "",
		TemplateField:            tplField,
		Disabled:                 noinput,
		Hidden:                   hidden,
		ValueID:                  wrapVariable(fn),
		IsMandatory:              mand,
		IsUserField:              userField,
		IsBaseField:              baseField,
		IsLookup:                 isLookup,
		LookupObject:             lkObject,
		LookupField:              lkKeyField,
		LookupValue:              lkValueField,
		RangeHTML:                lkRange,
		IsExtra:                  isExtra,
		IsOverride:               isOverride,
		IsListLookup:             isListLookup,
		HasCallout:               hasAPI,
		IsAudit:                  isAudit(origfn),
		FieldType:                "text",
		IsKey:                    isKey,
		WrapPropsMsgType:         wrapVariable(fn + "_props.MsgType"),
		WrapPropsMsgFeedBackType: wrapVariable(fn + "_props.MsgFeedBackType"),
		WrapPropsMsgMessage:      wrapVariable(fn + "_props.MsgMessage"),
		WrapPropsMsgGlyph:        wrapVariable(fn + "_props.MsgGlyph"),
		IsFilteredLookup:         isFiltered,
		IsCheckedHTML:            wrapVariable(fn + "_checked"),
	})

	//logs.Information(info, "")

	return en.FieldsList
}

func isAudit(fn string) bool {
	if len(fn) < 1 {
		return false
	}
	return fn[0:1] == "_"
}

func mergeEnrichmentDefinitions(filePath string, en ObjectDefinition) ObjectDefinition {

	//logs.Information("Read CSV", filePath)
	f, _ := os.Open(filePath)
	//logs.Information("File Open", filePath)
	// Create a new reader.
	r := csv.NewReader(f)
	//logs.Information("New Reader", filePath)
	//displayTableHeader("Enrichment")
	var enrichmentDefinitions [][]string

	// for i := 0; i < len(en.FieldsList); i++ {
	// 	fmt.Printf("b4 en: %d %v\n", i, en.FieldsList[i])
	// }
	for {
		enrichmentDefinition, err := r.Read()
		//fmt.Printf("record: %v\n", record)
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			logs.Warning("Read Enri :" + err.Error())

			panic(err)
		}
		if enrichmentDefinition[enri_Type] == "Type" && enrichmentDefinition[enri_Field] == "Field" {
			//logs.Information("Found", "Enrichments")
		} else {

			if enrichmentType(enrichmentDefinition[enri_Type], extraField) {
				// Add additional "Extra" fields to the object definition as specfied in .../?.enri
				//			logs.Information("Found", "Extra Field")
				en.FieldsList = addExtraTypeFields(enrichmentDefinition, en)
			} else {
				// add enrichmentDefinition to enrichmentDefinitions list

				enrichmentDefinitions = append(enrichmentDefinitions, enrichmentDefinition)
			}
		}
	}
	//fmt.Printf("enrichmentDefinitions: %v\n", enrichmentDefinitions)
	//loop through enrichmentDefinitions
	for _, thisEnrichment := range enrichmentDefinitions {
		//	fmt.Printf("enrichmentOverride: %v\n", enrichmentOverride)
		//fmt.Printf("thisEnrichment: %v\n", thisEnrichment)
		//fmt.Printf("thisEnrichmentType: %v\n", thisEnrichment[enri_Type])

		switch {
		case enrichmentType(thisEnrichment[enri_Type], listField):
			//logs.Processing(listField + " " + thisEnrichment[enri_Field])
			en.DoesListLookup = true
			en.FieldsList = mergeComplexField(en, thisEnrichment[enri_Field], thisEnrichment[enri_Type], thisEnrichment)

		case enrichmentType(thisEnrichment[enri_Type], lookupField):
			//logs.Processing(lookupField + " " + thisEnrichment[enri_Field])
			en.DoesLookup = true
			en.FieldsList = mergeComplexField(en, thisEnrichment[enri_Field], thisEnrichment[enri_Type], thisEnrichment)

		case enrichmentType(thisEnrichment[enri_Type], helperField):
			//logs.Processing(helperField + " " + thisEnrichment[enri_Field])
			en.DoesLookup = true
			en.FieldsList = mergeComplexField(en, thisEnrichment[enri_Field], thisEnrichment[enri_Type], thisEnrichment)

		case enrichmentType(thisEnrichment[enri_Type], extraField):
			// Do Nothing
			logs.Information(extraField, thisEnrichment[enri_Field])

		case enrichmentType(thisEnrichment[enri_Type], overrideField):
			// Do Nothing
			//logs.Information(overrideField, thisEnrichment[enri_Field])
			en.FieldsList = mergeComplexField(en, thisEnrichment[enri_Field], thisEnrichment[enri_Type], thisEnrichment)

		case enrichmentType(thisEnrichment[enri_Type], fetchField):
			// Do Nothing
			//logs.Processing(fetchField + " " + thisEnrichment[enri_Field])
			en.FieldsList = mergeComplexField(en, thisEnrichment[enri_Field], thisEnrichment[enri_Type], thisEnrichment)

		case enrichmentType(thisEnrichment[enri_Type], defaultField):

			//logs.Processing(fetchField + " " + thisEnrichment[enri_Field])

			en.FieldsList = mergeComplexField(en, thisEnrichment[enri_Field], thisEnrichment[enri_Type], thisEnrichment)

		default:
			// Do Nothing
			logs.Warning("Unkown Enrichment Type" + thisEnrichment[enri_Type] + thisEnrichment[enri_Field])
		}
	}
	noRows := len(en.FieldsList)
	for i := 0; i < noRows; i++ {
		//fmt.Printf("AF en: %d %v\n", i, en.FieldsList[i])
		op := en.FieldsList[i]
		lkVal := ""
		switch {
		case op.IsLookup:
			lkVal = "OL"
		case op.IsListLookup:
			lkVal = "LL"
		case op.IsFetch:
			lkVal = "FL"
		}
		//fmt.Printf("lkVal: %v\n", lkVal)
		ipVal := "Y"
		if op.Disabled == html_disabled {
			ipVal = "N"
		}
		if op.Hidden == html_hidden {
			ipVal = "H"
		}
		info := fmt.Sprintf(tableRow, op.FieldName, op.Type, op.Default, tf(op.IsMandatory), tf(op.IsBaseField), tf(op.IsExtra), tf(op.IsOverride), lkVal, op.LookupObject, op.LookupField, op.LookupValue, ipVal)
		en.MessageList = append(en.MessageList, messages{Message: info})
		//	fmt.Printf("info: %v\n", info)
	}
	//logs.Break()
	// for i := 0; i < len(en.FieldsList); i++ {
	// 	fmt.Printf("af en: %d %v\n", i, en.FieldsList[i])
	// }
	return en
}

func addExtraTypeFields(record []string, en ObjectDefinition) []FieldProperties {
	colMand := false
	if record[enri_IsMandatory] == "true" {
		colMand = true
	}

	isLookup := false
	isExtra := false
	isOverride := false
	isListLookup := false
	isFetch := false
	lkObject := ""
	lkKeyField := ""
	lkValueField := ""
	lkRange := ""
	isHidden := false
	isFilteredLookup := false

	suffix := "_Unknown"
	if enrichmentType(record[enri_Type], lookupField) {
		suffix = "_list"
		isLookup = true

		lkObject = record[enri_LookupObject]
		lkKeyField = record[enri_LookupKey]
		lkValueField = record[enri_LookupValue]

		lkRange = buildRangeHTML(record[1])
	}

	if enrichmentType(record[enri_Type], extraField) {
		isExtra = true
		suffix = ""
	}

	if enrichmentType(record[enri_Type], overrideField) {
		isOverride = true
		suffix = ""
	}

	if enrichmentType(record[enri_Type], listField) {
		isListLookup = true
		suffix = "_list"
		lkObject = record[2]
		lkRange = buildRangeHTML(record[1])
	}

	if enrichmentType(record[enri_Type], fetchField) {
		suffix = "_fetch"
		isFetch = true
		lkObject = record[enri_LookupObject]
		lkKeyField = record[enri_LookupKey]
		lkValueField = record[enri_LookupValue]
	}

	noInput := true
	if record[enri_IsInputtable] == "true" {
		noInput = false

	}
	if record[enri_Filter] == "true" {
		isFilteredLookup = true
	}

	return addComplexField(en, record[enri_Field]+suffix, "String", record[enri_DefaultValue], colMand, false, isLookup, lkObject, lkKeyField, lkValueField, lkRange, noInput, isExtra, isOverride, isListLookup, isFetch, isHidden, isFilteredLookup)
}

func buildRangeHTML(inObject string) string {
	return fmt.Sprintf(rangeHTMLString,
		inObject+"_lookup",
		wrapVariable("ID"),
		"$."+inObject,
		wrapVariable("Name"), inObject)
	//		wrapParentVariable(inObject))
}

func mergeComplexField(en ObjectDefinition, fn string, tp string, enrichmentOverride []string) []FieldProperties {

	// log parameters

	//log.Println("addComplexField:"+fn+" "+tp+" "+df+" "+strconv.FormatBool(mand)+" "+strconv.FormatBool(baseField)+" "+strconv.FormatBool(isLookup)+" "+lkObject+" "+lkKeyField+" "+lkValueField+" "+lkRange+" "+strconv.FormatBool(noinp), strconv.FormatBool(isExtra), strconv.FormatBool(isOverride))

	//origfn := ""

	//if first charachter of fieldName is _ then replace _ with SYS

	//noinput := ""
	//hidden := ""
	//userField := true
	noFields := len(en.FieldsList)

	//logs.Break()

	//fmt.Printf("noFields: %v\n", noFields)
	//fmt.Printf("fn: %v\n", fn)
	//fmt.Printf("tp: %v\n", tp)
	for i := 0; i < noFields; i++ {
		//	fmt.Printf("b4 en: %d %v\n", i, en.FieldsList[i])
		if en.FieldsList[i].FieldName == fn {
			//logs.Success("Found " + fn)

			switch {
			case tp == listField:
				//logs.Processing("LIST")
				en.FieldsList[i].IsListLookup = true
				en.FieldsList[i].LookupObject = enrichmentOverride[enri_LookupObject]
				en.FieldsList[i].LookupField = enrichmentOverride[enri_LookupKey]
				en.FieldsList[i].LookupValue = enrichmentOverride[enri_LookupValue]
				en.FieldsList[i].RangeHTML = buildRangeHTML(fn)

				en.FieldsList[i] = commonOverrides(enrichmentOverride, en.FieldsList[i])
			case tp == lookupField:
				//logs.Processing("LOOKUP")
				en.FieldsList[i].IsLookup = true
				en.FieldsList[i].LookupObject = enrichmentOverride[enri_LookupObject]
				en.FieldsList[i].LookupField = enrichmentOverride[enri_LookupKey]
				en.FieldsList[i].LookupValue = enrichmentOverride[enri_LookupValue]
				en.FieldsList[i].RangeHTML = buildRangeHTML(fn)

				en.FieldsList[i] = commonOverrides(enrichmentOverride, en.FieldsList[i])
			case tp == extraField:
				//logs.Skipping("EXTRA")
			case tp == overrideField:
				//logs.Processing("OVERRIDE")
				en.FieldsList[i].IsOverride = true
				en.FieldsList[i] = commonOverrides(enrichmentOverride, en.FieldsList[i])
			case tp == fetchField:
				//logs.Processing("FETCH")
				en.FieldsList[i].IsFetch = true
				en.FieldsList[i].LookupObject = enrichmentOverride[enri_LookupObject]
				en.FieldsList[i].LookupField = enrichmentOverride[enri_LookupKey]
				en.FieldsList[i].LookupValue = enrichmentOverride[enri_LookupValue]
			case tp == defaultField:
				//logs.Processing("DEFAULTING")
				if enrichmentOverride[enri_DefaultValue] != "" {
					en.FieldsList[i].Default = enrichmentOverride[enri_DefaultValue]
				}
				en.FieldsList[i] = commonOverrides(enrichmentOverride, en.FieldsList[i])
			case tp == helperField:
				en.FieldsList[i].IsHelper = true
				en.FieldsList[i].LookupObject = enrichmentOverride[enri_LookupObject]
				en.FieldsList[i].LookupField = enrichmentOverride[enri_LookupKey]
				en.FieldsList[i].LookupValue = enrichmentOverride[enri_LookupValue]
			//	en.FieldsList[i].RangeHTML = buildRangeHTML(fn)

			default:
				logs.Warning("UNKNOWN Enrichment Type: " + tp)
			}

		}

	}

	return en.FieldsList
}

func commonOverrides(commonOverrides []string, fieldsList FieldProperties) FieldProperties {
	//if commonOverrides[enri_IsInputtable] != "" {

	if commonOverrides[enri_IsInputtable] == "false" {
		fieldsList.Disabled = html_disabled
	} else {
		fieldsList.Disabled = ""
	}
	if commonOverrides[enri_IsMandatory] == "true" {
		fieldsList.IsMandatory = true
	} else {
		fieldsList.IsMandatory = false
	}
	if commonOverrides[enri_NoChange] == "true" {
		fieldsList.IsNoChange = true
	} else {
		fieldsList.IsNoChange = false
	}
	if commonOverrides[enri_HasCallout] == "true" {
		fieldsList.HasCallout = true
	} else {
		fieldsList.HasCallout = false
	}
	if commonOverrides[enri_FieldType] != "" {
		fieldsList.FieldType = core.FieldTypes[commonOverrides[enri_FieldType]]
	}
	if commonOverrides[enri_FieldType] == "number" {
		fieldsList.NumericStep = "0.25"
	}
	if commonOverrides[enri_Mask] != "" {
		fieldsList.FieldMask = commonOverrides[enri_Mask]
	}
	if commonOverrides[enri_Hidden] != "" {
		fieldsList.Hidden = "hidden"
	}
	fieldsList.IsFilteredLookup = false
	if commonOverrides[enri_Filter] == "true" {
		fieldsList.IsFilteredLookup = true
	}
	//}
	return fieldsList
}
