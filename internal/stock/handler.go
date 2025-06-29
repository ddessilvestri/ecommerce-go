package stock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/models"
	"github.com/ddessilvestri/ecommerce-go/tools"
)

// Handler struct wires the service (depends on Service)
type Handler struct {
	service *Service
}

// NewCategoryHandler creates a new handler with injected service
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {

	type StockUpdate struct {
		Delta int `json:"delta"`
	}

	var stockUpdate StockUpdate
	body := requestWithContext.Request().Body

	err := json.Unmarshal([]byte(body), &stockUpdate)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	pId := requestWithContext.RequestPathParameters()["productId"]
	pIdn, err := strconv.Atoi(pId)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid ProductId: "+err.Error())
	}

	err = h.service.UpdateStock(pIdn, stockUpdate.Delta)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error())
	}

	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"Stock incremented in %d for ProductId": %d}`, stockUpdate.Delta, pIdn))
}
