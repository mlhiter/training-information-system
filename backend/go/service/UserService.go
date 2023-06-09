package service

import (
	"backend/go/dao"
	"backend/go/entity"
)

/*
新建User信息
*/
func CreateUser(user *entity.User) (err error) {
	if err = dao.SqlSession.Create(user).Error; err != nil {
		return err
	}
	return
}

/*
获取user集合
*/
func GetAllUser() (userList []*entity.User, err error) {
	if err = dao.SqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

/*
根据id删除user
*/
func DeleteUserById(id string) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.User{}).Error
	return
}

/*
根据userName查询用户User
*/
func GetUserByUserName(userName string) (returnUser *entity.User, err error) {
	var user entity.User
	if err = dao.SqlSession.Table("user").Where("user_name=?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	returnUser = &user
	return
}

/*
根据id查询用户User
*/
func GetUserById(id uint) (user *entity.User, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新用户信息
*/
func UpdateUser(user *entity.User) (err error) {
	err = dao.SqlSession.Save(user).Error
	return
}
