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
func (h *Handler) Post(requestWithContext models.RequestWithContext) (*events.APIGatewayProxyResponse, error) {
	var c models.Category
	body := requestWithContext.RequestBody()

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
func (h *Handler) Put(requestWithContext models.RequestWithContext) (*events.APIGatewayProxyResponse, error) {
	var c models.Category
	body := requestWithContext.RequestBody()

	// 1. Try to parse the incoming JSON
	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error()), nil
	}

	// 2. Try to parse the incoming id
	id := requestWithContext.RequestPathParameters()["id"]
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
	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"Updated CategID": %d}`, idn)), nil
}

// Post handles the HTTP DELETE request to delete a category
func (h *Handler) Delete(requestWithContext models.RequestWithContext) (*events.APIGatewayProxyResponse, error) {

	// 1. Try to parse the incoming id
	id := requestWithContext.RequestPathParameters()["id"]
	idn, err := strconv.Atoi(id)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid CategoryId: "+err.Error()), nil
	}

	// 2. Call service to update category
	err = h.service.DeleteCategory(idn)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error()), nil
	}

	// 3. Return success response
	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"Deleted CategID": %d}`, idn)), nil
}

// Post handles the HTTP DELETE request to delete a category
func (h *Handler) Get(requestWithContext models.RequestWithContext) (*events.APIGatewayProxyResponse, error) {

	// 1 - First check if id is a query string parameter
	idstr := requestWithContext.RequestQueryStringParameters()["id"]
	slug := requestWithContext.RequestQueryStringParameters()["slug"]
	if idstr != "" {

		// 1. Try to parse the incoming idstr
		idn, err := strconv.Atoi(idstr)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid CategoryId: "+err.Error()), nil
		}

		// 2. Call service to update category
		c, err := h.service.GetCategory(idn)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error()), nil
		}

		body, err := json.Marshal(c)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusBadRequest, "Error converting model to json body: "+err.Error()), nil
		}

		// 3. Return success response
		return tools.CreateAPIResponse(http.StatusOK, string(body)), nil

	} else if slug != "" {
		// 2 - Second check if slug is a query string parameter

		// 1. Call service to update category
		c, err := h.service.GetCategoryBySlug(slug)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error()), nil
		}

		body, err := json.Marshal(c)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusBadRequest, "Error converting model to json body: "+err.Error()), nil
		}

		// 3. Return success response
		return tools.CreateAPIResponse(http.StatusOK, string(body)), nil

	}

	// 3  - Third retrieve all rows
	categories, err := h.service.GetCategories()
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error()), nil
	}

	body, err := json.Marshal(categories)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error converting model to json body: "+err.Error()), nil
	}

	// 3. Return success response
	return tools.CreateAPIResponse(http.StatusOK, string(body)), nil

}
