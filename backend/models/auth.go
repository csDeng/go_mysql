package models

import "fmt"

type Auth struct {
	// 转码成gorm的时候，ID改成primary_key, 转码成json的时候,叫id
	Model
	Username string `json:"username" validate:"required" gorm:"not null;type:string;size:100"`
	Password string `json:"password" validate:"required" gorm:"not null;type:string;size:256"`
	Level    string `json:"level" gorm:"default:0;comment:0是普通用户,1是管理员;type:string;size:1"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	return auth.ID > 0
}

func CheckAdmin(username, password string) (user Auth) {
	db.Where("username = ? AND password = ?", username, password).First(&user)
	return
}

func UsernameIsExisted(username string) bool {
	// 获取第一条匹配的记录
	var auth Auth
	db.Select("id").Where("username = ?", username).First(&auth)
	return auth.ID > 0

}
func UserIsExistedBYId(id int) bool {
	var auth Auth
	db.Select("id").Where("id = ?", id).First(&auth)
	return auth.ID > 0
}

func AddUser(username, password string) bool {
	db.Create(&Auth{
		Username: username,
		Password: password,
	})
	return true
}

func GetUserById(id int) (user Auth) {
	db.Where("id = ?", id).First(&user)
	return
}

func GetUsers(pageNum int,
	pageSize int,
	maps interface{}) (users []Auth) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users)
	return
}

func GetUserTotal(maps interface{}) (count int) {
	db.Model(&Auth{}).Where(maps).Count(&count)
	return
}

func EditUser(id int, data interface{}) bool {
	fmt.Println("model edit=>", data)
	db.Model(&Auth{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteUser(id int) bool {
	db.Where("id = ?", id).Delete(&Auth{})
	return true
}
