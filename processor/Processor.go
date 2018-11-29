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

