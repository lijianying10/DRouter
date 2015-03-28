package main

import (
	"DRouter/Router"
	"DRouter/conf"
	"DRouter/logD"
	"net/http"
)

func main() {
	logD.Logger.Println("Welcome to DRouter for reverse proxy and static file")
	Router.Bindings()
	if err := http.ListenAndServe(":"+(conf.Gconf["port"]).(string), nil); err != nil {
		logD.Logger.Println("http server Error : ", err.Error())
	}

}
