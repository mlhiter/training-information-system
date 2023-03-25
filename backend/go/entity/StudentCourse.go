package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (StudentCourse) TableName() string {
	return "student_course"
}

type StudentCourse struct {
	gorm.Model
	CourseId  uint `json:"courseId"`
	StudentId uint `json:"studentId"`
}
