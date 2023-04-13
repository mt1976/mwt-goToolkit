package main

import (
	_ "github.com/denisenkom/go-mssqldb"
)

type ObjectDefinition struct {
	ObjectName             string
	ObjectNameLower        string
	ObjectCamelCase        string
	ObjectGlyph            string
	ObjectTextClass        string
	EndpointRoot           string
	QueryString            string
	QueryField             string
	QueryFieldID           string
	SourceType             string
	Version                string
	Date                   string
	Time                   string
	Who                    string
	Host                   string
	FieldsList             []FieldProperties
	FriendlyName           string
	SQLTableName           string
	SQLSearchID            string
	SearchKey              string
	MessageList            []messages
	Path                   string
	ProjectRepo            string
	UUID                   string
	Title                  string
	PageTitle              string
	UserMenu               string
	MenuHeader             string
	RangeUserMenuStart     string
	RangeEnd               string
	MenuHREF               string
	MenuOnClick            string
	MenuGlyph              string
	MenuTextClass          string
	MenuText               string
	ItemsOnPageWc          string
	ItemList               string
	RangeItemList          string
	CanView                bool
	CanEdit                bool
	CanSave                bool
	CanNew                 bool
	CanDelete              bool
	CanList                bool
	CanAPI                 bool
	CanSoftDelete          bool
	PropertiesName         string
	HasStoreAdaptor        bool
	HasFetchAdaptor        bool
	HasAudit               bool
	CanExport              bool
	ProvidesReverseLookup  bool
	ReverseLookup          string
	IsSpecial              bool
	ProvidesLookup         bool
	LookupID               string
	LookupName             string
	TemplateHeader         string
	TemplateUserFooter     string
	TemplatePageFooter     string
	TemplateScripts        string
	TemplateAudit          string
	TemplateBody           string
	TemplateListControls   string
	TemplateExportControls string
	TitleText              string
	MonitorPath            string
	HasMonitor             bool
	CanOverrideID          bool
	Artifacts              []artifact
	HasJob                 bool
	CanImport              bool
	CanDo                  bool
	DoesLookup             bool
	DoesListLookup         bool
	ObjectPackage          string
	WrapContext            string
	HasCrossval            bool
	IsFilteredLookup       bool
	HasPostPutAction       bool
}

type FieldProperties struct {
	FieldName                string
	Type                     string
	Default                  string
	FieldSQL                 string
	Formatted                string
	TemplateField            string
	Disabled                 string
	Hidden                   string
	ValueID                  string
	IsMandatory              bool
	IsUserField              bool
	IsBaseField              bool
	IsLookup                 bool
	IsOverride               bool
	IsExtra                  bool
	IsListLookup             bool
	IsFetch                  bool
	IsHelper                 bool
	HelperHTML               string
	LookupObject             string
	LookupField              string
	LookupValue              string
	RangeHTML                string
	WrapFieldName            string
	IsNoChange               bool
	HasCallout               bool
	IsAudit                  bool
	FieldType                string
	FieldMask                string
	IsKey                    bool
	WrapPropsMsgType         string
	WrapPropsMsgFeedBackType string
	WrapPropsMsgMessage      string
	WrapPropsMsgGlyph        string
	NumericStep              string
	IsFilteredLookup         bool
	IsCheckedHTML            string
}

type messages struct {
	Message string
}

type artifact struct {
	Name     string
	Path     string
	Type     string
	FilePath string
}

const (
	go_template   = ".go_template"
	html_template = ".html_template"
	json_template = ".json_template"
	nfo_template  = ".nfo_template"
	tableHeader   = "| %-35s | %-10s | %-10s | %-2s | %-2s | %-2s | %-2s | %-2s | %-15s | %-24s | %-24s | %-2s = No."
	tableRow      = "| %-35s | %-10s | %-10s | %-2s | %-2s | %-2s | %-2s | %-2s | %-15s | %-24s | %-24s | %-2s |"
	overrideField = "Override"
	lookupField   = "Lookup"
	extraField    = "Extra"
	listField     = "List"
	fetchField    = "Fetch"
	defaultField  = "Default"
	helperField   = "Helper"

	enri_Type         = 0
	enri_Field        = 1
	enri_LookupObject = 2
	enri_LookupKey    = 3
	enri_LookupValue  = 4
	enri_IsInputtable = 5
	enri_IsMandatory  = 6
	enri_DefaultValue = 7
	enri_FieldType    = 8
	enri_NoChange     = 9
	enri_HasCallout   = 10
	enri_Mask         = 11
	enri_Hidden       = 12
	enri_Min          = 13
	enri_Max          = 14
	enri_Filter       = 15

	html_disabled  = "readonly=\"true\""
	html_hidden    = "hidden"
	html_mandatory = "required"

	rangeHTMLString = "{{range .%s}}<option value=\"%s\" {{if eq .ID %s}}selected{{end}} data-mdb-secondary-text=\"{{.ID}}\">%s</option>{{end}}<option value=\"\" {{if eq \"\" .%s}}selected{{end}} data-mdb-secondary-text=\"\"></option>"
)
