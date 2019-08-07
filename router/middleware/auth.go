package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hyper-xx/mygin/handler"
	"github.com/hyper-xx/mygin/pkg/errnum"
	"github.com/hyper-xx/mygin/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errnum.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
	}
}
