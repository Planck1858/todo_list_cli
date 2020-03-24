package router

import (
	"github.com/Planck1858/todo_list_cli/server/cache"
	"github.com/Planck1858/todo_list_cli/server/models"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	l := models.NewList()
	l.NewTask("First-Test")
	c := cache.NewRedisCache("localhost:6379", 1, 0)

	r.GET("/", IndexHandler())
	r.GET("/getList", GetListHandler(l))
	r.GET("/getLastTask", GetLastTaskHandler(l, c))
	r.POST("/newTask", NewTaskHandler(l, c))
	r.DELETE("/clearList", ClearListHandler(l))
	r.DELETE("/deleteTask/:index", DeleteTaskHandler(l))
	r.PUT("/changeTask/:index", ChangeTaskHandler(l))

	return r
}

//создать лист, обнулить лист, вернуть лист
//новый таск, удалить таск, изменить таск
