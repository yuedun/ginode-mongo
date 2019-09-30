package user

import (
	"fmt"
	"github.com/yuedun/ginode/db"
)

type UserService interface {
	GetUserInfo() (user User, err error)
	GetUserInfoBySql() (user User, err error)
	CreateUser(user *User) (err error)
	UpdateUser(userId int, user *User) (err error)
	DeleteUser(userId int) (err error)
}
type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}
func (u *userService) GetUserInfo() (user User, err error) {
	err = db.Mysql.First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userService) GetUserInfoBySql() (user User, err error) {
	err = db.Mysql.Raw("select * from user where id=?", user.Id).Scan(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userService) CreateUser(user *User) (err error) {
	err = db.Mysql.Create(user).Error
	fmt.Println(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) UpdateUser(userId int, user *User) (err error) {
	err = db.Mysql.Model(user).Where("id = ?", userId).Update(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) DeleteUser(userId int) (err error) {
	db.Mysql.Where("id = ?", userId).Delete(User{})
	if err != nil {
		return err
	}
	return nil
}
