package bootstrap

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"log"
)


var (
	appEnv string //env likes: dev prod test
)

func InitEnv() error {

	keys := []string{ "app_env", "mysql"}
	err := config.LoadFlags(keys)
	appEnv = config.String("app_env")
	if appEnv == ""{
		appEnv = config.GetEnv("app_env")
	}

	yamlFile := appEnv + ".yaml"
	// add driver for support yaml content
	config.AddDriver(yamlv3.Driver)
	err = config.LoadFiles(yamlFile)
	if err != nil {
		return err
	}

	//init config dir
	log.Println("----------------------bootstrap begin----------------------")
	log.Println("[INFO]  appEnv=", config.String("app_env"))
	if config.Bool("debug_mode"){
		log.Println("config data:", config.Data())
	}
	log.Println("[INFO]", " start loading resources.")
	return nil
}