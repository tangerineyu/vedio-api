package model

import "gorm.io/gorm"

type UserRelation struct {
	gorm.Model
	UserID   uint `gorm:"not null;uniqueIndex;idx_user_follow"` // 关注者ID
	FollowID uint `gorm:"not null;uniqueIndex;idx_user_follow"` // 被关注者ID
	User     User `gorm:"foreignKey:UserID"`
	Follow   User `gorm:"foreignKey:FollowID"`
}
