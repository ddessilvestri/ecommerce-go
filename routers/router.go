package routers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/auth"
	authContext "github.com/ddessilvestri/ecommerce-go/auth/context"
	"github.com/ddessilvestri/ecommerce-go/models"

	"github.com/ddessilvestri/ecommerce-go/internal/category"
	"github.com/ddessilvestri/ecommerce-go/internal/product"
	"github.com/ddessilvestri/ecommerce-go/internal/stock"
	"github.com/ddessilvestri/ecommerce-go/tools"
)

// HTTP method constants
const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// Router determines which entity router should handle the request.
func Router(request events.APIGatewayV2HTTPRequest, urlPrefix string, db *sql.DB) *events.APIGatewayProxyResponse {
	path := strings.Replace(request.RawPath, urlPrefix, "", 1)
	method := request.RequestContext.HTTP.Method
	header := request.Headers

	firstSegment := getFirstPathSegment(path)

	entityRouter, err := CreateRouter(firstSegment, db)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Unable to route request: "+err.Error())
	}

	var authUser *models.AuthUser

	if !(path == "product" && method == "GET") && !(path == "category" && method == "GET") {
		authUser, err = auth.ExtractAuthUser(header)
		if err != nil {
			return tools.CreateAPIResponse(http.StatusUnauthorized, "Unable to authenticate user: "+err.Error())
		}
	}

	context := authContext.WithUser(context.Background(), authUser)
	requestWithContext := models.NewRequestWithContext(request, context)

	switch method {
	case GET:
		return entityRouter.Get(requestWithContext)
	case POST:
		return entityRouter.Post(requestWithContext)
	case PUT:
		return entityRouter.Put(requestWithContext)
	case DELETE:
		return entityRouter.Delete(requestWithContext)
	default:
		return tools.CreateAPIResponse(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// CreateRouter maps entity names to their router implementations
func CreateRouter(entity string, db *sql.DB) (EntityRouter, error) {
	switch entity {
	case "category":
		return category.NewRouter(db), nil
	case "product":
		return product.NewRouter(db), nil
	case "stock":
		return stock.NewRouter(db), nil
	case "address":
		return product.NewRouter(db), nil
	default:
		return nil, fmt.Errorf("entity '%s' not implemented", entity)
	}
}

// Extract first part of path: "/category/1" -> "category"
func getFirstPathSegment(path string) string {
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}
