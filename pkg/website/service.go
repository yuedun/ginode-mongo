package website

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type (
	WebsiteService interface {
		GetPostInfo() (website Website, err error)
		GetPostInfoBySQL() (website Website, err error)
		CreatePost(website *Website) (err error)
		UpdatePost(websiteID int, website *Website) (err error)
		DeletePost(websiteID int) (err error)
	}
)
type websiteService struct {
	mysql *gorm.DB
}

func NewService(mysql *gorm.DB) WebsiteService {
	return &websiteService{
		mysql: mysql,
	}
}

func (u *websiteService) GetPostInfo() (post Website, err error) {
	err = u.mysql.First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (u *websiteService) GetPostInfoBySQL() (post Website, err error) {
	err = u.mysql.Raw("select * from post where id=?", post.Id).Scan(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (u *websiteService) CreatePost(post *Website) (err error) {
	err = u.mysql.Create(post).Error
	fmt.Println(post)
	if err != nil {
		return err
	}
	return nil
}

func (u *websiteService) UpdatePost(websiteID int, post *Website) (err error) {
	err = u.mysql.Model(post).Where("id = ?", websiteID).Update(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *websiteService) DeletePost(websiteID int) (err error) {
	u.mysql.Where("id = ?", websiteID).Delete(Website{})
	if err != nil {
		return err
	}
	return nil
}
