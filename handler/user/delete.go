package user

import (
	"strconv"

	"github.com/hyper-xx/mygin/pkg/errnum"

	"github.com/hyper-xx/mygin/model"

	"github.com/gin-gonic/gin"
	. "github.com/hyper-xx/mygin/handler"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errnum.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
