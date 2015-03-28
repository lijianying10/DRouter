package conf

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// UsageInfo show how to use drouter
func UsageInfo() {
	fmt.Fprintf(os.Stderr, "DRouter [config_file]")
	flag.PrintDefaults()
	os.Exit(2)
}

func readConfigFile(filePath string) {
	if fp, err := os.Open(filePath); err == nil {
		fpb, _ := ioutil.ReadAll(fp)
		if err = json.Unmarshal(fpb, &Gconf); err != nil {
			fmt.Println("config syntax error: ", err.Error())
			os.Exit(2)
		}
	} else {
		fmt.Println("config file " + filePath + " error")
	}
}

// FlagInit read flags and read config file
func FlagInit() {
	flag.Usage = UsageInfo
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		if _, err := os.Stat("/etc/drouter.conf"); err == nil {
			fmt.Println("config file in the /etc/")
			readConfigFile("/etc/drouter.conf")
			return
		}
		if _, err := os.Stat("./drouter.conf"); err == nil {
			fmt.Println("config file in the runtime floader")
			readConfigFile("./drouter.conf")
			return
		}
		fmt.Println("Please input config file")
		os.Exit(2)
	}
	if _, err := os.Stat(args[0]); err == nil {
		fmt.Println("Using Config file " + args[0])
		readConfigFile(args[0])
		return
	}
	fmt.Println("input file not exist")
	os.Exit(2)

}
