// routers/entity_router.go
package routers

import "github.com/aws/aws-lambda-go/events"

type EntityRouter interface {
	Get(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse
	Post(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse
	Put(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse
	Delete(request events.APIGatewayV2HTTPRequest) *events.APIGatewayProxyResponse
}
