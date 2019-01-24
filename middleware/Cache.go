package middleware

import (
	"encoding/json"
	"fmt"
	"forder_confirmer/database"
	/*"forder_confirmer/model"*/
	"net/http"
)


func  CacheMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method=="GET" {
			datas, err := database.RedisConn2().Get(r.URL.String()).Result()

			if err != nil {
				next.ServeHTTP(w, r)
				fmt.Println("aaaa")
			} else {
				fmt.Println("bbbb")
				bytes := []byte(datas)

				/*var orders []model.FOrderR

				err = json.Unmarshal(bytes, &orders)
				if err != nil {
					fmt.Println(err)
				}*/

				var raw interface{}
				err = json.Unmarshal(bytes, &raw)

				if err != nil {
					fmt.Println(err)
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(raw)
				//next.ServeHTTP(w, r)
				return
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})

}
