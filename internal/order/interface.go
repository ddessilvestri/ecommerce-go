package order

import "github.com/ddessilvestri/ecommerce-go/models"

type Storage interface {
	Insert(o models.Orders) (int64, error)
	GetById(id int) (models.Orders, error)
	GetAllByUserUUID(page int, limit int, fromDate string, toDate string, userUUID string) ([]models.Orders, error)
	Update(o models.Orders) error
	Delete(id int, userUUID string) error
}
