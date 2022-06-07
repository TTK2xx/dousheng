package controller

import "dousheng/model"

var DemoUser = model.User{
	ID:       100,
	Username: "admin",
	Password: "123456",
}

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
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

type Video struct {
	Id            int64      `gorm:"primaryKey autoIncrement" json:"id,omitempty"`
	Author        model.User `json:"author"`
	PlayUrl       string     `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string     `json:"cover_url,omitempty"`
	FavoriteCount int64      `json:"favorite_count,omitempty"`
	CommentCount  int64      `json:"comment_count,omitempty"`
	IsFavorite    bool       `json:"is_favorite,omitempty"`
	Title         string     `json:"title,omitempty"`
}
