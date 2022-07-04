package controller

import (
	"github.com/gin-gonic/gin"
	"gooooo/dao"
	"gooooo/models"
)

//URL---controller---logic---model
//请求来了--控制器--业务逻辑--模型层增删改查
func IndexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func CreatTodo(c *gin.Context) {
	//1.把数据从请求中拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	//err = DB.Create(&todo).Error
	err := models.CreatTodo(&todo)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, todo)
	}
	//3.返回响应

}

func GetTodolist(c *gin.Context) {
	todolist, err := models.GetallTodolist()
	if err = dao.DB.Find(&todolist).Error; err != nil {
		c.JSON(200, gin.H{"error": err})
	} else {
		c.JSON(200, todolist)
	}
}

func Update(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(200, gin.H{
			"error": "无效id",
		})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err := models.UpdateATodo(todo); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})

	} else {
		c.JSON(200, todo)
	}
}

func Delete(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(200, gin.H{
			"error": "无效id",
		})
	}
	if err := models.DeleteAid(id); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{id: "deleted"})
	}
}
