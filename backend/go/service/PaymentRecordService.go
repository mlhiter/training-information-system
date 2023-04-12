package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建学生信息
*/
func CreatePaymentRecord(paymentRecord *entity.PaymentRecord) (err error) {
	if err = dao.SqlSession.Create(paymentRecord).Error; err != nil {
		return err
	}
	return
}

/*
获取学生集合
*/
func GetAllPaymentRecordByStudentId(id uint) (paymentRecordList []*entity.PaymentRecord, err error) {
	if err = dao.SqlSession.Where("id=?", id).Find(&paymentRecordList).Error; err != nil {
		return nil, err
	}
	return
}

/*
获取学生集合
*/
func GetAllPaymentRecord() (paymentRecordList []*entity.PaymentRecord, err error) {
	if err = dao.SqlSession.Find(&paymentRecordList).Error; err != nil {
		return nil, err
	}
	return
}

/*
根据id删除学生
*/
func DeletePaymentRecordById(id uint) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.PaymentRecord{}).Error
	return
}

/*
根据id查询学生
*/
func GetPaymentRecordById(id uint) (paymentRecord *entity.PaymentRecord, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(paymentRecord).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新学生信息
*/
func UpdatePaymentRecord(paymentRecord *entity.PaymentRecord) (err error) {
	err = dao.SqlSession.Save(paymentRecord).Error
	return
}
