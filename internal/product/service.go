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

func (s *Service) Update(c models.Product) error {
	if c.Title == "" {
		return ErrInvalidProduct
	}
	if c.Id < 1 {
		return ErrInvalidProductId
	}
	return s.repo.Update(c)

}

func (s *Service) Delete(id int) error {
	if id < 1 {
		return ErrInvalidProductId
	}
	return s.repo.Delete(id)

}

func (s *Service) GetById(id int) (models.Product, error) {

	if id < 1 {
		return models.Product{}, ErrInvalidProductId
	}

	return s.repo.GetById(id)

}

func (s *Service) GetByCategoryId(id int) ([]models.Product, error) {

	if id < 1 {
		return []models.Product{}, ErrInvalidProductId
	}

	return s.repo.GetByCategoryId(id)

}

func (s *Service) GetAll(page, limit int, sortBy, order string) ([]models.Product, error) {
	offset := (page - 1) * limit
	return s.repo.GetAll(offset, limit, sortBy, order)
}

func (s *Service) GetBySlug(slug string) (models.Product, error) {

	if slug == "" {
		return models.Product{}, ErrInvalidProductSlug
	}

	return s.repo.GetBySlug(slug)

}

func (s *Service) GetByCategorySlug(slug string) ([]models.Product, error) {

	if slug == "" {
		return []models.Product{}, ErrInvalidProductSlug
	}

	return s.repo.GetByCategorySlug(slug)

}

func (s *Service) SearchByText(text string, page, limit int, sortBy, order string) ([]models.Product, error) {
	offset := (page - 1) * limit
	return s.repo.SearchByText(text, offset, limit, sortBy, order)
}

var ErrInvalidProduct = errors.New("invalid product: title is required")
var ErrInvalidProductId = errors.New("invalid product Id: Id < 1 ")
var ErrInvalidProductSlug = errors.New("invalid product Slug: empty slug ")
