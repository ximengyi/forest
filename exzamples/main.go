package main

import (
	"context"
	"fmt"
	"forest/bootstrap"
	"forest/config"
	"forest/db"
	"forest/log"
	"os"
	"time"
)

func main() {

	//err :=bootstrap.Init("./conf/local/")
	err := bootstrap.InitModule( []string{"base", "mysql", "redis"})

	if err != nil {
		log.Info("bootstrap init env fail check it  ")
		os.Exit(1)
	}
	defer bootstrap.Destroy()
	rdb, err := db.RedisConnFactory(0)
	ctx := context.Background()
	rdb.Set(ctx, "meng", "hello meng", 360*time.Second)

	conf := config.GetStringConf("base.base.master_secret")
	fmt.Println(conf)
	fmt.Println(config.ConfBase)
	log.Info("iam a test log")

   //优雅关停服务，其实一点也不优雅，采用这种方式当nginx 反向代理服务时c.client 无法获取真实ip
	//quit := make(chan os.Signal)
	//signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//
	//server.HttpServerStop()
}
