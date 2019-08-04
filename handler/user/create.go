package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/hyper-xx/mygin/handler"
	"github.com/hyper-xx/mygin/model"
	"github.com/hyper-xx/mygin/pkg/errnum"
	"github.com/hyper-xx/mygin/util"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

//Create a nuew user
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errnum.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	//Validate the date
	if err := u.Validate(); err != nil {
		SendResponse(c, errnum.ErrValidation, nil)
		return
	}

	//Encrypt the user password
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errnum.ErrEncrypt, nil)
		return
	}

	//Insert the user to the database
	if err := u.Create(); err != nil {
		//fmt.Printf(err.Error())
		SendResponse(c, errnum.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	//Show the user info
	SendResponse(c, nil, rsp)
}
