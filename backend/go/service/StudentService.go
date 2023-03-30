package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建学生信息
*/
func CreateStudent(student *entity.Student) (err error) {
	if err = dao.SqlSession.Create(student).Error; err != nil {
		return err
	}
	return
}

/*
获取学生集合
*/
func GetAllStudent() (studentList []*entity.Student, err error) {
	if err = dao.SqlSession.Find(&studentList).Error; err != nil {
		return nil, err
	}
	return
}

/*
根据id删除学生
*/
func DeleteStudentById(id uint) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.Student{}).Error
	return
}

/*
根据id查询学生
*/
func GetStudentById(id uint) (student *entity.Student, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(student).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新学生信息
*/
func UpdateStudent(student *entity.Student) (err error) {
	err = dao.SqlSession.Save(student).Error
	return
}
