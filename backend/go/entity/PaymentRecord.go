package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (PaymentRecord) TableName() string {
	return "payment_record"
}

type PaymentRecord struct {
	gorm.Model
	CourseId      uint    `json:"courseId"`
	CourseName    string  `json:"courseName"`
	StudentId     uint    `json:"studentId"`
	StudentName   string  `json:"studentName"`
	PaymentAmount float64 `json:"paymentAmount"`
	PaymentTime   string  `json:"paymentTime"`
	Payee         string  `json:"payee"`
	PayAccount    string  `json:"payAccount"`
}
