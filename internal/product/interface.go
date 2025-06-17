package product

import "github.com/ddessilvestri/ecommerce-go/models"

type Storage interface {
	Insert(c models.Product) (int64, error)
	Update(c models.Product) error
	Delete(id int) error
	GetById(id int) (models.Product, error)
	GetAll() ([]models.Product, error)
	GetBySlug(slug string) ([]models.Product, error)
}
