
package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	"gin.api/constant"
	"gin.api/business"
)


func main() {
	fmt.Println("hello world")
	r := gin.Default()
	customerController := r.Group(Constant.API_URL + "customers/")
	{
		customerController.GET("" , business.GetAllCustomer)
		customerController.GET(":id" , business.GetDetailCustomer)
		customerController.POST("" , business.CreateCustomer)
	}

	bookController := r.Group(Constant.API_URL + "books/")
	{
		bookController.GET("" , business.GetAllBook)
		bookController.GET(":id" , business.GetDetailBook)
		bookController.POST("" , business.CreateBook)
	}
	r.Run()
}