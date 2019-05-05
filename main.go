package main

import (
	"encoding/json"
	"fmt"
	"forder_confirmer/middleware"
	"forder_confirmer/model"
	"forder_confirmer/processor"
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"io/ioutil"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var conf model.Config
var db *gorm.DB
var err error

type dbStore struct {
	db *gorm.DB
	client *redis.Client
}

var dbConn dbStore

func init(){

	//load config file
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}

	//open db connection
	db, err := gorm.Open("mysql", conf.Username+":"+conf.Password+"@/"+conf.DB+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	//defer db.Close()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisHost+":"+conf.RedisPort,
		Password: conf.RedisPassword,
		DB:       2,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	dbConn = dbStore{db: db,client:client}

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/forders/index/", dbConn.Index).Methods("GET")
	router.HandleFunc("/orders/index/", dbConn.IndexO).Methods("GET")
	router.HandleFunc("/forders/create", dbConn.Create).Methods("POST")
	router.Use(middleware.AuthenticationMiddleware,middleware.RecoverFromPanic,middleware.Logger,middleware.CacheMiddleware)
	http.ListenAndServe(":8060", router)

	defer db.Close()

}

func (conn *dbStore) Index(w http.ResponseWriter, r *http.Request){

	orders,err := processor.Index(conn.db,r.URL.Query().Get("id"),r.URL.String())

    if err != nil {
   		panic(err)
    }

 	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
 	json.NewEncoder(w).Encode(orders)

}

func (conn *dbStore) IndexO(w http.ResponseWriter, r *http.Request){

	orders,err := processor.IndexO(conn.db,r.URL.Query().Get("id"),r.URL.String())

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)

}

func (conn *dbStore) Create(w http.ResponseWriter, req *http.Request){

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fo := model.FOrder{}

	err = json.Unmarshal(body, &fo)
	if err != nil {
		panic(err)
	}

	err = processor.Create(conn.db,fo,conf)

	if err != nil {
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"Error:": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	res := model.TemplateResponse{"Successfully saved item"}
	json.NewEncoder(w).Encode(res)

}


