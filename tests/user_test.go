package tests

import (
	"testing"

	"github.com/yuedun/ginode-mongo/db"
	_ "github.com/yuedun/ginode-mongo/db"
	"github.com/yuedun/ginode-mongo/pkg/page"
	"github.com/yuedun/ginode-mongo/pkg/user"
	"github.com/yuedun/ginode-mongo/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func TestQuery(t *testing.T) {
	q := page.Page{Name: "主页", WebsiteID: primitive.NewObjectID()}
	r, _ := util.Qeury(q)
	t.Log(r)
}
