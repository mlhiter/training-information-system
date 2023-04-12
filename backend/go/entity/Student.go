package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (Student) TableName() string {
	return "student"
}

type Student struct {
	gorm.Model
	StudentId uint   `json:"studentId"`
	CheckTime string `json:"checkTime"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Contact   string `json:"contact"`
	Company   string `json:"company"`
	Post      string `json:"post"`
	Level     string `json:"level"`
}
