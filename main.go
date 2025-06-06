package main

import (
	"context"

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

	db.ReadSecret()

	status, message := routers.Router(request, os.Getenv("UrlPrefix"))

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
