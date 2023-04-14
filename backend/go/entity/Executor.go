package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (Executor) TableName() string {
	return "executor"
}

type Executor struct {
	gorm.Model
	Name        string   `json:"name"`
	Course      string   `json:"course"`
	Condition   string   `json:"condition"`
}
