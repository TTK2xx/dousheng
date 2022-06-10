package service

import (
	"dousheng/database"
	"dousheng/model"
	"log"
)

func CreateVideo(video *model.Video) {
	res := database.MySQLDB.Create(&video)
	if res.Error != nil {
		log.Println("Insert video failed!" + res.Error.Error())
	}

}
