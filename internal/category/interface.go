package category

import "github.com/ddessilvestri/ecommerce-go/models"

type Storage interface {
	InsertCategory(c models.Category) (int64, error)
	UpdateCategory(c models.Category) error
	DeleteCategory(id int) error
	GetCategory(id int) (models.Category, error)
	GetCategories() ([]models.Category, error)
	GetCategoryBySlug(slug string) ([]models.Category, error)
}
