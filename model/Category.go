package model

import (
	"Devops/utils/errMsg"
	"gorm.io/gorm"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null " json:"name"`
}

// CheckCategory 查询分类否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errMsg.ErrorCateNameUsed
	}
	return errMsg.SUCCESS
}

// CreateCategory 新增分类
func CreateCategory(data *Category) (code int) {
	err := db.Create(data).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// todo 查询分类下的所有文章

// GetCategory 查询分类列表
func GetCategory(pageSize int, pageNum int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return cate
}

// EditCategory 编辑分类
func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err !=nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}