package product

import "github.com/ddessilvestri/ecommerce-go/models"

type Storage interface {
	Insert(c models.Product) (int64, error)
	Update(c models.Product) error
	Delete(id int) error
	GetById(id int) (models.Product, error)
	GetByCategoryId(id int) ([]models.Product, error)
	GetAll(page, limit int, sortBy, order string) ([]models.Product, error)
	GetBySlug(slug string) (models.Product, error)
	GetByCategorySlug(slug string) ([]models.Product, error)
	SearchByText(text string, page, limit int, sortBy, order string) ([]models.Product, error)
}
