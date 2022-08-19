package router

import (
	"Devops/api/v1"
	"Devops/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		//用户模块
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUserList)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.EditUser)

		//分类模块


		//文章模块
	}

	r.Run(utils.HttpPort)
}
