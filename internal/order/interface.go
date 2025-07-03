package order

import "github.com/ddessilvestri/ecommerce-go/models"

type Storage interface {
	Insert(o models.Orders) (int, error)
	GetById(id int, userUUID string) (models.Orders, error)
	GetByUserUUID(userUUID string) ([]models.Orders, error)
	Update(o models.Orders) error
	Delete(id int, userUUID string) error
}
