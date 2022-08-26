package router

import (
	"Devops/api/v1"
	"Devops/middleware"
	"Devops/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1").Use(middleware.JwtToken())
	{
		//用户模块

		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		//分类模块
		router.POST("category/add", v1.AddCategory)
		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)

		//文章模块
	}
	public := r.Group("api/v1")
	{
		public.POST("user/add", v1.AddUser)
		public.GET("users", v1.GetUserList)
		public.GET("category", v1.GetCategory)
		public.POST("login", v1.Login)
	}


	r.Run(utils.HttpPort)
}
