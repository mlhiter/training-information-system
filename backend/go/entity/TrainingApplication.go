package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (TrainingApplication) TableName() string {
	return "training_application"
}

type TrainingApplication struct {
	gorm.Model
	TrainingDate    string `json:"name"`
	TrainingContext string `json:"sex"`
	NumberOfPeople  int    `json:"companyName"`
	Applicant       string `json:"applicant"`
}
