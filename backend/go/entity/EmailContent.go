package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (EmailContent) TableName() string {
	return "email_content"
}

type EmailContent struct {
	gorm.Model
	content string `json:"content""`
}
