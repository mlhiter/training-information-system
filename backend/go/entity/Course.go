package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (Course) TableName() string {
	return "course"
}

type Course struct {
	gorm.Model
	Name             string  `json:"name"`
	Time             string  `json:"time"`
	Place            string  `json:"place"`
	Price            float64 `json:"price"`
	SelectedNumber   uint    `json:"selectedNumber"`
	Income			 float64 `json:"income"`
	LecturerId       uint    `json:"lecturerId"`
	Lecturer         string  `json:"lecturer"`
}
