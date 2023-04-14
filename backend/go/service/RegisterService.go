package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建学生信息
*/
func CreateRegister(register *entity.Register) (err error) {
	if err = dao.SqlSession.Create(register).Error; err != nil {
		return err
	}
	return
}

/*
获取学生集合
*/
func GetAllRegister() (registerList []*entity.Register, err error) {
	if err = dao.SqlSession.Find(&registerList).Error; err != nil {
		return nil, err
	}
	return
}

/*
根据id删除学生
*/
func DeleteRegisterById(id uint) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.Register{}).Error
	return
}

/*
根据id查询学生
*/
func GetRegisterById(id uint) (register *entity.Register, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(register).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新学生信息
*/
func UpdateRegister(register *entity.Register) (err error) {
	err = dao.SqlSession.Model(register).Update("status", register.Status).Error
	return
}
