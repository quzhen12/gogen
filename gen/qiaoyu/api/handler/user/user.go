package user

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/quzhen12/plugins/res"
	"qiaoyu/pkg/user"
)

type userHandler struct {
	srv user.Service
}

func NewUserHandler() *userHandler {
	return &userHandler{
		srv: user.NewService(),
	}
}

type register struct {
	Email    string
	Password string
}

func (u *userHandler) Register(c *gin.Context) {
	var r register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		res.Json(c, "", err)
	}
	ret, err := u.srv.Register(context.Background(), r.Email, r.Password)
	res.Json(c, ret, err)
}
