package user

import "gorm.io/gorm"

// 用户信息
type Info struct {
	gorm.Model
	Core     Core    `gorm:"foreignkey:Cid"`
	NickName string  `json:"nickName" gorm:"type:varchar(20);"`
	Avatar   string  `json:"avatar" gorm:"type:varchar(20);"`
	Line     string  `json:"line" gorm:"type:varchar(500);"`
	Duration float64 `json:"duration" gorm:"type:bigint;"`
	Cid      uint    `gorm:"type:int;not null" json:"cid"`
}
