package model

import (
	"Devops/utils/errMsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
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
	data.Password = ScryptPw(data.Password)
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

// EditUser 编辑用户信息只限定密码意外的信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err !=nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errMsg.ERROR
	}
	return errMsg.SUCCESS
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw


}

// CheckLogin 登录验证
func CheckLogin(username, password string) int {
	var user User
	err = db.Where("username = ?", username).First(&user).Error
	if user.ID == 0 {
		return errMsg.ErrorUserNotExist
	}
	if ScryptPw(password) != user.Password {
		return errMsg.ErrorPasswordWrong
	}
	if user.Role != 0 {
		return errMsg.ErrorUserNoRight
	}
	return errMsg.SUCCESS
}