package category

import (
	"database/sql"

	"github.com/aws/aws-lambda-go/events"
)

// Router struct contains all dependencies
type Router struct {
	handler *Handler
}

// NewCategoryRouter sets up the repository, service, and handler
func NewCategoryRouter(db *sql.DB) *Router {
	repo := NewSQLRepository(db)
	service := NewCategoryService(repo)
	handler := NewCategoryHandler(service)
	return &Router{handler: handler}
}

// Implements the EntityRouter interface

func (r *Router) Post(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Post(request)
	return resp
}

// Future implementations (stubs for now)
func (r *Router) Get(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	return nil
}

func (r *Router) Put(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Put(request)
	return resp
}

func (r *Router) Delete(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Delete(request)
	return resp
}
