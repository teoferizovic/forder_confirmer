package model

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
