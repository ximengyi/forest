package main

import (
	"forest/internal/pkg/server"
	"forest/internal/router"
	"forest/pkg/bootstrap"
	"forest/pkg/log"
	"os"
)

func main() {

	//err := bootstrap.InitModule([]string{"base", "mysql", "redis", "mongodb"})
	err := bootstrap.InitModule([]string{"base", "mysql", "redis"})
	//err := bootstrap.InitModule([]string{"base", "redis"})
	if err != nil {
		log.Info("bootstrap init env fail check it  ")
		os.Exit(1)
	}
	defer bootstrap.Destroy()
	route := router.InitRouter()
	server.HttpServerRun(route)

}
