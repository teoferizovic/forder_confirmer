package model

import (
	"errors"
	"time"
)

type FOrder struct {
	//gorm.Model
	ID int64              		`json:"id" gorm:"primary_key"`
	User_id int64		  		`json:"user_id"`
	User   User			  		`json:"user"`
	Order_id int64        		`json:"order_id"`
	Payment_id int64      		`json:"payment_id"`
	Payment   Payment			`json:"payment"`
	Status string         		`json:"status"`
	Final_price float64   		`json:"final_price"`
	Created_at string			`json:"created_at"`
	Updated_at *time.Time		`json:"updated_at"`
	Deleted_at *string			`json:"deleted_at"`
}

func (fo *FOrder) BeforeSave() error {
	fo.Status = "F"
	return nil
}

func (fo *FOrder) Validate() error {

	// check if the user_id empty
	if fo.User_id == 0 {
		return errors.New("The field user_id is required!")
	}

	// check if the order_id empty
	if fo.Order_id == 0 {
		return errors.New("The field order_id is required!")
	}

	// check if the payment_id empty
	if fo.Payment_id == 0 {
		return errors.New("The field payment_id is required!")
	}

	// check if the status empty
	if fo.Status == "" {
		return errors.New("The field status is required!")
	}

	// check the status field is bigger then 1
	if len(fo.Status) > 1  {
		return errors.New("The field status can contain only 1 letter!")
	}

	// check if the created_at empty
	if fo.Final_price == 0 {
		return errors.New("The field final_price is required!")
	}

	return nil
}

type FOrderR struct {
	ID int64              `json:"id"`
	Order_id int64        `json:"order_id"`
	//User User			  `json:"User"`
	//Payment Payment		  `json:"Payment"`
	Status string         `json:"status"`
	Final_price float64   `json:"final_price"`
	Created_at string     `json:"created_at"`
}

type TemplateResponse struct {
	Message string
	//Error   error
}
