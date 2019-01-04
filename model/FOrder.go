package model

import (
	"errors"
)

type FOrder struct {
	ID int64              `json:"id"`
	User_id int64		  `json:"user_id"`
	Order_id int64        `json:"order_id"`
	Payment_id int64      `json:"payment_id"`
	Status string         `json:"status"`
	Final_price float64   `json:"final_price"`
	Deleted_at string     `json:"deleted_at"`
	Created_at string     `json:"created_at"`
	Updated_at string     `json:"updated_at"`
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
	if fo.Created_at == "" {
		return errors.New("The field created_at is required!")
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
	User User			  `json:"User"`
	Payment Payment		  `json:"Payment"`
	Status string         `json:"status"`
	Final_price float64   `json:"final_price"`
	Created_at string     `json:"created_at"`
}

type TemplateResponse struct {
	Message string
	//Error   error
}
