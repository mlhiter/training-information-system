package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建课程信息
*/
func CreateExecutor(executor *entity.Executor) (err error) {
	if err = dao.SqlSession.Create(executor).Error; err != nil {
		return err
	}
	return
}

/*
获取课程集合
*/
func GetAllExecutor() (executorList []*entity.Executor, err error) {
	if err = dao.SqlSession.Find(&executorList).Error; err != nil {
		return nil, err
	}
	return
}

/*
根据id删除课程
*/
func DeleteExecutorById(id uint) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.Executor{}).Error
	return
}

/*
根据id查询课程
*/
func GetExecutorById(id uint) (executor *entity.Executor, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(executor).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新课程信息
*/
func UpdateExecutor(executor *entity.Executor) (err error) {
	err = dao.SqlSession.Save(executor).Error
	return
}
