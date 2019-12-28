package model

type Book struct {
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
	BookID int `json:"book_id" bson:"book_id"`
	Status int `json:"status" bson:"status"`
	Quality int `json:"quality" bson:"quality"`
	TotalBorrow int `json:"total_borrow" bson:"total_borrow"`
	Description string `json:"description" bson:"description"`
	Author string `json:"author" bson:"author"`
	Image string `json:"image" bson:"image"`
}