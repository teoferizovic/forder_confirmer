package processor

import (
	"encoding/json"
	"forder_confirmer/database"
	"forder_confirmer/model"
	"forder_confirmer/storage"
	"github.com/jinzhu/gorm"
	"time"
)

func Index(db *gorm.DB,id string,url string) ([]model.FOrder,error) {

		orders, err := storage.GetAllFOrders(db,id,url)

		if err != nil{
			return nil,err
		}

		orderBytes, _ := json.Marshal(orders)

		//set route in cache with 160 second of expiration
		err = database.RedisConn2().Set(url, string(orderBytes), 160*time.Second).Err()
		if err != nil {
			return nil,err
		}

		return orders,nil

}

func Create(db *gorm.DB,order model.FOrder,conf model.Config) error {

		if validErrs := order.Validate();validErrs!=nil {
			return validErrs
		}

		err := storage.Create(db,order)

		if err != nil{
			return err
		}

		//publish to Redis channel to comunicate with other Go microservice
		database.RedisConn2().Publish(conf.RedisChannel,"Successfuly saved order.")

		//clear redis cache by route
		database.RedisConn2().FlushDb()

		return nil
}

func IndexO(db *gorm.DB,id string,url string) ([]model.Order,error) {

		orders, err := storage.GetAllOrders(db)

		if err != nil {
			return nil,err
		}

		orderBytes, _ := json.Marshal(orders)

		//set route in cache with 160 second of expiration
		err = database.RedisConn2().Set(url, string(orderBytes), 160*time.Second).Err()
		if err != nil {
			return nil,err
		}

		return orders,nil
}


