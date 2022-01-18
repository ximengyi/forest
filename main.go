package main

import (
	"forest/bootstrap"
	"forest/log"
	"forest/router"
	"forest/server"
	"os"
)


func main() {


	err := bootstrap.InitModule([]string{"base", "mysql", "redis","mongodb"})
	if err != nil {
		log.Info("bootstrap init env fail check it  ")
		os.Exit(1)
	}
	defer bootstrap.Destroy()
	route := router.InitRouter()
	server.HttpServerRun(route)


}

