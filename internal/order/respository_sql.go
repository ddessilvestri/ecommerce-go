package order

import (
	"database/sql"

	"github.com/ddessilvestri/ecommerce-go/models"
)

type Repository interface {
	Insert(o models.Orders) (int, error)
	GetById(id int, userUUID string) (models.Orders, error)
	GetByUserUUID(userUUID string) ([]models.Orders, error)
	Update(o models.Orders) error
	Delete(id int, userUUID string) error
}

type repositorySQL struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) Repository {
	return &repositorySQL{db: db}
}

func (r *repositorySQL) Insert(o models.Orders) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	res, err := tx.Exec(`
		INSERT INTO orders (Order_UserUUID, Order_AddId, Order_Date, Order_Total)
		VALUES (?, ?, NOW(), ?)`,
		o.UserUUID, o.AddId, o.Total,
	)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	orderID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, d := range o.OrderDetails {
		_, err := tx.Exec(`
			INSERT INTO orders_details (OrderId, ProdId, Quantity, Price)
			VALUES (?, ?, ?, ?)`,
			orderID, d.ProdId, d.Quantity, d.Price,
		)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(orderID), nil
}

func (r *repositorySQL) GetById(id int, userUUID string) (models.Orders, error) {
	var o models.Orders
	err := r.db.QueryRow(`
		SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total
		FROM orders
		WHERE Order_Id = ? AND Order_UserUUID = ?`,
		id, userUUID,
	).Scan(&o.Id, &o.UserUUID, &o.AddId, &o.Date, &o.Total)
	if err != nil {
		return models.Orders{}, err
	}

	rows, err := r.db.Query(`
		SELECT Id, OrderId, ProdId, Quantity, Price
		FROM orders_details
		WHERE OrderId = ?`,
		id,
	)
	if err != nil {
		return models.Orders{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var d models.OrdersDetails
		if err := rows.Scan(&d.Id, &d.OrderId, &d.ProdId, &d.Quantity, &d.Price); err != nil {
			return models.Orders{}, err
		}
		o.OrderDetails = append(o.OrderDetails, d)
	}

	return o, nil
}

func (r *repositorySQL) GetByUserUUID(userUUID string) ([]models.Orders, error) {
	rows, err := r.db.Query(`
		SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total
		FROM orders
		WHERE Order_UserUUID = ?
		ORDER BY Order_Id DESC`,
		userUUID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Orders

	for rows.Next() {
		var o models.Orders
		err := rows.Scan(&o.Id, &o.UserUUID, &o.AddId, &o.Date, &o.Total)
		if err != nil {
			return nil, err
		}

		detailsRows, err := r.db.Query(`
			SELECT Id, OrderId, ProdId, Quantity, Price
			FROM orders_details
			WHERE OrderId = ?`,
			o.Id,
		)
		if err != nil {
			return nil, err
		}
		defer detailsRows.Close()

		for detailsRows.Next() {
			var d models.OrdersDetails
			if err := detailsRows.Scan(&d.Id, &d.OrderId, &d.ProdId, &d.Quantity, &d.Price); err != nil {
				return nil, err
			}
			o.OrderDetails = append(o.OrderDetails, d)
		}

		orders = append(orders, o)
	}

	return orders, nil
}

func (r *repositorySQL) Update(o models.Orders) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE orders
		SET Order_AddId = ?, Order_Total = ?
		WHERE Order_Id = ? AND Order_UserUUID = ?`,
		o.AddId, o.Total, o.Id, o.UserUUID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		DELETE FROM orders_details
		WHERE OrderId = ?`,
		o.Id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, d := range o.OrderDetails {
		_, err := tx.Exec(`
			INSERT INTO orders_details (OrderId, ProdId, Quantity, Price)
			VALUES (?, ?, ?, ?)`,
			o.Id, d.ProdId, d.Quantity, d.Price,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *repositorySQL) Delete(id int, userUUID string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		DELETE FROM orders_details
		WHERE OrderId = ?`,
		id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		DELETE FROM orders
		WHERE Order_Id = ? AND Order_UserUUID = ?`,
		id, userUUID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
