package server

import (
	"context"
	"forest/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun(r *gin.Engine) {
	gin.SetMode(config.ConfBase.DebugMode)
	//r := InitRouter()
	//HttpSrvHandler = &http.Server{
	//	Addr:           bootstrap.GetStringConf("base.http.addr"),
	//	Handler:        r,
	//	ReadTimeout:    time.Duration(bootstrap.GetIntConf("base.http.read_timeout")) * time.Second,
	//	WriteTimeout:   time.Duration(bootstrap.GetIntConf("base.http.write_timeout")) * time.Second,
	//	MaxHeaderBytes: 1 << uint(bootstrap.GetIntConf("base.http.max_header_bytes")),
	//}

	//go func() {
	log.Printf(" [INFO] HttpServerRun:%s\n", config.GetStringConf("base.http.addr"))
	if err := r.Run(config.GetStringConf("base.http.addr")); err != nil {
		log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", config.GetStringConf("base.http.addr"), err)
	}
	//}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}

	log.Printf(" [INFO] HttpServerStop stopped\n")
}
