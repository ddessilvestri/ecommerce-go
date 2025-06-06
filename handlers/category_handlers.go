package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/ddessilvestri/ecommerce-go/db"
	"github.com/ddessilvestri/ecommerce-go/models"
)

func PostCategory(body string, User string) (int, string) {
	var t models.Category
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "data requested Error " + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "Must specify category name (Title)"
	}

	if len(t.CategPath) == 0 {
		return 400, "Must specify category Path (Route)"
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := db.InsertCategory(t)
	if err2 != nil {
		return 400, "Error while trying to add a category " + t.CategName + " > " + err2.Error()
	}

	return 200, "{CategID: " + strconv.Itoa(int(result)) + "}"
}
