package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"

	core "github.com/mt1976/mwt-goToolkit/core"
)

func wrapVariable(in string) string {
	return "{{." + in + "}}"
}

func wrapTemplate(in string) string {
	return "{{template " + enquote(in) + " .}}"
}

func enquote(in string) string {
	return "\"" + in + "\""
}

func getUsername() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if usr.Name != "" {
		return usr.Username + " (" + usr.Name + ")"
	}

	return usr.Username
}

func genUUID() string {
	id := uuid.New()
	//fmt.Printf("github.com/google/uuid:         %s\n", id.String())
	return id.String()
}

func genReleaseName() string {
	return fmt.Sprintf("%s [r%s-%s]",
		core.Properties["releaseid"],
		core.Properties["releaselevel"],
		core.Properties["releasenumber"])
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
	if core.Properties["deliverto"] != "" {
		do = core.Properties["deliverto"]
	} else {
		do = getPWD() + core.Properties["data_out"]
	}
	return do
}

func data_in() string {
	//fmt.Printf("core.Properties: %v\n", core.Properties)
	return strings.TrimSpace(core.Properties["data_in"])
}

func enrichmentType(inSource string, inType string) bool {
	return strings.EqualFold(inSource, inType)
}

func tf(in bool) string {
	if in {
		return "Y"
	}
	return ""
}

func getProperty(name string, props map[string]string) bool {
	return strings.EqualFold(props[name], "Y")
}
