package handler

import (
	"github.com/gin-gonic/gin"
	"qiaoyu/api/handler/user"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	u := user.NewUserHandler()

	api := r.Group("/api/v1")
	{
		api.POST("user", u.Register)
	}
	return r
}
