package v1

import (
	"Devops/model"
	"Devops/utils/errMsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//查询分类是否存在

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errMsg.SUCCESS {
		model.CreateCategory(&data)
	}
	if code == errMsg.ErrorCateNameUsed {
		code = errMsg.ErrorCateNameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"msg": errMsg.GetErrMsg(code),
	})
}

// todo 查询单个分类下的所有文章

// GetCategory 查询分类列表
func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCategory(pageSize, pageNum)
	code = errMsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"msg": errMsg.GetErrMsg(code),
	})
}

// EditCategory 编辑分类
func EditCategory(c *gin.Context) {
	var cate model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&cate)
	code = model.CheckCategory(cate.Name)
	if code == errMsg.SUCCESS {
		model.EditCategory(id, &cate)
	}
	if code == errMsg.ErrorCateNameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg": errMsg.GetErrMsg(code),
	})

}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCategory(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg": errMsg.GetErrMsg(code),
	})
}