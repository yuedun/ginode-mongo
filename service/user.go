package service

import (
	"github.com/yuedun/ginode/model"
)

func GetUserInfo() (model.User, error) {
	var user model.User
	err := model.Db.First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
	}

func GetUserInfoBySql() (model.User, error) {
	var user model.User
	err := model.Db.Raw("select * from user where id=?", 2).Scan(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}