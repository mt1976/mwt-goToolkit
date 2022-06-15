package main

import (
	"runtime"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/mt1976/templateBuilder/logs"
	core "github.com/mt1976/templatebuiler/core"
)

func displayApplicationHeader() {

	logs.Break()

	logs.Header("Application Information")
	logs.Break()

	logs.Header("Application")
	logs.Information("Name", core.Properties["appname"])
	logs.Information("Host Name", getHostName())

	logs.Information("Server Release", genReleaseName())
	logs.Information("Server Date", time.Now().Format(core.DATEFORMATUSER))

	logs.Information("Licence", core.Properties["licname"])
	logs.Information("Lic URL", core.Properties["liclink"])
	logs.Header("Runtime")
	logs.Information("GO Version", runtime.Version())
	logs.Information("Operating System", runtime.GOOS+" ("+runtime.GOARCH+")")

	logs.Default("Working Directory", getPWD())
	logs.Information("User", getUsername())
	logs.Header("Connectivity")
	logs.Default("Input", data_in())
	logs.Default("Output", data_out())
}

func displayTableHeader(in string) {
	// logs.Break()
	// logs.Header(in + " Information")
	// logs.Break()
	// logs.Information("MD ", "Mandatory")
	// logs.Information("CR ", "Core Fields")
	// logs.Information("EX ", "Extra Fields")
	// logs.Information("OV", " Override of a Core Field")
	// logs.Information("OL", " You can use this field to lookup a value from a table")
	// logs.Information("LL", " You can use this field to lookup from a fixed list in lists.cfg")
	// logs.Information("FL", " The display value of this field is displayed a value from a lookup")
	// logs.Information("IN", " Input Field")

	// logs.Break()
	// info := fmt.Sprintf(tableHeader, "Field Name", "Type", "Default", "MD", "CR", "EX", "OV", "L⬆", "⬆ Object", "⬆ Field", "⬇ Value", "IN")
	// logs.Information(info, "")
	// logs.Break()
}
