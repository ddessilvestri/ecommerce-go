package tools

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

func DateMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func CreateAPIResponse(status int, body string) *events.APIGatewayProxyResponse {

	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
