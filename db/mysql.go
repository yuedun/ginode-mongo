package db

import "github.com/jinzhu/gorm"

var Mysql *gorm.DB

func init() {
	var err error
	Mysql, err = gorm.Open("mysql", "root:root@/issue?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	Mysql.LogMode(true)
	//Db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	//defer Db.Close()
}
