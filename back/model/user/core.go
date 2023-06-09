package user

import "gorm.io/gorm"

// 用户的核心信息
type Core struct {
	gorm.Model
	UserName string `json:"username" gorm:"type:varchar(20);not null" validate:"required,min=2,max=10"`
	Password string `json:"password" gorm:"type:varchar(500);not null" validate:"required,min=8,max=18"`
	Email    string `json:"email" gorm:"type:varchar(500);not null"`
	Role     string `json:"role" gorm:"type:int;DEFAULT:2;not null"`
}

// 添加用户（注册）
func AddUser() {

}
