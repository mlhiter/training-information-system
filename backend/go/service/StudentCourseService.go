package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建课程信息
*/
func CreateStudentCourse(studentCourse *entity.StudentCourse) (err error) {
	if err = dao.SqlSession.Create(studentCourse).Error; err != nil {
		return err
	}
	return
}

/*
获取课程集合
*/
func GetAllStudentCourse() (studentCourseList []*entity.StudentCourse, err error) {
	if err = dao.SqlSession.Find(&studentCourseList).Error; err != nil {
		return nil, err
	}
	return
}

func GetStudentCourseListByCourseName(courseId uint) (studentCourseList []*entity.StudentCourse, err error) {
	if err = dao.SqlSession.Where("course_id", courseId).Find(&studentCourseList).Error; err != nil {
		return nil, err
	}
	return
}


/*
根据id删除课程
*/
func DeleteStudentCourseById(id uint) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.StudentCourse{}).Error
	return
}

/*
根据id查询课程
*/
func GetStudentCourseById(id uint) (studentCourse *entity.StudentCourse, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(studentCourse).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新课程信息
*/
func UpdateStudentCourse(studentCourse *entity.StudentCourse) (err error) {
	err = dao.SqlSession.Save(studentCourse).Error
	return
}
