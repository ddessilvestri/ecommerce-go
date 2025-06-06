package routers

import (
	"fmt"

	"github.com/ddessilvestri/ecommerce-go/common"
)

func CreateRouter(entity string) (EntityRouter, error) {
	switch entity {
	case common.User:
		return nil, fmt.Errorf("invalid entity: %s", entity)
	case common.Product:
		return nil, fmt.Errorf("invalid entity: %s", entity)
	case common.Stock:
		return nil, fmt.Errorf("invalid entity: %s", entity)
	case common.Address:
		return nil, fmt.Errorf("invalid entity: %s", entity)
	case common.Category:
		return &CategoryRouter{}, nil
	case common.Order:
		return nil, fmt.Errorf("invalid entity: %s", entity)
	default:
		return nil, fmt.Errorf("invalid entity: %s", entity)

	}
}
