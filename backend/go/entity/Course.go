package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (Course) TableName() string {
	return "course"
}

type Course struct {
	gorm.Model
	Name             string  `json:"name"`
	ClassTime        string  `json:"classTime"`
	PlaceOfClass     string  `json:"placeOfClass"`
	Price            float64 `json:"price"`
	SelectedNumber   uint
	Income			 float64
	LecturerId       uint    `json:"lecturerId"`
	ExecutorId       uint    `json:"executorId"`
	EmailContentId   uint    `json:"EmailContentId"`
	TrainingNoticeId uint    `json:"TrainingNoticeId"`
}
