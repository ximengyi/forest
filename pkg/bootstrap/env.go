package bootstrap

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"os"
)


var (
	appEnv string //配置环境名 比如：dev prod test
)

func InitEnv() error {

	keys := []string{ "app_env", "mysql"}
	err := config.LoadFlags(keys)
	fmt.Println(config.Data())
	appEnv = config.String("app_env")
	fmt.Println("=========appenv======",appEnv)
	if appEnv == ""{
		appEnv = config.GetEnv("app_env")
	}

    yamlFile := appEnv + ".yaml"
	config.WithOptions(config.ParseEnv)
	// add driver for support yaml content
	config.AddDriver(yamlv3.Driver)
	err = config.LoadFiles(yamlFile)
	if err != nil {
		return err
	}

	//初始化配置文件目录
	fmt.Println("----------------------bootstrap begin----------------------")
	fmt.Printf("[INFO]  appEnv=%s\n", config.String("app_env"))
    if config.Bool("debug_mode"){

		fmt.Printf("config data: \n %#v\n", config.Data())
	}
	fmt.Printf("[INFO] %s\n", " start loading resources.")
       os.Exit(1)
	return nil
}
