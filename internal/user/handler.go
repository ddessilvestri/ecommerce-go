package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/ddessilvestri/ecommerce-go/auth/context"
	"github.com/ddessilvestri/ecommerce-go/models"
	"github.com/ddessilvestri/ecommerce-go/tools"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	var u models.User
	body := requestWithContext.RequestBody()

	err := json.Unmarshal([]byte(body), &u)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "invalid JSON body: "+err.Error())
	}

	u.UUID, err = context.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateAPIResponse(http.StatusUnauthorized, "user not found in context")
	}
	err = h.service.Update(u)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "error: "+err.Error())
	}

	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf("user updated: %s", u.UUID))
}

func (h *Handler) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	// userUUID, err := authctx.UserUUIDFromContext(requestWithContext.Context())

	userUUID, err := context.UserUUIDFromContext(requestWithContext.Context())
	if err != nil {
		return tools.CreateAPIResponse(http.StatusUnauthorized, "user not found in context")
	}
	user, err := h.service.GetByUUID(userUUID)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusNotFound, "user not found: "+err.Error())
	}
	body, err := json.Marshal(user)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, "error converting to JSON: "+err.Error())
	}
	return tools.CreateAPIResponse(http.StatusOK, string(body))
}
