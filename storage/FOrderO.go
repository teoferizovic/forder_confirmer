package storage

import (
	"forder_confirmer/model"
	"github.com/jinzhu/gorm"
	"strconv"
)

func GetAllFOrders(db *gorm.DB,id string,url string) ([]model.FOrder,error) {

	orders := []model.FOrder{}

	if len(id) > 0 {

		id, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			return nil, err
		}

		order := model.FOrder{}

		if err := db.Preload("User").Preload("Payment").First(&order,id).Error; gorm.IsRecordNotFoundError(err) {
			return orders,nil
		}

		orders = append(orders,order)
		return orders,nil
	}

	db.Preload("User").Preload("Payment").Find(&orders)
	return orders,nil

}