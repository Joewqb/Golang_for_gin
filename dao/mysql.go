package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //必须导入数据库驱动
)

var (
	DB *gorm.DB
)

func Initmysql() (err error) {
	dsn := "root:@tcp(127.0.0.1:3306)/bub?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	DB.DB().Ping()
	return
}
