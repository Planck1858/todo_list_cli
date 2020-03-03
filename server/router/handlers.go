package router

import (
	"github.com/Planck1858/todo_list_cli/server/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type textResp struct {
	data string `json:"data"`
}

/// GET ///
func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to server !")
	}
}

func GetListHandler(l *models.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, l.GetList())
	}
}

/// POST ///
func NewTaskHandler(l *models.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := c.Request.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			panic(err)
		}

		l.NewTask(string(data))
		c.String(http.StatusOK, "New task created !")
	}
}

/// DELETE ///
func ClearListHandler(l *models.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		l.ClearList()
		c.String(http.StatusOK, "List is empty now !")
	}
}

func DeleteTaskHandler(l *models.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			panic(err)
		}

		l.DeleteTask(index - 1)
		c.String(http.StatusOK, "Element was deleted !")
	}
}

/// PUT ///
func ChangeTaskHandler(l *models.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index"))
		if err != nil {
			panic(err)
		}

		l.ChangeTask(index - 1)
		c.String(http.StatusOK, "Task was changed !")
	}
}
