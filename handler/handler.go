package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyper-xx/mygin/pkg/errnum"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//ReWrite response func with errcode
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errnum.DecodeErr(err)

	c.JSON(http.StatusOK, Response{Code: code, Message: message, Data: data})
}
