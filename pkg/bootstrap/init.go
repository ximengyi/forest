package bootstrap

import (
	"fmt"
	"forest/pkg/config"
	"forest/pkg/db"
	"forest/pkg/log"
	"forest/pkg/utils"
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
	// 加载base配置
	if utils.InArrayString("base", modules) {
		if err := config.InitBaseConf(config.GetConfPath("base")); err != nil {
			fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitBaseConf:"+err.Error())
		}
	}

	opts := &log.Options{
		Level:             config.ConfBase.Log.Level,
		Format:            config.ConfBase.Log.Format,
		EnableColor:       config.ConfBase.Log.EnableColor,
		DisableCaller:     config.ConfBase.Log.DisableCaller,
		Development:       config.ConfBase.Log.Development,
		DisableStacktrace: config.ConfBase.Log.DisableStacktrace,
		OutputPaths:       config.ConfBase.Log.OutputPaths,
		ErrorOutputPaths:  config.ConfBase.Log.ErrorOutputPaths,
	}

	// 初始化全局logger
	log.Init(opts)
	defer log.Flush()

	// 加载mysql配置并初始化实例
	if utils.InArrayString("mysql", modules) {
		mysqlPath := config.GetConfPath("mysql")
		MysqlConfMap := &db.MysqlMapConf{}
		err = config.ParseConfig(mysqlPath, MysqlConfMap)
		if err != nil {
			return err
		}
		if err = db.InitMysqlPool(MysqlConfMap); err != nil {
			fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitMysqlPool:"+err.Error())
			return err
		}
	}

	if utils.InArrayString("mongodb", modules) {
		mongodbPath := config.GetConfPath("mongodb")
		MongodbConf := &db.MongodbConf{}
		err = config.ParseKeyConfig("mongodb", mongodbPath, MongodbConf)
		if err != nil {
			return err
		}
		if err = db.InitMongodbPool(MongodbConf); err != nil {
			fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitMongodbPool:"+err.Error())
			return err
		}
	}

	// 加载redis配置并初始化实例
	if utils.InArrayString("redis", modules) {

		if err = db.InitRedisConf(config.ConfBase.RedisConfig.Addr, config.ConfBase.RedisConfig.Password); err != nil {
			fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitRedisDBPool:"+err.Error())
			return err
		}
	}

	// 设置时区
	if location, err := time.LoadLocation(config.ConfBase.TimeLocation); err != nil {
		return err

	} else {
		TimeLocation = location
	}

	fmt.Printf("[INFO] %s\n", " success loading resources.")
	fmt.Println("----------------------bootstrap end------------------------")

	return nil
}

//公共销毁函数
func Destroy() {
	fmt.Println("----------------------destroy resources---------------------")
	fmt.Printf("[INFO] %s\n", " start destroy resources.")
	fmt.Printf("[INFO] %s\n", " success destroy resources.")
}
