package model

type Customer struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string `json:"address"`
	Status int `json:"status"`
	Reputation int `json:"reputation"`
	Image string `json:"image"`
	Books []int `json:"books"`
	Phone string `json:"phone"`
	Friends []int `json:"friends"`
	Password string `json:"password"`
	Token string `json:"token"`
	CustomerID int `json:"customer_id"`
}

