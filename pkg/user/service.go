package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

/**
面向接口开发：
面向接口开发的好处是要对下面的函数进行测试时，不需要依赖一个全局的mysql连接，只需要调用NewService传一个mysql连接参数即可测试
 */
type UserService interface {
	GetUserInfo() (user User, err error)
	GetUserInfoBySql() (user User, err error)
	CreateUser(user *User) (err error)
	UpdateUser(userId int, user *User) (err error)
	DeleteUser(userId int) (err error)
}
type userService struct {
	mysql *gorm.DB
}

func NewService(mysql *gorm.DB) UserService {
	return &userService{
		mysql: mysql,
	}
}
func (u *userService) GetUserInfo() (user User, err error) {
	err = u.mysql.First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userService) GetUserInfoBySql() (user User, err error) {
	err = u.mysql.Raw("select * from user where id=?", user.Id).Scan(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userService) CreateUser(user *User) (err error) {
	err = u.mysql.Create(user).Error
	fmt.Println(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) UpdateUser(userId int, user *User) (err error) {
	err = u.mysql.Model(user).Where("id = ?", userId).Update(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) DeleteUser(userId int) (err error) {
	u.mysql.Where("id = ?", userId).Delete(User{})
	if err != nil {
		return err
	}
	return nil
}
