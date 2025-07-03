package address

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

func (s *Service) Create(a models.Address, userUUID string) (int64, error) {
	if userUUID == "" {
		return 0, ErrMissingUUID
	}
	if a.Address == "" {
		return 0, ErrMissingAddress
	}
	if a.Name == "" {
		return 0, ErrMissingName
	}
	if a.Title == "" {
		return 0, ErrMissingTitle
	}
	if a.City == "" {
		return 0, ErrMissingCity
	}
	if a.Phone == "" {
		return 0, ErrMissingPhone
	}
	if a.PostalCode == "" {
		return 0, ErrMissingPostalCode
	}

	return s.repo.Insert(a, userUUID)
}

func (s *Service) Update(a models.Address) error {
	if a.Id < 1 {
		return ErrInvalidId
	}
	if !s.repo.Exists(a.Id) {
		return ErrIdNotFound
	}
	if a.Address == "" {
		return ErrMissingAddress
	}
	if a.Name == "" {
		return ErrMissingName
	}
	if a.Title == "" {
		return ErrMissingTitle
	}
	if a.City == "" {
		return ErrMissingCity
	}
	if a.Phone == "" {
		return ErrMissingPhone
	}
	if a.PostalCode == "" {
		return ErrMissingPostalCode
	}

	return s.repo.Update(a)

}

func (s *Service) Delete(id int) error {
	if id < 1 {
		return ErrInvalidId
	}
	if !s.repo.Exists(id) {
		return ErrIdNotFound
	}
	return s.repo.Delete(id)

}

func (s *Service) GetAllByUserUUID(userUUID string) ([]models.Address, error) {
	return s.repo.GetAllByUserUUID(userUUID)
}

var ErrMissingAddress = errors.New("missing address ")
var ErrMissingName = errors.New("missing name ")
var ErrMissingTitle = errors.New("missing title ")
var ErrMissingCity = errors.New("missing city ")
var ErrMissingPostalCode = errors.New("missing postal code ")
var ErrMissingPhone = errors.New("missing phone ")
var ErrMissingUUID = errors.New("missing UUID ")
var ErrIdNotFound = errors.New("missing Id ")
var ErrInvalidId = errors.New("invalid Id ")
