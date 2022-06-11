package model

type Video struct {
	Id            int64  `gorm:"primaryKey autoIncrement" json:"video_id,omitempty"`
	Author        User   `gorm:"-" json:"author"` //默认关联主键
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}
