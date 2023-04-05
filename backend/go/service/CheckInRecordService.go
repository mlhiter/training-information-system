package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建课程信息
*/
func CreateCheckInRecord(checkInRecord *entity.CheckInRecord) (err error) {
	if err = dao.SqlSession.Create(checkInRecord).Error; err != nil {
		return err
	}
	return
}

/*
获取课程集合
*/
func GetAllCheckInRecord() (checkInRecordList []*entity.CheckInRecord, err error) {
	if err = dao.SqlSession.Find(&checkInRecordList).Error; err != nil {
		return nil, err
	}
	return
}

/*
根据id删除课程
*/
func DeleteCheckInRecordById(id uint) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.CheckInRecord{}).Error
	return
}

/*
根据id查询课程
*/
func GetCheckInRecordById(id uint) (checkInRecord *entity.CheckInRecord, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(checkInRecord).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新课程信息
*/
func UpdateCheckInRecord(checkInRecord *entity.CheckInRecord) (err error) {
	err = dao.SqlSession.Save(checkInRecord).Error
	return
}
