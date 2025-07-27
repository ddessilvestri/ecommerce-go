package models

type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DBClusterIdentifier string `json:"dbClusterIdentifier"`
}

type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}

type Category struct {
	CategID   int    `json:"categID"`
	CategName string `json:"categName"`
	CategPath string `json:"categPath"`
}

type Product struct {
	Id          int     `json:"prodID"`
	Title       string  `json:"prodTitle"`
	Description string  `json:"prodDescription"`
	CreatedAt   string  `json:"prodCreatedAt"`
	Updated     string  `json:"prodUpdated"`
	Price       float64 `json:"prodPrice,omitempty"`
	Stock       int     `json:"prodStock"`
	CategId     int     `json:"prodCategId"`
	Path        string  `json:"prodPath"`
	Search      string  `json:"search,omitempty"`
	CategPath   string  `json:"categPath,omitempty"`
}

type Address struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Phone      string `json:"phone"`
}

type OrdersDetails struct {
	Id       int     `json:"id"`
	OrderId  int     `json:"orderId"`
	ProdId   int     `json:"prodId"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Orders struct {
	Id           int     `json:"orderId"`
	UserUUID     string  `json:"orderUserUUID"`
	AddId        int     `json:"orderAddId"`
	Date         string  `json:"orderDate"`
	Total        float64 `json:"orderTotal"`
	OrderDetails []OrdersDetails
}

type User struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    int    `json:"status"`
	DateAdd   string `json:"dateAdd"`
	DateUpg   string `json:"dateUpg"`
}
