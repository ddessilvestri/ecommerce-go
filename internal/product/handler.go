package product

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

func (h *Handler) Post(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {

	var c models.Product
	body := request.Body

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

func (h *Handler) Put(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {

	var c models.Product
	body := request.Body

	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid JSON body: "+err.Error())
	}

	id := request.PathParameters["id"]
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
func (h *Handler) Delete(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {

	id := request.PathParameters["id"]
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

func (h *Handler) Get(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {

	return tools.CreateAPIResponse(http.StatusMethodNotAllowed, "Not implemented")
	//query := request.QueryStringParameters

	// idstr := query["id"]

	// if idstr != "" {
	// 	id, err := strconv.Atoi(idstr)
	// 	if err != nil || id < 0 {
	// 		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid id parameter")
	// 	}

	// 	product, err := h.service.GetById(id)
	// 	if err != nil {
	// 		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid getting product :"+err.Error())
	// 	}
	// 	body, err := json.Marshal(product)
	// 	if err != nil {
	// 		return tools.CreateAPIResponse(http.StatusBadRequest, "Error converting model to json body: "+err.Error())

	// 	}
	// 	return tools.CreateAPIResponse(http.StatusOK, string(body))
	// }

	// page := 1 // default value 1
	// val := query["page"]
	// if val != "" {
	// 	parsed, err := strconv.Atoi(val)
	// 	if err != nil || parsed < 1 {
	// 		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid page parameter")
	// 	}
	// 	page = parsed
	// }

	// limit := 10 // default value 10
	// val = query["limit"]
	// if val != "" {
	// 	parsed, err := strconv.Atoi(val)
	// 	if err != nil || parsed < 1 {
	// 		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid limit parameter")
	// 	}
	// 	limit = parsed
	// }

	// sortBy := "id" // default value id sort by Product Id
	// val = query["sort_by"]
	// if val != "" {
	// 	allowed := map[string]bool{
	// 		"id":          true, // default value
	// 		"title":       true,
	// 		"description": true,
	// 		"price":       true,
	// 		"category_id": true,
	// 		"stock":       true,
	// 		"created_at":  true,
	// 	}
	// 	if !allowed[val] {
	// 		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid sort by parameter")
	// 	}

	// 	sortBy = val

	// }

	// order := "ASC" // default value is ASC
	// val = query["order"]
	// if val != "" {
	// 	upper := strings.ToUpper(val)
	// 	if upper != "ASC" && upper != "DESC" {
	// 		return tools.CreateAPIResponse(http.StatusBadRequest, "Invalid order parameter")

	// 	}
	// 	order = upper
	// }

	// products, err := h.service.GetAll(page, limit, sortBy, order)
	// if err != nil {
	// 	return tools.CreateAPIResponse(http.StatusBadRequest, "Error : "+err.Error())
	// }

	// body, err := json.Marshal(products)
	// if err != nil {
	// 	return tools.CreateAPIResponse(http.StatusBadRequest, "Error converting model to json body: "+err.Error())

	// }

	// return tools.CreateAPIResponse(http.StatusOK, string(body))
}
