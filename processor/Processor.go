package processor

import (
	"database/sql"
	"forder_confirmer/model"
)

func Index(db *sql.DB) ([]model.FOrder,error) {

	rows, err := db.Query("select id,user_id from f_orders")

		if err != nil {
			return nil,err
		}

		orders := []model.FOrder{}

		for rows.Next() {

			var or model.FOrder
			err = rows.Scan(&or.ID, &or.User_id)

			if err != nil {
				return nil,err
			}

			orders = append(orders, or)
		}

		return orders,nil

}

func Create(db *sql.DB,order model.FOrder) error {

	insert, err := db.Prepare("INSERT INTO f_orders(user_id, order_id,payment_id,status,final_price,created_at) VALUES(?,?,?,?,?,?)")

	if err != nil {
		return err
	}

	insert.Exec(order.User_id,order.Order_id,order.Payment_id,order.Status,order.Final_price,order.Created_at)
	insert.Close()

	return nil
}

