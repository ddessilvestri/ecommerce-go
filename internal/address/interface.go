package address

import "github.com/ddessilvestri/ecommerce-go/models"

type Storage interface {
	Insert(a models.Address, userUUID string) (int64, error)
	Update(a models.Address) error
	Delete(id int) error
	GetById(id int) (models.Address, error)
	GetAllByUserUUID(userUUID string) ([]models.Address, error)
	Exists(id int) bool
}
