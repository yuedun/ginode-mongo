package component

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type (
	ComponentService interface {
		GetComponentList(offset, limit int, search Component) (component []Component, count int, err error)
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

func (u *componentService) GetComponentList(offset, limit int, search Component) (components []Component, count int, err error) {
	if search.Name != "" {
		u.mysql = u.mysql.Where("name like ?", search.Name+"%")
	}
	if search.Category != "" {
		u.mysql = u.mysql.Where("category = ?", search.Category)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = u.mysql.Offset(offset).Limit(limit).Find(&components).Error
	}()
	go func() {
		defer wg.Done()
		err = u.mysql.Model(&Component{}).Count(&count).Error
	}()
	wg.Wait()
	return components, count, err
}

func (u *componentService) CreateComponent(component *Component) (err error) {
	// using unaddressable value Create(指针类型)
	err = u.mysql.Create(component).Error
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
