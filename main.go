
package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	"gin.api/model"
)


func main() {
	fmt.Println("hello world")
	r := gin.Default()
	r.GET("api/v1" , getData)

	r.Run()
}

func getData(c *gin.Context) {
	c.JSON(200, model.Response{
		Message: "Api is working",
		StatusCode: 200,
	})
}