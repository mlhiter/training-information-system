package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (TrainingApplication) TableName() string {
	return "training_application"
}

type TrainingApplication struct {
	gorm.Model
	Applicant       string `json:"applicant"`
	TrainingDate    string `json:"trainingDate"`
	TrainingContext string `json:"trainingContext"`
	NumberOfPeople  int    `json:"numberOfPeople"`
	Status          string `json:"status"`
}
