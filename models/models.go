package models

type SecretRDSJson struct {
	Username            string `json: "username"`
	Password            string `json: "password"`
	Engine              string `json: "engine"`
	Host                string `json: "host"`
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
