package model

type Customer struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName string `json:"last_name" bson:"last_name"`
	Address string `json:"address" bson:"address"`
	Status int `json:"status" bson:"status"`
	Reputation int `json:"reputation" bson:"reputation"`
	Image string `json:"image" bson:"image"`
	Books []int `json:"books" bson:"books"`
	Phone string `json:"phone" bson:"phone"`
	Friends []int `json:"friends" bson:"friends"`
	Password string `json:"password" bson:"password"`
	Token string `json:"token" bson:"token"`
	CustomerID int `json:"customer_id" bson:"customer_id"`
}

