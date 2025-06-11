package category

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
func NewCategoryHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Post handles the HTTP POST request to create a category
func (h *Handler) Post(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	var c models.Category
	body := request.Body

	// 1. Try to parse the incoming JSON
	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error()), nil
	}

	// 2. Call service to create category
	id, err := h.service.CreateCategory(c)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error()), nil
	}

	// 3. Return success response
	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"CategID": %d}`, id)), nil
}

// Post handles the HTTP POST request to create a category
func (h *Handler) Put(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	var c models.Category
	body := request.Body

	// 1. Try to parse the incoming JSON
	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error()), nil
	}

	// 2. Try to parse the incoming id
	id := request.PathParameters["id"]
	idn, err := strconv.Atoi(id)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid CategoryId: "+err.Error()), nil
	}

	c.CategID = idn

	// 3. Call service to update category
	err = h.service.UpdateCategory(c)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error()), nil
	}

	// 3. Return success response
	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"CategID": %d}`, idn)), nil
}
