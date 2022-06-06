package service

import (
	"dousheng/database"
	"dousheng/model"
	"log"
)

func GetUserByUsername(username string) (user *model.User, err error) {
	var u model.User
	res := database.MySQLDB.Model(&model.User{}).Where("username = ?", username).First(&u)
	if res.Error != nil {
		log.Println(res.Error.Error())
	}
	return &u, res.Error
}

func IsUserExisted(username string) bool {
	var user model.User
	res := database.MySQLDB.Where("username=?", username).First(&user)
	if res.RowsAffected >= 1 {
		return true
	}
	return false
}

func CreateUser(user *model.User) {
	res := database.MySQLDB.Create(&user)
	if res.Error != nil {
		log.Println("Insert user failed!" + res.Error.Error())
	}
	//return user
}
