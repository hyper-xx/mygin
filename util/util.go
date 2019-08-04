package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

//GenShortId uniq id
func GenShortId() (string, error) {
	return shortid.Generate()
}

//Get request uniq id
func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}