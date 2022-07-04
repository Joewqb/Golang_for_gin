package routers

import (
	"github.com/gin-gonic/gin"
	"gooooo/controller"
)

func SetupRouter() *gin.Engine {
	//模型绑定
	r := gin.Default()
	r.Static("/static", "static")
	//告诉gin框架模版文件引用的静态文件去static找，对应根目录下的static
	r.LoadHTMLGlob("templates/*")
	//告诉gin框架去哪里找模板文件
	r.GET("/", controller.IndexHandler)
	//定义接口v1
	v1Group := r.Group("v1")
	{
		//增删改查
		//添加
		v1Group.POST("/todo", controller.CreatTodo)
		//查看所有代办事件
		v1Group.GET("/todo", controller.GetTodolist)
		////查看某个代办事件
		//v1Group.GET("/todo/:id", func(c *gin.Context) {
		//
		//})
		//修改某一个
		v1Group.PUT("/todo/:id", controller.Update)
		//删除某一个
		v1Group.DELETE("/todo/:id", controller.Delete)
	}
	return r
}
