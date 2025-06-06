package routers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ddessilvestri/ecommerce-go/auth"
	"github.com/ddessilvestri/ecommerce-go/handlers"
)

func Router(path string, method string, body string, header map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Processing " + path + " > " + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := authValidation(path, method, header)

	if !isOk {
		return statusCode, user
	}

	switch path {
	case "/user":
		return ProcessUser(body, path, method, user, id, request)
	case "/product":
		return ProcessProducts(body, path, method, user, idn, request)
	case "/stock":
		return ProcessStock(body, path, method, user, idn, request)
	case "/address":
		return ProcessAddress(body, path, method, user, idn, request)
	case "/category":
		return ProcessCategory(body, path, method, user, idn, request)
	case "/order":
		return ProcessOrder(body, path, method, user, idn, request)

	}

	return 400, "Invalid Method"

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

// func authValidation(path string, method string, header map[string]string) (bool, int, string) {
// 	if (path == "product" && method == "GET") ||
// 		(path == "category" && method == "GET") {
// 		return true, 200, ""
// 	}
// 	rawAuth := header["authorization"] // e.g. "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9…"
// 	if len(rawAuth) == 0 {
// 		return false, 401, "Required Token"
// 	}

// 	// 1) Separar “Bearer” del JWT real
// 	parts := strings.SplitN(rawAuth, " ", 2)
// 	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
// 		return false, 401, "Invalid Authorization header format"
// 	}
// 	token := parts[1] // aquí ya tenemos solo “eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9…”

// 	isOk, msg, err := auth.TokenValidation(token)
// 	if !isOk {
// 		if err != nil {
// 			fmt.Println("Token Error " + err.Error())
// 			return false, 401, err.Error()
// 		}
// 		fmt.Println("Token Error " + msg)
// 		return false, 401, msg
// 	}

// 	fmt.Println("Token OK")
// 	return true, 200, msg

// }

func ProcessUser(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
func ProcessProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
func ProcessCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	switch method {
	case "POST":
		return handlers.PostCategory(body, user)
	}

	return 400, "Invalid Method"
}
func ProcessStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
func ProcessAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
func ProcessOrder(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
