package service

import (
	"fmt"
	"github.com/yuedun/ginode/model"
)

func GetUserInfo() (user model.User, err error) {
	err = model.Db.First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserInfoBySql() (user model.User, err error) {
	err = model.Db.Raw("select * from user where id=?", user.Id).Scan(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user *model.User) (err error) {
	err = model.Db.Create(user).Error
	fmt.Println(user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(userId int, user *model.User) (err error) {
	err = model.Db.Model(user).Where("id = ?", userId).Update(user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(userId int) (err error) {
	model.Db.Where("id = ?", userId).Delete(model.User{})
	if err != nil {
		return err
	}
	return nil
}
