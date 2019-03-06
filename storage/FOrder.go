package storage

import (
	"forder_confirmer/model"
	"github.com/jinzhu/gorm"
)

func Create(db *gorm.DB,order model.FOrder) error {

	if err := db.Create(&order).Error; err != nil {
		return err
	}

	return nil

}
