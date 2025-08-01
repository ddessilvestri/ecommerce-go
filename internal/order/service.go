package order

import (
	"errors"
	"fmt"

	"github.com/ddessilvestri/ecommerce-go/models"
)

type Service struct {
	repo Storage
}

func NewService(repo Storage) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(o models.Orders) (int64, error) {
	fmt.Printf("DEBUG: Received order with Total: %v (type: %T)\n", o.Total, o.Total)

	if o.Total <= 0 {
		return 0, errors.New("order total must be > 0")
	}
	if o.UserUUID == "" {
		return 0, errors.New("user UUID must be provided")
	}
	if o.AddId <= 0 {
		return 0, errors.New("address ID must be provided")
	}
	if len(o.OrderDetails) == 0 {
		return 0, errors.New("order must have at least one order detail")
	}
	for _, detail := range o.OrderDetails {
		if detail.ProdId <= 0 {
			return 0, errors.New("product ID must be provided")
		}
		if detail.Quantity <= 0 {
			return 0, errors.New("quantity must be greater than 0")
		}
		if detail.Price <= 0 {
			return 0, errors.New("price must be greater than 0")
		}
	}

	return s.repo.Insert(o)
}

func (s *Service) GetAllByUserUUID(page, limit int, fromDate, toDate string, userUUID string) ([]models.Orders, error) {
	return s.repo.GetAllByUserUUID(page, limit, fromDate, toDate, userUUID)
}

func (s *Service) GetById(id int) (models.Orders, error) {
	return s.repo.GetById(id)
}

func (s *Service) GetByIdWithUserValidation(id int, userUUID string) (models.Orders, error) {
	order, err := s.repo.GetById(id)
	if err != nil {
		return models.Orders{}, err
	}

	// Validate that the order belongs to the authenticated user
	if order.UserUUID != userUUID {
		return models.Orders{}, errors.New("order not found or access denied")
	}

	return order, nil
}

func (s *Service) Update(o models.Orders) error {
	return s.repo.Update(o)
}

func (s *Service) Delete(id int, userUUID string) error {
	return s.repo.Delete(id, userUUID)
}
