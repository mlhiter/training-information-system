package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (CheckInRecord) TableName() string {
	return "check_in_record"
}

type CheckInRecord struct {
	gorm.Model
	CourseId  uint   `json:"courseId"`
	StudentId uint   `json:"studentId"`
	CheckTime string `json:"checkTime"`
}
