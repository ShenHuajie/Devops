package v1

import (
	"Devops/middleware"
	"Devops/model"
	"Devops/utils/errMsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	var data model.User
	ctx.ShouldBindJSON(&data)

	var token string
	var code int

	code = model.CheckLogin(data.Username, data.Password)
	if code == errMsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg": errMsg.GetErrMsg(code),
		"token": token,
	})
}