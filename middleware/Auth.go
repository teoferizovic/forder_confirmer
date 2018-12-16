package middleware

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis"
	"net/http"
	"strings"
	"forder_confirmer/model"
)

type redisStore struct {
	client *redis.Client
}

var conf model.Config
var redisConn redisStore

func init(){

	//load config file
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisHost+":"+conf.RedisPort,
		Password: conf.RedisPassword,
		DB:       conf.RedisDB,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	redisConn = redisStore{client:client}
}

func AuthenticationMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		_, err := redisConn.client.Get(token).Result()

		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)

	})

}

//https://github.com/gorilla/mux#middleware
//https://stackoverflow.com/questions/24790175/when-is-the-init-function-run
