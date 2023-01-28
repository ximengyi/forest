package middleware

import (
	"errors"
	"fmt"
	"forest/internal/constant"
	"forest/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Accepted = append(c.Accepted, "")
		isMatched := false
		for _, host :=range config.StringMap("http.allow_ip") {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched {
			response.Error(c, constant.InternalErrorCode, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
