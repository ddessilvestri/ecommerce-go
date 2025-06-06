package routers

type EntityRouter interface {
	Get(user string, id string, query map[string]string) (int, string)
	Post(body string, user string) (int, string)
	Put(body string, user string, id string) (int, string)
	Delete(user string, id string) (int, string)
}
