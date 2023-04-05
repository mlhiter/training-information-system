package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建课程信息
*/
func CreateTrainingNotice(trainingNotice *entity.TrainingNotice) (err error) {
	if err = dao.SqlSession.Create(trainingNotice).Error; err != nil {
		return err
	}
	return
}

/*
获取课程集合
*/
func GetAllTrainingNotice() (trainingNoticeList []*entity.TrainingNotice, err error) {
	if err = dao.SqlSession.Find(&trainingNoticeList).Error; err != nil {
		return nil, err
	}
	return
}

/*
根据id删除课程
*/
func DeleteTrainingNoticeById(id uint) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.TrainingNotice{}).Error
	return
}

/*
根据id查询课程
*/
func GetTrainingNoticeById(id uint) (trainingNotice *entity.TrainingNotice, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(trainingNotice).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新课程信息
*/
func UpdateTrainingNotice(trainingNotice *entity.TrainingNotice) (err error) {
	err = dao.SqlSession.Save(trainingNotice).Error
	return
}
