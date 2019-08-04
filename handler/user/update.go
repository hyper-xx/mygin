package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	. "github.com/hyper-xx/mygin/handler"
	"github.com/hyper-xx/mygin/model"
	"github.com/hyper-xx/mygin/pkg/errnum"
	"github.com/hyper-xx/mygin/util"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errnum.ErrBind, nil)
		return
	}

	u.Id = uint64(userId)

	if err := u.Validate(); err != nil {
		SendResponse(c, errnum.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errnum.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		SendResponse(c, errnum.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
