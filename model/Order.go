package model

type Order struct {
	ID int64                       `json:"id"`
	User_id int64		  		   `json:"user_id,omitempty"`
	User User			  		   `json:"User,omitempty"`
	Payment Payment		  		   `json:"Payment,omitempty"`
	Payment_id int64               `json:"payment_id,omitempty"`
	Final_price float64   		   `json:"final_price,omitempty"`
	Status string         		   `json:"status,omitempty"`
	Deleted_at string      		   `json:"deleted_at,omitempty"`
	Created_at string              `json:"created_at,omitempty"`
	Updated_at string              `json:"updated_at,omitempty"`
}

