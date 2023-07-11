package model

import "gorm.io/gorm"

// 歌曲分类
type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(50); not null"`
}
