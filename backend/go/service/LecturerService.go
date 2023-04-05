package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建课程信息
*/
func CreateLecturer(lecturer *entity.Lecturer) (err error) {
	if err = dao.SqlSession.Create(lecturer).Error; err != nil {
		return err
	}
	return
}

/*
获取课程集合
*/
func GetAllLecturer() (lecturerList []*entity.Lecturer, err error) {
	if err = dao.SqlSession.Find(&lecturerList).Error; err != nil {
		return nil, err
	}
	return
}

/*
根据id删除课程
*/
func DeleteLecturerById(id uint) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.Lecturer{}).Error
	return
}

/*
根据id查询课程
*/
func GetLecturerById(id uint) (lecturer *entity.Lecturer, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(lecturer).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新课程信息
*/
func UpdateLecturer(lecturer *entity.Lecturer) (err error) {
	err = dao.SqlSession.Save(lecturer).Error
	return
}
