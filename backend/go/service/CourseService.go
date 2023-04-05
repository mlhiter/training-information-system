package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建课程信息
*/
func CreateCourse(course *entity.Course) (err error) {
	if err = dao.SqlSession.Create(course).Error; err != nil {
		return err
	}
	return
}

/*
获取课程集合
*/
func GetAllCourse() (courseList []*entity.Course, err error) {
	if err = dao.SqlSession.Find(&courseList).Error; err != nil {
		return nil, err
	}
	return
}

/*
根据id删除课程
*/
func DeleteCourseById(id uint) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.Course{}).Error
	return
}

/*
根据id查询课程
*/
func GetCourseById(id uint) (course *entity.Course, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(course).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新课程信息
*/
func UpdateCourse(course *entity.Course) (err error) {
	err = dao.SqlSession.Save(course).Error
	return
}
