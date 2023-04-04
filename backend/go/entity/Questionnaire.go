package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (Questionnaire) TableName() string {
	return "questionnaire"
}

type Questionnaire struct {
	gorm.Model
	CourseId     uint   `json:"courseId"`
	CourseName   string `json:"courseName"`
	StudentId    uint   `json:"studentId"`
	StudentName  string `json:"studentName"`
	Satisfaction int    `json:"satisfaction"`
	Comment      string `json:"comment"`
	Suggestion   string `json:"suggestion"`
}
