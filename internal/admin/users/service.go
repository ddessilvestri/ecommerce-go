package adminusers

import (
	"errors"

	"github.com/ddessilvestri/ecommerce-go/models"
)

type Service struct {
	repo Storage
}

func NewService(repo Storage) *Service {
	return &Service{repo: repo}
}

func (s *Service) Delete(uuid string) error {
	if uuid == "" {
		return ErrInvalidUUID
	}

	return s.repo.Delete(uuid)
}

func (s *Service) GetAll(page, limit int, sortBy, order string) ([]models.User, error) {
	offset := (page - 1) * limit
	return s.repo.GetAll(offset, limit, sortBy, order)
}

var ErrInvalidUUID = errors.New("invalid user id")
