package processor

import (
	"database/sql"
	"encoding/json"
	"forder_confirmer/database"
	"forder_confirmer/model"
	"forder_confirmer/storage"
	"time"
)

func Index(db *sql.DB,id string,url string) ([]model.FOrderR,error) {


		orders, err := storage.GetAllFOrders(db,id,url)

		if err != nil{
			return nil,err
		}

		orderBytes, _ := json.Marshal(orders)

		//set route in cache with 10 second of expiration
		err = database.RedisConn2().Set(url, string(orderBytes), 160*time.Second).Err()
		if err != nil {
			return nil,err
		}

		return orders,nil

}

func Create(db *sql.DB,order model.FOrder) error {

	if validErrs := order.Validate();validErrs!=nil {
		return validErrs
	}

	err := storage.Create(db,order)

	if err != nil{
		return err
	}

	return nil
}

func IndexO(db *sql.DB,id string,url string) ([]model.Order,error) {

	orders, err := storage.GetAllOrders(db)

	if err != nil {
		return nil,err
	}

	orderBytes, _ := json.Marshal(orders)

	//set route in cache with 10 second of expiration
	err = database.RedisConn2().Set(url, string(orderBytes), 160*time.Second).Err()
	if err != nil {
		return nil,err
	}

	return orders,nil
}


