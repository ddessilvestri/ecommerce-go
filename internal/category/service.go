package category

import (
	"errors"

	"github.com/ddessilvestri/ecommerce-go/models"
)

// Service provides methods for business logic related to category.
type Service struct {
	repo Storage // This is the interface, so it's decoupled from repositorySQL
}

// NewCategoryService acts like a constructor.
func NewCategoryService(repo Storage) *Service {
	return &Service{repo: repo}
}

// CreateCategory performs validation and delegates to the repository.
func (s *Service) CreateCategory(c models.Category) (int64, error) {
	// Simple validation (can be more elaborate in real use cases)
	if c.CategName == "" || c.CategPath == "" {
		return 0, ErrInvalidCategory
	}

	// Business logic: could include auditing, formatting, etc.
	return s.repo.InsertCategory(c)
}

// CreateCategory performs validation and delegates to the repository.
func (s *Service) UpdateCategory(c models.Category) error {
	// Simple validation (can be more elaborate in real use cases)
	if c.CategName == "" || c.CategPath == "" {
		return ErrInvalidCategory
	}
	if c.CategID < 1 {
		return ErrInvalidCategoryId
	}
	// Business logic: could include auditing, formatting, etc.
	return s.repo.UpdateCategory(c)

}

// DeleteCategory performs validation and delegates to the repository.
func (s *Service) DeleteCategory(id int) error {
	// Simple validation (can be more elaborate in real use cases)
	if id < 1 {
		return ErrInvalidCategoryId
	}
	// Business logic: could include auditing, formatting, etc.
	return s.repo.DeleteCategory(id)

}

// CreateCategory performs validation and delegates to the repository.
func (s *Service) GetCategory(id int) (models.Category, error) {

	// Simple validation (can be more elaborate in real use cases)
	if id < 1 {
		return models.Category{}, ErrInvalidCategoryId
	}

	// Business logic: could include auditing, formatting, etc.

	return s.repo.GetCategory(id)

}

// ErrInvalidCategory represents a validation error.
var ErrInvalidCategory = errors.New("invalid category: name and path are required")
var ErrInvalidCategoryId = errors.New("invalid category Id: Id < 1 ")
