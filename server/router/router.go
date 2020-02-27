package router

import (
	"github.com/Planck1858/todo_list_cli/server/models"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	l := models.NewList()

	r.GET("/", IndexHandler())
	r.GET("/getList", GetListHandler(l))
	r.POST("/newTask", NewTaskHandler(l))
	r.DELETE("/clearList", ClearListHandler(l))
	r.DELETE("/deleteTask/:index", DeleteTaskHandler(l))
	r.PUT("/changeTask/:index/:flag", ChangeTaskHandler(l))

	return r
}

//создать лист, обнулить лист, вернуть лист
//новый таск, удалить таск, изменить таск
