package model

import (
	"time"

	"gorm.io/gorm"
)

// 歌曲信息
type Music struct {
	gorm.Model
	Category      Category  `gorm:"foreignkey:Cid"`
	User          User      `gorm:"foreignkey:User_id"`
	Title         string    `json:"title" gorm:"type:varchar(255); not null"`
	ArtList       string    `json:"artlist" gorm:"type:varchar(255); not null"`
	Album         string    `json:"album" gorm:"type:varchar(255)"` // 专辑
	Cover         string    `json:"cover" gorm:"type:varchar(255)"`
	Url           string    `json:"url" gorm:"type:varchar(255); not null"`
	Lyrics        string    `json:"lyrics" gorm:"type:text"` // 歌词
	Create_time   time.Time `json:"create_time" gorm:"type:datetime; not null"`
	User_id       string    `json:"user_id" gorm:"type:bigint; not null"` // 上传用户
	Like_count    uint      `json:"like_count" gorm:"type:bigint"`
	Comment_count uint      `json:"comment_count" gorm:"type:bigint"`
	Tag           string    `json:"tag" gorm:"type:varchar(255);"`
	Cid           uint      `json:"cid" gorm:"type:bigint; not null" `
}
