package main

import (
	"runtime"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	core "github.com/mt1976/mwt-goToolkit/core"
	"github.com/mt1976/mwt-goToolkit/logs"
)

func displayApplicationHeader() {
	//fmt.Printf("core.Properties: %v\n", core.Properties)
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
