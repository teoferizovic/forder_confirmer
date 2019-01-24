package database

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis"
	"forder_confirmer/model"
)

var conf model.Config

func RedisConn2()  *redis.Client {

	//load config file
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisHost+":"+conf.RedisPort,
		Password: conf.RedisPassword,
		DB:       2,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client

}