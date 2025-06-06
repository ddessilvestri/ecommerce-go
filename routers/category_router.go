package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/handlers"
)

type CategoryRouter struct{}

func (router *CategoryRouter) Route(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return handlers.PostCategory(body, user)
	default:
		return 400, "Invalid Method " + method
	}

}
