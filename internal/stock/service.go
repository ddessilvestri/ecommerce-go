package stock

import (
	"errors"
)

// Service provides methods for business logic related to category.
type Service struct {
	repo Storage // This is the interface, so it's decoupled from repositorySQL
}

func NewService(repo Storage) *Service {
	return &Service{repo: repo}
}

func (s *Service) UpdateStock(productId, delta int) error {
	if productId < 1 {
		return ErrInvalidProductId
	}
	return s.repo.UpdateStock(productId, delta)

}

var ErrInvalidProductId = errors.New("invalid product Id: Id < 1 ")
var ErrInvalidStock = errors.New("invalid stock value")
