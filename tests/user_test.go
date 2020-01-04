package tests

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yuedun/ginode/db"
	"github.com/yuedun/ginode/pkg/user"
	"testing"
)

func TestGetUser(t *testing.T) {
	userService := user.NewService(db.Mysql)
	user, err := userService.GetUserInfo(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}

func TestCreateUser(t *testing.T) {
	userService := user.NewService(db.Mysql)
	newUser := new(user.User)
	newUser.Mobile = "17864345978"
	err := userService.CreateUser(newUser)
	if err != nil {
		t.Error(err)
	}
	t.Log(newUser)
}
