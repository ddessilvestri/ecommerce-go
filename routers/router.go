package routers

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/auth"
)

func Router(request events.APIGatewayV2HTTPRequest, urlPrefix string) (int, string) {

	path := strings.Replace(request.RawPath, urlPrefix, "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	fmt.Println("Processing " + path + " > " + method)

	id := request.PathParameters["id"]
	// idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := authValidation(path, method, header)

	if !isOk {
		return statusCode, user
	}
	firstSegment := getFirstPathSegment(path)
	entityRouter, err := CreateRouter(firstSegment)
	if err != nil {
		return 400, "unable to get router " + err.Error()
	}

	switch method {
	case GET:
		return entityRouter.Get(user, id, request.QueryStringParameters)
	case PUT:
		return entityRouter.Put(body, user, id)
	case POST:
		return entityRouter.Post(body, user)
	case DELETE:
		return entityRouter.Delete(user, id)
	default:
		return 405, "Method not allowed"
	}

}

func getFirstPathSegment(path string) string {
	// Remove leading/trailing slashes
	trimmed := strings.Trim(path, "/")
	segments := strings.Split(trimmed, "/")
	if len(segments) > 0 && segments[0] != "" {
		return segments[0]
	}
	return ""
}

func authValidation(
	path string,
	method string,
	header map[string]string,
) (bool, int, string) {
	// Rutas públicas (sin token)
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
		return true, 200, ""
	}

	rawAuth := header["authorization"]
	if len(rawAuth) == 0 {
		return false, 401, "Required Token"
	}

	var token string
	// Si viene con "Bearer <espacio>" (minúsculas o mayúsculas), lo cortamos
	if strings.HasPrefix(strings.ToLower(rawAuth), "bearer ") {
		token = rawAuth[len("Bearer "):]
	} else {
		// No viene con prefijo, asumimos que rawAuth es directamente el token
		token = rawAuth
	}

	isOk, msg, err := auth.TokenValidation(token)
	if !isOk {
		if err != nil {
			fmt.Println("Token Error " + err.Error())
			return false, 401, err.Error()
		}
		fmt.Println("Token Error " + msg)
		return false, 401, msg
	}

	fmt.Println("Token OK")
	return true, 200, msg
}
