package tests

import (
	"testing"

	"github.com/yuedun/ginode-mongo/db"
	_ "github.com/yuedun/ginode-mongo/db"
	"github.com/yuedun/ginode-mongo/pkg/user"
)

func TestGetUser(t *testing.T) {
	userService := user.NewService(db.NewDB("website"))
	user, err := userService.GetUserInfoByName("test")
	if err != nil {
		t.Error(err)
	}
	t.Log(user)
}

func TestCreateUser(t *testing.T) {
	userService := user.NewService(db.NewDB("website"))
	newUser := new(user.User)
	newUser.Mobile = "17864345978"
	err := userService.CreateUser(newUser)
	if err != nil {
		t.Error(err)
	}
	t.Log(newUser)
}
