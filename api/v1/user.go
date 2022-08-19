package v1

import (
	"Devops/model"
	"Devops/utils/errMsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// UserEXIST 查询用户是否存在
func UserEXIST() {
	// todo
}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	code = model.CheckUser(user.Username)
	if code == errMsg.SUCCESS {
		model.CreateRUser(&user)
	}
	if code == errMsg.ErrorUsernameUsed {
		code = errMsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": user,
		"msg": errMsg.GetErrMsg(code),
	})
}
//查询单个用户

// GetUserList 查询用户列表
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	code = errMsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"msg": errMsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {

}