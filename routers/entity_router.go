package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/models"
)

type EntityRouter interface {
	Get(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse
	Post(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse
	Put(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse
	Delete(requestWithContext models.RequestWithContext) *events.APIGatewayProxyResponse
}
