package model

type FOrder struct {
	ID int64
	User_id int64
	Order_id int64
	Payment_id int64
	Status string
	Final_price float64
	Deleted_at string
	Created_at string
	Updated_at string
}

type TemplateResponse struct {
	Message string
	//Error   error
}
