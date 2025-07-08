package adminusers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/ddessilvestri/ecommerce-go/models"
	"github.com/ddessilvestri/ecommerce-go/tools"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	// Verify if the user is admin

	query := requestWithContext.RequestQueryStringParameters()
	page, limit, sortBy, order, err := tools.ParsePaginationAndSorting(query)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, err.Error())
	}

	users, err := h.service.GetAll(page, limit, sortBy, order)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, err.Error())
	}
	body, err := json.Marshal(users)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
	}
	return tools.CreateAPIResponse(http.StatusOK, string(body))
}

func (h *Handler) Delete(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	// Verify if the user is admin

	uuid := requestWithContext.RequestPathParameters()["id"]
	if uuid == "" {
		return tools.CreateAPIResponse(http.StatusBadRequest, "invalid UUID")
	}

	err := h.service.Delete(uuid)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "error: "+err.Error())
	}

	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf("product deleted: %s", uuid))
}
