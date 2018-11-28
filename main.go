package main

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_"github.com/go-sql-driver/mysql"
	"forder_confirmer/model"

)

var conf model.Config


func init(){
	//load config file
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%#v\n", conf)
	//open db connection
	db, err := sql.Open("mysql", conf.Username+":"+conf.Password+"@tcp("+conf.Host+":"+conf.Port+")/"+conf.DB)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
func main() {
	fmt.Println("aaaaa")
}
