package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" valdate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" valdate:"required,min=6,max=120" label:"密码"`
	Role     string `gorm:"type:int;default:2" json:"role" valdate:"required,gte=2" label:"角色码"`
}
