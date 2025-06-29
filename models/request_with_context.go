package models

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type RequestWithContext struct {
	req events.APIGatewayV2HTTPRequest
	ctx context.Context
}

func NewRequestWithContext(req events.APIGatewayV2HTTPRequest, ctx context.Context) RequestWithContext {
	return RequestWithContext{
		req: req,
		ctx: ctx,
	}
}

func (r RequestWithContext) Context() context.Context {
	return r.ctx
}

func (r RequestWithContext) Request() events.APIGatewayV2HTTPRequest {
	return r.req
}

func (r RequestWithContext) RequestBody() string {
	return r.req.Body
}

func (r RequestWithContext) RequestQueryStringParameters() map[string]string {
	return r.req.QueryStringParameters
}

func (r RequestWithContext) RequestPathParameters() map[string]string {
	return r.req.PathParameters
}
