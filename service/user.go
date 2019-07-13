package service

import (
	"fmt"

	. "github.com/yuedun/ginode/model"
)

type UserService struct {
}

func (s *UserService) GetUserInfo() {
	user := new(User)
	fmt.Println(user)
}
