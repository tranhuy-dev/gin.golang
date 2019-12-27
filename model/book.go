package model

type Book struct {
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
	BookID int `json:"book_id" bson:"book_id"`
	Status int `json:"status" bson:"status"`
	Owner int `json:"owner" bson:"owner"`
	BorrowDate int `json:"borrow_date" bson:"borrow_date"`
	PayDate int `json:"pay_date" bson:"pay_date"`
	Quality int `json:"quality" bson:"quality"`
	TotalBorrow int `json:"total_borrow" bson:"total_borrow"`
	Description string `json:"description" bson:"description"`
	Author string `json:"author" bson:"author"`
	Image string `json:"image" bson:"image"`
}