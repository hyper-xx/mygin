package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/hyper-xx/mygin/handler"
	"github.com/hyper-xx/mygin/model"
	"github.com/hyper-xx/mygin/pkg/auth"
	"github.com/hyper-xx/mygin/pkg/errnum"
	"github.com/hyper-xx/mygin/pkg/token"
)

// @Summary Login generates the authentication token
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /login [post]
func Login(c *gin.Context) {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errnum.ErrBind, nil)
		return
	}

	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errnum.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errnum.ErrPasswordIncorrect, nil)
		return
	}

	//Sign the json web token
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errnum.ErrToken, nil)
		return
	}
	SendResponse(c, nil, model.Token{Token: t})
}
