package middleware

import (
	"errors"
	"fmt"
	"forest/config"
	"forest/constant"
	"forest/response"
	"github.com/gin-gonic/gin"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Accepted = append(c.Accepted, "")
		isMatched := false
		for _, host := range config.GetStringSliceConf("base.http.allow_ip") {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched{
			response.Error(c, constant.InternalErrorCode, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
