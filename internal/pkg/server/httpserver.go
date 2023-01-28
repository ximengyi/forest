package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
	"log"
)

//var (
//	HttpSrvHandler *http.Server
//)

func HttpServerRun(r *gin.Engine) {
	gin.SetMode(config.String("debug_mode"))

	log.Println(" [INFO] HttpServerRun:", config.String("http.addr"))
	if err := r.Run(config.String("http.addr")); err != nil {
		log.Fatalln(" [ERROR] HttpServerRun:", config.String("http.addr")," err:", err)
	}

}

//func HttpServerStop() {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
//		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
//	}
//
//	log.Printf(" [INFO] HttpServerStop stopped\n")
//}
