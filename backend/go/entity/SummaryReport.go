package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (SummaryReport) TableName() string {
	return "summary_report"
}

type SummaryReport struct {
	gorm.Model
}
