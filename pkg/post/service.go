package post

import (
	"fmt"
	"github.com/yuedun/ginode/db"
)

type PostService interface {
	GetPostInfo() (user Post, err error)
	GetPostInfoBySql() (user Post, err error)
	CreatePost(user *Post) (err error)
	UpdatePost(userId int, user *Post) (err error)
	DeletePost(userId int) (err error)
}
type userService struct {
}

func NewPostService() PostService {
	return &userService{}
}

func (u *userService) GetPostInfo() (post Post, err error) {
	err = db.Mysql.First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (u *userService) GetPostInfoBySql() (post Post, err error) {
	err = db.Mysql.Raw("select * from post where id=?", post.Id).Scan(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (u *userService) CreatePost(post *Post) (err error) {
	err = db.Mysql.Create(post).Error
	fmt.Println(post)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) UpdatePost(userId int, post *Post) (err error) {
	err = db.Mysql.Model(post).Where("id = ?", userId).Update(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) DeletePost(userId int) (err error) {
	db.Mysql.Where("id = ?", userId).Delete(Post{})
	if err != nil {
		return err
	}
	return nil
}
