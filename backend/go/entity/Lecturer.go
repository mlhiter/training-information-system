package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (Lecturer) TableName() string {
	return "lecturer"
}

type Lecturer struct {
	gorm.Model
	Name           string  `json:"name"`
	Profession     string  `json:"profession"`
	Field          string  `json:"field"`
	Email          string  `json:"email"`
	Telephone      string  `json:"telephone"`
}
