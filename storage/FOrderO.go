package storage

import (
	"forder_confirmer/model"
	"github.com/jinzhu/gorm"
	"strconv"
)

func GetAllFOrders(db *gorm.DB,id string,url string) ([]model.FOrder,error) {

	orders := []model.FOrder{}
	//user := model.User{}
	order := model.FOrder{}

	if len(id) > 0 {

		id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, err
		}



		if err := db.First(&order,id).Error; gorm.IsRecordNotFoundError(err) {
			return orders,nil
		}

		orders = append(orders,order)
		return orders,nil
	}


	//db.Model(&orders).Related(&user,"User")
	//db.Model(&order).Association("Users").Find(&order.User)
	db.Find(&orders)
	return orders,nil

}