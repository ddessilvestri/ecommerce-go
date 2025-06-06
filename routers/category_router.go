package routers

import (
	"github.com/ddessilvestri/ecommerce-go/handlers"
)

type CategoryRouter struct{}

func (router *CategoryRouter) Post(body string, user string) (int, string) {
	return handlers.CreateCategory(body, user)
}

func (router *CategoryRouter) Get(user string, id string, query map[string]string) (int, string) {
	return 405, GET + " Not implemented"
}

func (router *CategoryRouter) Put(body string, user string, id string) (int, string) {
	return 405, PUT + " Not implemented"
}

func (router *CategoryRouter) Delete(user string, id string) (int, string) {
	return 405, DELETE + " Not implemented"
}
