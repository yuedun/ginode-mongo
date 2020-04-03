package tests

import (
	"github.com/yuedun/ginode-mongo/db"
	_ "github.com/yuedun/ginode-mongo/db"
	"github.com/yuedun/ginode-mongo/pkg/user"
	"testing"
)

func TestGetUser(t *testing.T) {
	userService := user.NewService(db.Mongodb)
	user, err := userService.GetUserInfo(user.User{})
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}

func TestCreateUser(t *testing.T) {
	userService := user.NewService(db.Mongodb)
	newUser := new(user.User)
	newUser.Mobile = "17864345978"
	err := userService.CreateUser(newUser)
	if err != nil {
		t.Error(err)
	}
	t.Log(newUser)
}
