// auth/auth_validation.go
package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ddessilvestri/ecommerce-go/models"
)

type tokenJSON struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ExtractAuthUser(header map[string]string) (*models.AuthUser, error) {
	rawAuth := header["authorization"]
	if rawAuth == "" {
		return nil, errors.New("missing authorization header")
	}

	var token string
	if strings.HasPrefix(strings.ToLower(rawAuth), "bearer ") {
		token = rawAuth[len("Bearer "):]
	} else {
		token = rawAuth
	}

	// Ahora validamos el token
	isOk, uuid, err := tokenValidation(token)
	if !isOk {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("invalid token")
	}

	// Creamos y devolvemos el AuthUser
	return &models.AuthUser{
		UUID: uuid,
	}, nil
}

func tokenValidation(token string) (bool, string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("Invalid Token. Length != 3")
		return false, "Invalid Token", nil
	}

	rawPayload := parts[1]
	payloadBytes, err := base64.RawURLEncoding.DecodeString(rawPayload)
	if err != nil {
		fmt.Println("Decoding token error:", err.Error())
		return false, err.Error(), err
	}

	var tkj tokenJSON
	if err := json.Unmarshal(payloadBytes, &tkj); err != nil {
		fmt.Println("Cannot be decoded as JSON struct:", err.Error())
		return false, err.Error(), err
	}

	now := time.Now()
	expiration := time.Unix(int64(tkj.Exp), 0)
	if expiration.Before(now) {
		fmt.Println("Expiration date taken =", expiration.String())
		fmt.Println("Expired Token")
		return false, "Expired Token", nil
	}

	return true, string(tkj.Username), err
}
