package main

import (
	"context"
	"strings"

	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/awsgo"
	"github.com/ddessilvestri/ecommerce-go/db"
	"github.com/ddessilvestri/ecommerce-go/routers"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.AWSInit()
	if !isValid() {
		panic("Paramater Error. Must send 'SecretName','UrlPrefix'")
	}
	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	db.ReadSecret()

	//
	// status, message := handlers.Handlers(path, method, body, header, request)
	status, message := routers.Router(path, method, body, header, request)

	headerResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headerResp,
	}
	return res, nil

}

func isValid() bool {
	_, hasParam := os.LookupEnv("SecretName")
	if !hasParam {
		return hasParam
	}

	_, hasParam = os.LookupEnv("UrlPrefix")
	return hasParam
}
