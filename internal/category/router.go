package category

import (
	"database/sql"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/models"
)

// Router struct contains all dependencies
type Router struct {
	handler *Handler
}

// NewCategoryRouter sets up the repository, service, and handler
func NewRouter(db *sql.DB) *Router {
	repo := NewSQLRepository(db)
	service := NewCategoryService(repo)
	handler := NewCategoryHandler(service)
	return &Router{handler: handler}
}

// Implements the EntityRouter interface

func (r *Router) Post(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Post(requestWithContext)
	return resp
}

// Future implementations (stubs for now)
func (r *Router) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Get(requestWithContext)
	return resp
}

func (r *Router) Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Put(requestWithContext)
	return resp
}

func (r *Router) Delete(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	resp, _ := r.handler.Delete(requestWithContext)
	return resp
}
