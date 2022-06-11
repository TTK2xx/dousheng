package service

import (
	"dousheng/database"
	"dousheng/model"
	"fmt"
	"log"
)

func CreateVideo(video *model.Video) {
	res := database.MySQLDB.Create(&video)
	if res.Error != nil {
		log.Println("Insert video failed!" + res.Error.Error())
	}

}
func GetVideoById(Id int64) (video *model.Video) {
	var v model.Video
	res := database.MySQLDB.Model(&model.Video{}).Where("id = ?", Id).First(&v)
	if res.Error != nil {
		log.Println(res.Error.Error())
	}
	return &v
}
func GetAllVideos() (video []model.Video) {
	var videos []model.Video
	database.MySQLDB.Model(&model.Video{}).Find(&videos)
	fmt.Println("%#v", videos)

	return videos
}
