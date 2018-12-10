package processor

import (
	"database/sql"
	"forder_confirmer/model"
	"strconv"
)

func Index(db *sql.DB,id string) ([]model.FOrderR,error) {

	rows, err := db.Query("SELECT fo.id,fo.order_id,fo.final_price,fo.created_at,fo.status,u.id,u.email,u.password,u.role_id,u.created_at,u.updated_at,p.id,p.name FROM f_orders as fo inner join users as u on fo.user_id = u.id inner join payments as p on fo.payment_id = p.id;")

		if err != nil {
			return nil,err
		}

		orders := []model.FOrderR{}

		for rows.Next() {

			var or model.FOrderR
			err = rows.Scan(&or.ID, &or.Order_id,&or.Final_price,&or.Created_at,&or.Status,&or.User.ID,&or.User.Email,&or.User.Password,&or.User.Role_id,&or.User.Created_at,&or.User.Updated_at,&or.Payment.ID,&or.Payment.Name)

			if err != nil {
				return nil,err
			}

			orders = append(orders, or)
		}

		if len(id) > 0 {

			id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				return nil,err
			}

			for _, order := range orders {
				if order.ID == id {
					orders = orders[:0]
					orders = append(orders, order)
					break
				} else {
					orders = orders[:0]
				}
			}

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

