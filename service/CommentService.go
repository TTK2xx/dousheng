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
