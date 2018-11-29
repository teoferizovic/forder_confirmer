package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forder_confirmer/model"
	"forder_confirmer/processor"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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


	fmt.Println(db)

	router := mux.NewRouter()
	router.HandleFunc("/forders/index", dbConn.Index).Methods("GET")
	//router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/forders/create", dbConn.Create).Methods("POST")
	//router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	http.ListenAndServe(":8060", router)
	defer db.Close()

}

func (conn *dbStore) Index(w http.ResponseWriter, r *http.Request){

	orders,err := processor.Index(conn.db)

    if err != nil {
   		panic(err)
    }

 	w.Header().Set("Content-Type", "application/json")
 	json.NewEncoder(w).Encode(orders)

}

func (conn *dbStore) Create(w http.ResponseWriter, r *http.Request){
fmt.Println("aaa")
fmt.Println(conn.db)
/*rows, err := db.db.Query("select id,user_id from f_orders")

if err != nil {
	log.Fatal(err)
}

fmt.Println(rows)*/
}


