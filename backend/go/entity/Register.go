package entity

import "github.com/jinzhu/gorm"

// TableName 数据库表明自定义，默认为model的复数形式
func (Register) TableName() string {
	return "register"
}

type Register struct {
	gorm.Model
	Name           string `json:"name"`
	Sex            string `json:"sex"`
	CompanyName    string `json:"companyName"`
	OperatingPost  string `json:"operatingPost"`
	TechnicalLevel string `json:"technicalLevel"`
	ContactInfo    string `json:"contactInfo"`
}
