package model

type Book struct {
	Name string `json:"name"`
	Type string `json:"type"`
	BookID int `json:"book_id"`
	Status int `json:"status"`
	Owner int `json:"owner"`
	BorrowDate int `json:"borrow_date"`
	PayDate int `json:"pay_date"`
	Quality int `json:"quality"`
	TotalBorrow int `json:"total_borrow"`
	Description string `json:"description"`
	Author string `json:"author"`
	Image string `json:"image"`
}