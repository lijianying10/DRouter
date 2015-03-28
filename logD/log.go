package logD

import (
	"DRouter/conf"
	"fmt"
	"log"
	"os"
)

// logger
var Logger *log.Logger

var outFile *os.File

func init() {
	if Logger == nil {
		logFileName := conf.Gconf["accesslog"].(string)
		var err error
		if outFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); err != nil {
			fmt.Println("proxy log ERROR : ", err.Error())
		}
		Logger = log.New(outFile, "", log.Ldate|log.Ltime|log.Llongfile)
	}
}

// CloseLogFile function: close logger file
func CloseLogFile() {
	outFile.Close()
}
