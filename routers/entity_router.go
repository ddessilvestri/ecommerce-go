package routers

import "github.com/aws/aws-lambda-go/events"

type EntityRouter interface {
	Route(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string)
}
