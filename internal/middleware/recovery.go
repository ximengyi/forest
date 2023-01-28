package middleware

import (
	"errors"
	"fmt"
	"forest/internal/response"
	"forest/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
	"runtime/debug"
)

// RecoveryMiddleware捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//先做一下日志记录
				fmt.Println(string(debug.Stack()))
				log.Infof("_com_panic", map[string]interface{}{
					"error": fmt.Sprint(err),
					"stack": string(debug.Stack()),
				})

				if config.String("debug_mode") != "debug" {
					response.Error(c, 500, errors.New("内部错误"))
					return
				} else {
					response.Error(c, 500, errors.New(fmt.Sprint(err)))
					return
				}
			}
		}()
		c.Next()
	}
}
