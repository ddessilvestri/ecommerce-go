package category

import "github.com/ddessilvestri/ecommerce-go/models"

type Storage interface {
	InsertCategory(c models.Category) (int64, error)
	UpdateCategory(c models.Category) error
}
