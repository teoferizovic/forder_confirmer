package storage

import ("database/sql"
	"forder_confirmer/model"
)

func GetAllOrders(db *sql.DB) ([]model.Order,error) {

	rows, err := db.Query("SELECT o.id,COALESCE(deleted_at, ''),u.email as deleted_at from orders as o inner join users as u on u.id=o.user_id")

	if err != nil {
		return nil,err
	}

	orders := []model.Order{}

	for rows.Next() {

		var o model.Order
		err = rows.Scan(&o.ID,&o.Deleted_at,&o.User.Email)

		if err != nil {
			return nil,err
		}

		orders = append(orders, o)
	}

	return orders,nil
}