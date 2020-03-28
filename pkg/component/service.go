package component

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type (
	ComponentService interface {
		GetComponentList(offset, limit int, search Component) (component []Component, err error)
		CreateComponent(component *Component) (err error)
		UpdateComponent(component *Component) (err error)
		DeleteComponent(componentID int) (err error)
	}
)
type componentService struct {
	mysql *gorm.DB
}

func NewService(mysql *gorm.DB) ComponentService {
	return &componentService{
		mysql: mysql,
	}
}

func (u *componentService) GetComponentList(offset, limit int, search Component) (components []Component, err error) {
	if search.Name != "" {
		u.mysql = u.mysql.Where("name like ?", search.Name+"%")
	}
	if search.Category != "" {
		u.mysql = u.mysql.Where("category = ?", search.Category)
	}
	err = u.mysql.Offset(offset).Limit(limit).Find(&components).Error
	if err != nil {
		return components, err
	}
	return components, nil
}

func (u *componentService) CreateComponent(component *Component) (err error) {
	err = u.mysql.Create(component).Error
	fmt.Println(component)
	if err != nil {
		return err
	}
	return nil
}

func (u *componentService) UpdateComponent(component *Component) (err error) {
	err = u.mysql.Model(component).Update(component).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *componentService) DeleteComponent(componentID int) (err error) {
	u.mysql.Where("id = ?", componentID).Delete(Component{})
	if err != nil {
		return err
	}
	return nil
}
