package middleware

import (
	"errors"

	"forest/internal/constant"
	"forest/internal/response"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if name, ok := session.Get("openid").(string); !ok || name == "" {
			response.Error(c, constant.InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
