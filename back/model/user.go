package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"type:varchar(20);not null" validate:"required,min=2,max=10"`
	Password string `json:"password" gorm:"type:varchar(500);not null" validate:"required,min=8,max=18"`
	Email    string `json:"email" gorm:"type:varchar(500);not null"`
	Phone    string `json:"phone" gorm:"type:bigint;not null"`
	Role     string `json:"role" gorm:"type:int;DEFAULT:2;not null"`
}
