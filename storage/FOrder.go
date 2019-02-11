package storage

import (
	"database/sql"
	"forder_confirmer/model"
)

func Create(db *sql.DB,order model.FOrder) error {

	insert, err := db.Prepare("INSERT INTO f_orders(user_id, order_id,payment_id,status,final_price,created_at) VALUES(?,?,?,?,?,?)")

	if err != nil {
		return err
	}

	insert.Exec(order.User_id,order.Order_id,order.Payment_id,order.Status,order.Final_price,order.Created_at)
	insert.Close()

	return nil
}
