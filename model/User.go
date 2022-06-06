package model

type User struct {
	ID       int64  `gorm:"primaryKey autoIncrement" json:"user_id"`
	Username string `gorm:"uniqueIndex type:varchar(32)" json:"username"` // required，最长32个字符
	Password string `gorm:"type:varchar(32)" json:"password"`             // required，最长32个字符
}
