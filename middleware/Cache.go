package middleware

import (
	"encoding/json"
	"fmt"
	"forder_confirmer/database"
	"forder_confirmer/model"
	"net/http"
)


func  CacheMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		datas, err := database.RedisConn2().Get(r.URL.String()).Result()

		if err != nil {
			next.ServeHTTP(w, r)
			fmt.Println("aaaa")
		} else {
			fmt.Println("bbbb")
			bytes := []byte(datas)

			var orders []model.FOrderR

			err = json.Unmarshal(bytes, &orders)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println(ps)
			//fmt.Println(datas)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(orders)
			//next.ServeHTTP(w, r)
			return
		}
	})

}
