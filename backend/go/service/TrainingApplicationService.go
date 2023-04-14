package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建学生信息
*/
func CreateTrainingApplication(trainingApplication *entity.TrainingApplication) (err error) {
	if err = dao.SqlSession.Create(trainingApplication).Error; err != nil {
		return err
	}
	return
}

/*
获取学生集合
*/
func GetAllTrainingApplication() (trainingApplicationList []*entity.TrainingApplication, err error) {
	if err = dao.SqlSession.Find(&trainingApplicationList).Error; err != nil {
		return nil, err
	}
	return
}

/*
根据id删除学生
*/
func DeleteTrainingApplicationById(id uint) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.TrainingApplication{}).Error
	return
}

/*
根据id查询学生
*/
func GetTrainingApplicationById(id uint) (trainingApplication *entity.TrainingApplication, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(trainingApplication).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新学生信息
*/
func UpdateTrainingApplication(trainingApplication *entity.TrainingApplication) (err error) {
	err = dao.SqlSession.Model(trainingApplication).Update("status", trainingApplication.Status).Error
	return
}
