package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forder_confirmer/middleware"
	"forder_confirmer/model"
	"forder_confirmer/processor"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var conf model.Config
var db *sql.DB
var err error

type dbStore struct {
	db *sql.DB
}

var dbConn dbStore


func init(){

	//load config file
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}

	//open db connection
	db, err = sql.Open("mysql", conf.Username+":"+conf.Password+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.DB)

	if err != nil {
		panic(err.Error())
	}

	dbConn = dbStore{db: db}

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/forders/index/", dbConn.Index).Methods("GET")
	router.HandleFunc("/forders/create", dbConn.Create).Methods("POST")
	router.Use(middleware.AuthenticationMiddleware,middleware.RecoverFromPanic)
	http.ListenAndServe(":8060", router)

	defer db.Close()

}

func (conn *dbStore) Index(w http.ResponseWriter, r *http.Request){

	orders,err := processor.Index(conn.db,r.URL.Query().Get("id"))

    if err != nil {
   		panic(err)
    }

 	w.Header().Set("Content-Type", "application/json")
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

	err = processor.Create(conn.db,fo)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	res := model.TemplateResponse{"Successfully saved item"}
	json.NewEncoder(w).Encode(res)

}


