package post

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type PostService interface {
	GetPostInfo() (user Post, err error)
	GetPostInfoBySql() (user Post, err error)
	CreatePost(user *Post) (err error)
	UpdatePost(userId int, user *Post) (err error)
	DeletePost(userId int) (err error)
}
type postService struct {
	mysql *gorm.DB
}

func NewService(mysql *gorm.DB) PostService {
	return &postService{
		mysql: mysql,
	}
}

func (u *postService) GetPostInfo() (post Post, err error) {
	err = u.mysql.First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (u *postService) GetPostInfoBySql() (post Post, err error) {
	err = u.mysql.Raw("select * from post where id=?", post.Id).Scan(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (u *postService) CreatePost(post *Post) (err error) {
	err = u.mysql.Create(post).Error
	fmt.Println(post)
	if err != nil {
		return err
	}
	return nil
}

func (u *postService) UpdatePost(userId int, post *Post) (err error) {
	err = u.mysql.Model(post).Where("id = ?", userId).Update(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *postService) DeletePost(userId int) (err error) {
	u.mysql.Where("id = ?", userId).Delete(Post{})
	if err != nil {
		return err
	}
	return nil
}
