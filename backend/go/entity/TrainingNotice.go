package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (TrainingNotice) TableName() string {
	return "training_notice"
}

type TrainingNotice struct {
	gorm.Model
	Notice string `json:"notice"`
}
