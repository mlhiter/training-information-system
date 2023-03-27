package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建Student信息
*/
func CreateStudent(student *entity.Student) (err error) {
	if err = dao.SqlSession.Create(student).Error; err != nil {
		return err
	}
	return
}
