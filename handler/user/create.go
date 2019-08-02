package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyper-xx/mygin/pkg/errnum"
	"github.com/lexkong/log"
)

//Create a nuew user
func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errnum.ErrBind})
		return
	}

	log.Debugf("username is :[%s],password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		err = errnum.New(errnum.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		log.Errorf(err, "Get an error")
	}

	if errnum.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}

	code, message := errnum.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
