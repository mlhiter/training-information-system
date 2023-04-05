package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (User) TableName() string {
	return "user"
}

type User struct {
	gorm.Model
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Role      string `json:"role"`
}
