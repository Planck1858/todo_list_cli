package cache

import (
	"encoding/json"
	"github.com/Planck1858/todo_list_cli/server/models"
	"github.com/go-redis/redis/v7"
	"time"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) *RedisCache {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (c *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.host,
		Password: "",
		DB:       c.db,
	})
}

func (c *RedisCache) Set(key string, value models.ToDo, sec time.Duration) {
	client := c.getClient()

	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(key, data, sec)
}

func (c *RedisCache) Get(key string) *models.ToDo {
	client := c.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	todo := models.ToDo{}
	err = json.Unmarshal([]byte(val), &todo)
	if err != nil {
		panic(err)
	}

	return &todo
}
