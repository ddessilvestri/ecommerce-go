package user

import (
	"database/sql"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/tools"
)

// Router struct contains all dependencies
type Router struct {
	handler *Handler
}

func NewRouter(db *sql.DB) *Router {
	repo := NewSQLRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	return &Router{handler: handler}
}

// Implements the EntityRouter interface

func (r *Router) Post(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return tools.CreateAPIResponse(http.StatusMethodNotAllowed, "not implemented")
}

// Future implementations (stubs for now)
func (r *Router) Get(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return r.handler.Get(request)
}

func (r *Router) Put(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return r.handler.Put(request)
}

func (r *Router) Delete(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return tools.CreateAPIResponse(http.StatusMethodNotAllowed, "not implemented")
}
