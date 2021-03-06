package logs

import (
	"fmt"
	"log"
	"os"
	"strings"

	colour "github.com/TwiN/go-color"
	viper "github.com/spf13/viper"
	"golang.org/x/term"
)

const (
	log_Warning       = "Warning "
	log_Info          = "Information"
	log_MUX           = "Publish"
	log_Error         = "Error"
	log_Success       = "Success"
	log_Panic         = "Panic"
	log_Header        = ""
	log_Poke          = "Fingering"
	log_Schedule      = "Scheduler"
	log_ScheduleStart = "Job Start"
	log_ScheduleEnd   = "Job Ends"
	log_System        = "System"
	log_Accessing     = "Accessing"
	log_Servicing     = "Servicing"
	log_Menu          = "Menu"
	log_Template      = "Template"
	log_Database      = "Database"
	log_URI           = "URI"
	log_Skipping      = "Skipping"
	log_Created       = "Created"
	log_Default       = "Default"
	log_Processing    = "Processing"
	log_Query         = "Query"
	log_Result        = "Result"
	log_Storing       = "Storing"
	log_Activity      = "Activity"

	colourReset       = "\033[0m"
	colourRed         = "\033[31m"
	colourGreen       = "\033[32m"
	colourYellow      = "\033[33m"
	colourBlue        = "\033[34m"
	colourPurple      = "\033[35m"
	colourCyan        = "\033[36m"
	colourWhite       = "\033[37m"
	Character_MapTo   = "⇄"
	Character_Job     = "⚙️"
	Character_Heart   = "🫀"
	Character_Poke    = "👉"
	Character_Time    = "🕒"
	Character_Break   = "≫"
	Character_Tick    = "☑️"
	Character_Warning = "⚠️"
	Character_Bike    = "🚴‍♂️"
	Character_Skip    = "⏭️"
	Character_Created = "💾"
	Character_Query   = "🔎"
	Character_Result  = "?"
	Character_Storing = "📀"
)

type Config struct {
	Verbose     string `mapstructure:"VERBOSE"`
	VerboseMode bool
}

var CFG Config

func init() {
	CFG, _ = getConfig()
	//spew.Dump(CFG)
}

func Poke(w string, v string) {
	msg_raw(log_Poke, w, v, colour.Yellow)
}

func Success(s string) {
	//msg_done(s)
	msg_raw(log_Success, s, Character_Tick, colour.Green)
}

func Skipping(s string) {
	//msg_done(s)
	msg_raw(log_Skipping, s, Character_Skip, colour.Yellow)
}

func Created(s string) {
	//msg_done(s)
	msg_raw(log_Created, s, Character_Created, colour.Blue)
}

func System(s string) {
	//msg_done(s)
	msg_raw(log_System, s, "", colour.Gray)
}

func Default(s string, w string) {
	//msg_done(s)
	msg_raw(log_Default, s, w, colour.Purple+colour.Bold)
}

func Information(w string, v string) {
	//msg_info(w, v)
	if len(v) == 0 {
		msg_raw(log_Info, w, "", colour.Reset)
	} else {
		msg_raw(log_Info, w+" =", colour.Bold+v, colour.Reset)
	}
	//msg_raw(log_Info, w, v, colour.Reset)

}

func Activity(w string, v string) {
	//msg_info(w, v)
	if len(v) == 0 {
		msg_raw(log_Activity, w, "", colour.Reset)
	} else {
		msg_raw(log_Activity, w+" =", colour.Bold+v, colour.Reset)
	}
	//msg_raw(log_Info, w, v, colour.Reset)

}

func Processing(s string) {
	//msg_done(s)
	msg_raw(log_Processing, s, "", colour.White)
}

func Schedule(w string) {
	//msg_info(w, v)
	msg_raw(log_Schedule, w, "", colour.Cyan)

}

func URI(w string) {
	//msg_info(w, v)
	msg_raw(log_URI, colour.Bold+colour.Purple+w, "", colour.Green)
}

func Servicing(w string) {
	//msg_info(w, v)
	msg_raw(log_Servicing, w, "", colour.Yellow)
}

func Menu(w string) {
	//msg_info(w, v)
	msg_raw(log_Menu, w, "", colour.Cyan)

}

func Template(w string) {
	//msg_info(w, v)
	msg_raw(log_Template, w, "", colour.Cyan)

}

func Accessing(w string) {
	//msg_info(w, v)
	msg_raw(log_Accessing, w, "", colour.Green)

}

func Query(w string) {
	//msg_info(w, v)
	if CFG.VerboseMode {
		msg_raw(log_Query, w, Character_Query, colour.White)
	}
}

func Storing(t string, w string) {
	//msg_info(w, v)
	if CFG.VerboseMode {

		msg_raw(log_Storing, t+" "+w, Character_Storing, colour.Yellow)
	}
}

func Result(w string, r string) {
	//msg_info(w, v)
	if CFG.VerboseMode {
		msg_raw(log_Result, w, Character_Query+" = "+colour.Bold+r, colour.White)
	}
}

func Database(w string, v string) {
	//msg_info(w, v)
	if CFG.VerboseMode {
		msg_raw(log_Database, w, v, colour.Green)
	}

}

func StartJob(w string) {
	//msg_info(w, v)
	msg_raw(log_ScheduleStart, w, Character_Bike, colour.Green)

}

func EndJob(w string) {
	//msg_info(w, v)
	msg_raw(log_ScheduleEnd, w, Character_Tick, colour.Green)

}

func Warning(s string) {
	msg_raw(log_Warning, s, Character_Warning, colour.Bold+colour.Yellow)

	//log.Println(colourYellow + "Warning       : " + s + " " + colourReset)
}

func Message(w string, v string) {
	//log.Println(colourReset + "Warning       : " + s + " " + colourReset)
	//output := fmt.Sprintf("%s : %s", w, v)
	//log.Println(colourReset + output + colourReset)
	msg_raw(w, v, "", colour.White)

}

func Error(s string, e error) {
	msg_raw(log_Error, s, e.Error(), colour.Bold+colour.Red)
	log.Fatalln(colour.Bold + colour.Red + e.Error())
}

func Fatal(s string, e error) {

	Error(s, e)
}

func Panic(s string, e error) {
	msg_raw(log_Panic, s, e.Error(), colour.Red)
	log.Panicln(colour.Red + s + e.Error())
}

func Publish(w string, v string) {
	op := v + " " + Character_MapTo + "  " + w + " " + Character_Tick
	//	msg_mux(v + " " + Character_MapTo + "  " + w)
	msg_raw(log_MUX, op, "", colour.Bold+colour.White)
}

func msg_raw(pref string, what string, value string, clr string) {

	op := fmt.Sprintf("%-13s : %s %s", pref, what, value)
	log.Println(clr + op + colour.Reset)
}

func Break() {
	width, _, _ := term.GetSize(0)
	log.Println(colour.Bold + strings.Repeat("-", width-20) + colour.Reset)
	//log.Println("width: ", width)
}

func Header(s string) {
	log.Println(colour.Green + colour.Bold + s + colour.Reset)
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func getConfig() (config Config, err error) {
	// get current os directory path
	pwd, _ := os.Getwd()

	//fmt.Println(pwd)
	viper.AddConfigPath(pwd + "/config/")
	viper.SetConfigName("logs")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = viper.Unmarshal(&config)
	//spew.Dump(config)
	if config.Verbose == "true" {
		config.VerboseMode = true
	} else {
		config.VerboseMode = false
	}
	return
}
