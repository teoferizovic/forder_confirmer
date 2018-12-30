package middleware

/*type redisStore2 struct {
	client *redis.Client
}

var conff model.Config
var redisConn2 redisStore2

func init(){
	//load config file
	if _, err := toml.DecodeFile("./config.toml", &conff); err != nil {
		fmt.Println(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     conff.RedisHost+":"+conf.RedisPort,
		Password: conff.RedisPassword,
		DB:       2,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	redisConn2 = redisStore2{client:client}
}

func  CacheMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		datas, err := redisConn2.client.Get(r.URL.String()).Result()
		fmt.Println(r.URL.String())
		if err != nil {
			next.ServeHTTP(w, r)
			fmt.Println("aaaa")
		} else {
			fmt.Println("bbbb")

			bytes := []byte(datas)
			var ps []model.FOrder{}

			err = json.Unmarshal(bytes, &ps)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(ps)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ps)
			next.ServeHTTP(w, r)
		}
	})

}*/
