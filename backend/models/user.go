package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required" gorm:"type:varchar(100);not null;"`
	Password string `json:"password" validate:"required" gorm:"type:varchar(255);not null;"`
	Level    string `json:"level" gorm:"type:char(1);default:'0';comment:'0是普通用户,1是管理员';"`
}

func CheckUser(username, password string) bool {
	var user User
	db.Select("id").Where(User{Username: username, Password: password}).First(&user)
	return user.ID > 0
}

func CheckAdmin(username, password string) (user User) {
	db.Where("username = ? AND password = ?", username, password).First(&user)
	return
}

func UsernameIsExisted(username string) bool {
	// 获取第一条匹配的记录
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	return user.ID > 0

}
func UserIsExistedBYId(id int) bool {
	var user User
	db.Select("id").Where("id = ?", id).First(&user)
	return user.ID > 0
}

func AddUser(username, password string) bool {
	db.Create(&User{
		Username: username,
		Password: password,
	})
	return true
}

func GetUserById(id int) (user User) {
	db.Where("id = ?", id).First(&user)
	return
}

func GetUsers(pageNum int,
	pageSize int,
	maps interface{}) (users []User) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users)
	return
}

func GetUserTotal(maps interface{}) (count int64) {
	db.Model(&User{}).Where(maps).Count(&count)
	return
}

func EditUser(id int, data interface{}) bool {
	fmt.Println("model edit=>", data)
	db.Model(&User{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteUser(id int) bool {
	db.Where("id = ?", id).Delete(&User{})
	return true
}
