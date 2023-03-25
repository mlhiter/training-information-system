package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (Lecturer) TableName() string {
	return "executor"
}

type Lecturer struct {
	gorm.Model
	Name              string `json:"name"`
	ProfessionalTitle string `json:"professionalTitle"`
	FieldOfExpertise  string `json:"fieldOfExpertise"`
	EmailAccount      string `json:"emailAccount"`
	Tel               string `json:"tel"`
}
