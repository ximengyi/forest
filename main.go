package main

import (
	"forest/internal/pkg/server"
	"forest/internal/router"
	"forest/pkg/bootstrap"
	"log"
	"os"
)

func main() {

	//err := bootstrap.InitModule([]string{"base", "mysql", "redis", "mongodb"})
	err := bootstrap.InitModule([]string{"base", "mysql", "redis"})
	if err != nil {
		log.Println("bootstrap init env fail check it ",err)
		os.Exit(1)
	}
	defer bootstrap.Destroy()

	route := router.InitRouter()
	server.HttpServerRun(route)

}
