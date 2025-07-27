package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func (h *Handler) Post(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {

	var c models.Product
	body := requestWithContext.RequestBody()

	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	id, err := h.service.Create(c)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error())
	}

	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"ProductID": %d}`, id))
}

func (h *Handler) Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {

	var c models.Product
	body := requestWithContext.RequestBody()

	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	id := requestWithContext.RequestPathParameters()["id"]
	idn, err := strconv.Atoi(id)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid ProductId: "+err.Error())
	}

	c.Id = idn

	err = h.service.Update(c)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error())
	}

	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"Updated ProductId": %d}`, idn))
}

// Post handles the HTTP DELETE request to delete a category
func (h *Handler) Delete(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {

	id := requestWithContext.RequestPathParameters()["id"]
	idn, err := strconv.Atoi(id)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid ProductId: "+err.Error())
	}

	err = h.service.Delete(idn)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error())
	}
	return tools.CreateAPIResponse(http.StatusOK, fmt.Sprintf(`{"Deleted ProductId": %d}`, idn))

}

func (h *Handler) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	query := requestWithContext.RequestQueryStringParameters()

	// === 1. Lookup by ID ===
	if idStr := query["id"]; idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid 'id' parameter")
		}
		product, err := h.service.GetById(id)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusNotFound, "Product not found: "+err.Error())
		}
		body, err := json.Marshal(product)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusInternalServerError, "Error converting to JSON: "+err.Error())
		}
		return tools.CreateAPIResponse(http.StatusOK, string(body))
	}

	// === 2. Lookup by Slug ===
	if slug := strings.TrimSpace(query["slug"]); slug != "" {
		product, err := h.service.GetBySlug(slug)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusNotFound, "Product not found: "+err.Error())
		}
		body, err := json.Marshal(product)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusInternalServerError, "Error converting to JSON: "+err.Error())
		}
		return tools.CreateAPIResponse(http.StatusOK, string(body))
	}

	// === 3. Search full-text ===
	if search := strings.TrimSpace(query["search"]); search != "" {
		page, limit, sortBy, order, err := tools.ParsePaginationAndSorting(query)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusBadRequest, err.Error())
		}
		products, err := h.service.SearchByText(search, page, limit, sortBy, order)

		if err != nil {
			return tools.CreateAPIResponse(http.StatusInternalServerError, err.Error())
		}
		body, err := json.Marshal(products)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusInternalServerError, "Error converting to JSON: "+err.Error())
		}
		return tools.CreateAPIResponse(http.StatusOK, string(body))
	}

	// === 4. Filter by category ID ===
	if catIdStr := strings.TrimSpace(query["categId"]); catIdStr != "" {
		catId, err := strconv.Atoi(catIdStr)
		if err != nil || catId <= 0 {
			return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid 'categId' parameter")
		}
		products, err := h.service.GetByCategoryId(catId)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusInternalServerError, err.Error())
		}
		body, err := json.Marshal(products)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusInternalServerError, "Error converting to JSON: "+err.Error())
		}
		return tools.CreateAPIResponse(http.StatusOK, string(body))
	}

	// === 5. Filter by category slug ===
	if slugCateg := strings.TrimSpace(query["slugCateg"]); slugCateg != "" {
		products, err := h.service.GetByCategorySlug(slugCateg)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusInternalServerError, err.Error())
		}
		body, err := json.Marshal(products)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusInternalServerError, "Error converting to JSON: "+err.Error())
		}
		return tools.CreateAPIResponse(http.StatusOK, string(body))
	}

	// === 6. Default: Get all paginated ===
	page, limit, sortBy, order, err := tools.ParsePaginationAndSorting(query)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, err.Error())
	}

	products, err := h.service.GetAll(page, limit, sortBy, order)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, err.Error())
	}
	body, err := json.Marshal(products)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusInternalServerError, "Error converting to JSON: "+err.Error())
	}
	return tools.CreateAPIResponse(http.StatusOK, string(body))
}
