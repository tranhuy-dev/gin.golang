
package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	"gin.api/model"
	"gin.api/constant"
)


func main() {
	fmt.Println("hello world")
	r := gin.Default()
	r.GET("api/v1" , getData)
	r.GET("api/v1/:name" , getDataParams)
	r.POST("api/v1/" , createData)
	r.Run()
}

func getData(c *gin.Context) {
	c.JSON(Constant.HTTP_STATUS_OKE, model.Response{
		Message: "Api is working",
		StatusCode: Constant.HTTP_STATUS_OKE,
	})
}

func getDataParams(c *gin.Context) {
	name := c.Param("name")
	c.JSON(Constant.HTTP_STATUS_OKE , model.Response{
		Message: "Hello " + name,
		StatusCode: Constant.HTTP_STATUS_OKE,
	})
}

func createData(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(Constant.HTTP_STATUS_BAD_REQUEST , model.Response{
			Message: "Bad Request",
			StatusCode: Constant.HTTP_STATUS_BAD_REQUEST,
		})
	}
	c.JSON(Constant.HTTP_STATUS_CREATED , model.Response{
		Message: "Hello " + customer.Name,
		StatusCode: Constant.HTTP_STATUS_CREATED,
	})
}