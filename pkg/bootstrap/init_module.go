package bootstrap

import (
	"fmt"
	"forest/pkg/log"
	"forest/pkg/mysql"
	"forest/pkg/redis"
	"forest/pkg/utils"
	"github.com/gookit/config/v2"
	"time"
)

var TimeLocation *time.Location
var TimeFormat = "2006-01-02 15:04:05"

//var DateFormat = "2006-01-02"

//公共初始化函数：支持两种方式设置配置文件
//
//函数传入配置文件 Init("./conf/dev/")
//如果配置文件为空，会从命令行中读取 	  -config conf/dev/
//func Init(configPath string) error {
//	return InitModule(configPath, []string{"base", "mysql", "redis"})
//}

//模块初始化

func InitModule(modules []string) error {
	err := InitEnv()
	if err != nil {
		return err
	}

	opts := &log.Options{

		Level:             config.String("log.level"),
		Format:            config.String("log.format"),
		EnableColor:       config.Bool("log.enable_color"),
		DisableCaller:     config.Bool("log.disable_caller"),
		Development:       config.Bool("log.development"),
		DisableStacktrace: config.Bool("log.disable_stacktrace"),
		OutputPaths:       config.Strings("log.OutputPaths"),
		ErrorOutputPaths:  config.Strings("log.ErrorOutputPaths"),
	}

	// 初始化全局logger
	log.Init(opts)
	defer log.Flush()

     if utils.InArrayString("redis", modules) {
		 if err = redis.InitRedisConf(config.String("redis.addr"),config.String("redis.password")); err != nil {
			 log.Fatalw("[ERROR] InitRedisDBPool: " , err.Error())
			 return err
		 }
	 }

	//初始化MySQL连接池
	if utils.InArrayString("mysql", modules) {
		MysqlConf := &mysql.Conf{}
		err = config.BindStruct("mysql", MysqlConf)
		fmt.Println("=========bind", MysqlConf)
		if err != nil {
			log.Fatalw("[ERROR] InitMysqlDBPool: ", err.Error())
			return err
		}
		mysql.InitMysqlPool(MysqlConf)
	}

	log.Info("[INFO] success loading resources.")
	log.Info("----------------------bootstrap end------------------------")
	return nil

}

//公共销毁函数
func Destroy() {
	fmt.Println("----------------------destroy resources---------------------")
	fmt.Printf("[INFO] %s\n", " start destroy resources.")
	fmt.Printf("[INFO] %s\n", " success destroy resources.")
}
