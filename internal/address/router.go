package address

import (
	"database/sql"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/models"
)

type Router struct {
	handler *Handler
}

func NewRouter(db *sql.DB) *Router {
	repo := NewSQLRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	return &Router{handler: handler}
}

func (r *Router) Post(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	return r.handler.Post(requestWithContext)
}

func (r *Router) Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	return r.handler.Get(requestWithContext)
}

func (r *Router) Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	return r.handler.Put(requestWithContext)
}

func (r *Router) Delete(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse {
	return r.handler.Delete(requestWithContext)
}
