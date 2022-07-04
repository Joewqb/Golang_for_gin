package main

import "C"
import (
	"gooooo/dao"
	"gooooo/models"
	"gooooo/routers"
)

func main() {
	//创建数据库CREATE DATABASE bub;
	//连接数据库

	err := dao.Initmysql()
	if err != nil {

		panic(err)
	}
	defer dao.DB.Close()
	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()
	r.Run(":8080")
}
