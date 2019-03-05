package storage

import (
	"forder_confirmer/model"
	"github.com/jinzhu/gorm"
)

func Create(db *gorm.DB,order model.FOrder) error {

	db.Create(&order)
	return nil

}
