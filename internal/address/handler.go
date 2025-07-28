package address

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/models"
	"github.com/ddessilvestri/ecommerce-go/tools"

	authContext "github.com/ddessilvestri/ecommerce-go/auth/context"
)

// Handler struct wires the service (depends on Service)
type Handler struct {
	service *Service
}

// NewHandler creates a new handler with injected service
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Post(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {

	var a models.Address
	body := requestWithContext.RequestBody()

	err := json.Unmarshal([]byte(body), &a)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}
	userUUID, err := authContext.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "User Not found In Context: "+err.Error())
	}

	id, err := h.service.Create(a, userUUID)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error())
	}

	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"AddressID": %d}`, id))
}

func (h *Handler) Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	id := requestWithContext.RequestPathParameters()["id"]
	idn, err := strconv.Atoi(id)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid AddressId: "+err.Error())
	}

	var a models.Address
	body := requestWithContext.RequestBody()

	err = json.Unmarshal([]byte(body), &a)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	a.Id = idn

	err = h.service.Update(a)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error())
	}

	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"Updated AddressId": %d}`, idn))
}

// Delete handles the HTTP DELETE request to delete an address
func (h *Handler) Delete(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {

	id := requestWithContext.RequestPathParameters()["id"]
	idn, err := strconv.Atoi(id)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid AddressId: "+err.Error())
	}

	err = h.service.Delete(idn)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error())
	}
	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"Deleted AddressId": %d}`, idn))

}

func (h *Handler) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	query := requestWithContext.RequestQueryStringParameters()

	// === 1. Lookup by ID ===
	if idStr := query["id"]; idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid 'id' parameter")
		}
		address, err := h.service.GetById(id)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusNotFound, "Address not found: "+err.Error())
		}
		body, err := json.Marshal(address)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusInternalServerError, "Error converting to JSON: "+err.Error())
		}
		return tools.CreateAPIResponse(http.StatusOK, string(body))
	}

	// === 2. Default: Get all addresses for user ===
	userUUID, err := authContext.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "User Not found In Context: "+err.Error())
	}

	addresses, err := h.service.GetAllByUserUUID(userUUID)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, err.Error())
	}
	body, err := json.Marshal(addresses)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, "Error converting to JSON: "+err.Error())
	}
	return tools.CreateAPIResponse(http.StatusOK, string(body))
}
