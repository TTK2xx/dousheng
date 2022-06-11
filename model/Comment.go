package model

// 拉取评论返回的信息

type Comment struct {
	VideoID     int64       `json:"video_id"`
	Content     string      `json:"content"`                                         // 评论内容
	CreateDate  string      `json:"create_date"`                                     // 评论发布日期，格式 mm-dd
	ID          int64       `gorm:"primaryKey autoIncrement" json:"id"`              // 评论id
	CommentUser CommentUser `json:"comment_user" gorm:"foreignKey:id;references:ID"` // 评论用户信息
}

// 评论中使用的用户信息

type CommentUser struct {
	FollowCount   int64  `json:"follow_count"`                       // 关注总数
	FollowerCount int64  `json:"follower_count"`                     // 粉丝总数
	ID            int64  `gorm:"primaryKey autoIncrement" json:"id"` // 用户id
	IsFollow      bool   `json:"is_follow"`                          // true-已关注，false-未关注
	Name          string `json:"name"`                               // 用户名称
}
