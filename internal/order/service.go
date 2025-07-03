package order

import (
	"errors"

	"github.com/ddessilvestri/ecommerce-go/models"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(o models.Orders) (int, error) {
	if o.Total <= 0 {
		return 0, errors.New("order total must be > 0")
	}
	return s.repo.Insert(o)
}

func (s *Service) GetById(id int, userUUID string) (models.Orders, error) {
	return s.repo.GetById(id, userUUID)
}

func (s *Service) GetByUser(userUUID string) ([]models.Orders, error) {
	return s.repo.GetByUserUUID(userUUID)
}

func (s *Service) Update(o models.Orders) error {
	return s.repo.Update(o)
}

func (s *Service) Delete(id int, userUUID string) error {
	return s.repo.Delete(id, userUUID)
}
