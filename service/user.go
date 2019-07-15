package service

import (
	. "github.com/yuedun/ginode/model"
)

func GetUserInfo() (User, error) {
	var user User
	err := Db.First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
