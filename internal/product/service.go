package product

import (
	"errors"

	"github.com/ddessilvestri/ecommerce-go/models"
)

// Service provides methods for business logic related to category.
type Service struct {
	repo Storage // This is the interface, so it's decoupled from repositorySQL
}

func NewService(repo Storage) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(c models.Product) (int64, error) {
	if c.Title == "" {
		return 0, ErrInvalidProduct
	}

	return s.repo.Insert(c)
}

func (s *Service) UpdateProduct(c models.Product) error {
	if c.Title == "" {
		return ErrInvalidProduct
	}
	// if c.CategID < 1 {
	// 	return ErrInvalidProductId
	// }
	return s.repo.Update(c)

}

func (s *Service) Delete(id int) error {
	if id < 1 {
		return ErrInvalidProductId
	}
	return s.repo.Delete(id)

}

func (s *Service) Get(id int) (models.Product, error) {

	if id < 1 {
		return models.Product{}, ErrInvalidProductId
	}

	return s.repo.GetById(id)

}

func (s *Service) GetAll() ([]models.Product, error) {

	return s.repo.GetAll()

}

func (s *Service) GetBySlug(slug string) ([]models.Product, error) {

	if slug == "" {
		return []models.Product{}, ErrInvalidProductSlug
	}

	return s.repo.GetBySlug(slug)

}

var ErrInvalidProduct = errors.New("invalid product: title is required")
var ErrInvalidProductId = errors.New("invalid product Id: Id < 1 ")
var ErrInvalidProductSlug = errors.New("invalid product Slug: empty slug ")
