package storage

import (
	"forder_confirmer/model"
	"github.com/jinzhu/gorm"
)

func GetAllOrders(db *gorm.DB) ([]model.Order,error) {

	orders := []model.Order{}
	db.Find(&orders)

	return orders,nil
}