package bootstrap

import (
	"flag"
	"fmt"
	"forest/config"
	"os"
)

type Env struct{
	Env    string `mapstructure:"env"`
	Config map[string]interface{}
}

// 解析配置文件目录
//
// 配置文件必须放到一个文件夹中
// 如：config=conf/dev/base.json 	ConfEnvPath=conf/dev	ConfEnv=dev
// 如：config=conf/base.json		ConfEnvPath=conf		ConfEnv=conf

var (
	appEnv string  //配置环境名 比如：dev prod test
    //配置文件夹
	//env *Env

	)

func InitEnv() error{
	//优先从命令行读取配置，命令行没有读到，再从环境变量读取，环境变量没有读取到从env 读取
	defaultEnv := ""
	envConf := flag.String("APP_ENV", defaultEnv, "input env file like local|stage|uat|production")
	flag.Parse()
	if *envConf == "" {
		appEnv = os.Getenv("APP_ENV")
	}else{
		appEnv = *envConf
	}

	if appEnv == ""{
		flag.Usage()
		fmt.Println("parse env fail check your env config")
		os.Exit(1)
	}
	//初始化配置文件目录
	config.ConfEnvPath = config.ConfEnvPath + appEnv
	fmt.Println("----------------------bootstrap begin----------------------")
	fmt.Printf("[INFO]  appEnv=%s\n", appEnv)
	fmt.Printf("[INFO]  configEnvPath=%s\n", config.ConfEnvPath)
	fmt.Printf("[INFO] %s\n", " start loading resources.")

	//初始化config 目录配置文件
	if err := config.InitViperConf(); err != nil {
		return err
	}
     return nil
}


//获取配置环境名
func GetAppEnv() string{
	return appEnv
}
