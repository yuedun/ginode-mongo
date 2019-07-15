package model

import "github.com/jinzhu/gorm"

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:root@/issue?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	Db.LogMode(true)
	Db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	//defer Db.Close()
}
