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

	"github.com/ddessilvestri/ecommerce-go/internal/address"
	adminusers "github.com/ddessilvestri/ecommerce-go/internal/admin/users"
	"github.com/ddessilvestri/ecommerce-go/internal/category"
	"github.com/ddessilvestri/ecommerce-go/internal/order"
	"github.com/ddessilvestri/ecommerce-go/internal/product"
	"github.com/ddessilvestri/ecommerce-go/internal/stock"
	"github.com/ddessilvestri/ecommerce-go/internal/user"
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

	segments := getPathSegments(path)

	entityRouter, err := CreateRouter(segments, db)
	if err != nil {
		return tools.CreateAPIResponse(http.StatusBadRequest, "Unable to route request: "+err.Error())
	}

	var authUser *models.AuthUser

	if !(segments[0] == "product" && method == "GET") && !(segments[0] == "category" && method == "GET") {
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
func CreateRouter(segments []string, db *sql.DB) (EntityRouter, error) {
	switch segments[0] {
	case "category":
		return category.NewRouter(db), nil
	case "product":
		return product.NewRouter(db), nil
	case "stock":
		return stock.NewRouter(db), nil
	case "address":
		return address.NewRouter(db), nil
	case "order":
		return order.NewRouter(db), nil
	case "admin":
		if segments[1] == "users" {
			return adminusers.NewRouter(db), nil
		}
		return nil, fmt.Errorf("path '%s'/'%s' not implemented", segments[0], segments[1])
	case "user":
		return user.NewRouter(db), nil // Change API response from HTTP to With Context
	default:
		return nil, fmt.Errorf("entity '%s' not implemented", segments[0])
	}
}

// Extract first part of path: "/category/1" -> "category"

func getPathSegments(path string) []string {
	path = strings.Trim(path, "/")
	if path == "" {
		return []string{}
	}
	return strings.Split(path, "/")
}
