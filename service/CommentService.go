package service

import (
	"dousheng/database"
	"dousheng/model"
	"log"
)

func GetCommentByVideoID(vid int64) (comment model.Comment) {
	var u model.Comment
	res := database.MySQLDB.Model(&model.Comment{}).Where("video_id = ?", vid).First(&u)
	if res.Error != nil {
		log.Println(res.Error.Error())
	}
	return u
}

func CreateComment(comment *model.Comment) {
	res := database.MySQLDB.Create(&comment)
	if res.Error != nil {
		log.Println("Insert user failed!" + res.Error.Error())
	}
	//return user
}

func DeleteComment(cid int64) {
	database.MySQLDB.Delete(&model.Comment{}, cid)
	//return user
}
