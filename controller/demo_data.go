package controller

import "dousheng/model"

var DemoUser = model.User{
	ID:       100,
	Username: "admin",
	Password: "123456",
}

var DemoVideos = []model.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://10.0.2.2:8080/static/bear.mp4",
		CoverUrl:      "http://10.0.2.2:8080/static/hcl_1.jpg",
		FavoriteCount: 2,
		CommentCount:  3,
		IsFavorite:    false,
		Title:         "hcl‘s title1",
	},
	{
		Id:            2,
		Author:        DemoUser,
		PlayUrl:       "http://10.0.2.2:8080/static/bear.mp4",
		CoverUrl:      "http://10.0.2.2:8080/static/hcl_1.jpg",
		FavoriteCount: 44,
		CommentCount:  55,
		IsFavorite:    false,
		Title:         "hcl‘s title2",
	},
}
