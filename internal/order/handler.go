package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	authContext "github.com/ddessilvestri/ecommerce-go/auth/context"
	"github.com/ddessilvestri/ecommerce-go/models"
	"github.com/ddessilvestri/ecommerce-go/tools"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Post(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	var o models.Orders
	body := requestWithContext.RequestBody()

	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	userUUID, err := authContext.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateAPIResponse(http.StatusUnauthorized, "User not found in context: "+err.Error())
	}

	o.UserUUID = userUUID

	id, err := h.service.Create(o)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, "Error creating order: "+err.Error())
	}

	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"OrderId": %d}`, id))
}

func (h *Handler) Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	idStr := requestWithContext.RequestPathParameters()["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid OrderId: "+err.Error())
	}

	var o models.Orders
	body := requestWithContext.RequestBody()

	err = json.Unmarshal([]byte(body), &o)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	userUUID, err := authContext.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateAPIResponse(http.StatusUnauthorized, "User not found in context: "+err.Error())
	}

	o.Id = id
	o.UserUUID = userUUID

	err = h.service.Update(o)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, "Error updating order: "+err.Error())
	}

	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"UpdatedOrderId": %d}`, id))
}

func (h *Handler) Delete(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	idStr := requestWithContext.RequestPathParameters()["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid OrderId: "+err.Error())
	}

	userUUID, err := authContext.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateAPIResponse(http.StatusUnauthorized, "User not found in context: "+err.Error())
	}

	err = h.service.Delete(id, userUUID)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, "Error deleting order: "+err.Error())
	}

	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"DeletedOrderId": %d}`, id))
}

func (h *Handler) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	userUUID, err := authContext.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateAPIResponse(http.StatusUnauthorized, "User not found in context: "+err.Error())
	}

	// === 1. Get By Id ===
	idStr := requestWithContext.RequestPathParameters()["id"]
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid OrderId: "+err.Error())
		}

		order, err := h.service.GetById(id)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusNotFound, "Order not found: "+err.Error())
		}

		body, err := json.Marshal(order)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusInternalServerError, "Error converting to JSON: "+err.Error())
		}

		return tools.CreateAPIResponse(http.StatusOK, string(body))
	}

	// === 2. Default: Get all paginated ===

	query := requestWithContext.RequestQueryStringParameters()
	page, fromDate, toDate, err := tools.ParseOrdersPaginationAndSorting(query)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, err.Error())
	}

	limit := 10 // Default limit

	orders, err := h.service.GetAllByUserUUID(page, limit, fromDate, toDate, userUUID)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, err.Error())
	}
	body, err := json.Marshal(orders)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, "Error converting to JSON: "+err.Error())
	}

	return tools.CreateAPIResponse(http.StatusOK, string(body))
}
