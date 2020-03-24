package router

import (
	"github.com/Planck1858/todo_list_cli/server/cache"
	"github.com/Planck1858/todo_list_cli/server/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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

func GetLastTaskHandler(l *models.List, cache *cache.RedisCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		lastIndex := len(l.GetList()) - 1
		var todo *models.ToDo = cache.Get(strconv.Itoa(lastIndex))

		c.JSON(http.StatusOK, todo)
	}
}

/// POST ///
func NewTaskHandler(l *models.List, cache *cache.RedisCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := c.Request.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			panic(err)
		}

		l.NewTask(string(data))
		index := len(l.GetList()) - 1
		rand.Seed(time.Now().UnixNano())
		min, max := 6, 20
		cache.Set(strconv.Itoa(index), l.GetList()[index], time.Second*time.Duration(rand.Intn(max-min+1)+min))
		//time.Second*time.Duration(rand.Intn(max-min+1)+min)

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
