package model

import (
	"Devops/utils/errMsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username"`
	Password string `gorm:"type:varchar(20);not null " json:"password"`
	Role int `gorm:"type:int" json:"role"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errMsg.ErrorUsernameUsed
	}
	return errMsg.SUCCESS
}

// CreateRUser AddUser 新增用户
func CreateRUser(data *User) (code int) {
	err := db.Create(data).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var user []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return user
}