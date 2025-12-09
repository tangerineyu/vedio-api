package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"not null;index"`
	User    User   `gorm:"foreignKey:UserID"`
	VideoID uint   `gorm:"not null;index"`
	Video   Video  `gorm:"foreignKey:VideoID"`
	Content string `gorm:"type:text;not null"`
	//父评论id，默认为0,表示这是对视频的评论
	ParentID uint `gorm:"default:0;index"`
}
type UserFavorite struct {
	gorm.Model
	UserID  uint  `gorm:"not null;uniqueIndex;idx_user_video"`
	VideoID uint  `gorm:"not null;uniqueIndex;idx_user_video"`
	User    User  `gorm:"foreignKey:UserID"`
	Video   Video `gorm:"foreignKey:VideoID"`
}
