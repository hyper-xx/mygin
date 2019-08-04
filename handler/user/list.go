package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/hyper-xx/mygin/handler"
	"github.com/hyper-xx/mygin/pkg/errnum"
	"github.com/hyper-xx/mygin/service"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errnum.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		// fmt.Println(err.Error())
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
