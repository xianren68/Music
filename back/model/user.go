package model

import (
	"time"

	"gorm.io/gorm"
)

// 创建用户数据库
type User struct {
	gorm.Model
	UserName    string    `json:"username" gorm:"type:varchar(50); not null"`
	Password    string    `json:"password" gorm:"type:varchar(255); not null"`
	NickName    string    `json:"nickname" gorm:"type:varchar(50); not null"`
	Avatar      string    `json:"avatar" gorm:"type:varchar(255)"`
	Emial       string    `json:"emial" gorm:"type:varchar(50); not null"`
	Line        string    `json:"line" gorm:"type:varchar(255)"`
	Create_time time.Time `json:"create_time" gorm:"type:datetime; not null"`
}
