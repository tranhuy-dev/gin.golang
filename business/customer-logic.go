package business

import (
	"github.com/gin-gonic/gin"
	"gin.api/model"
	"gin.api/constant"
	"gin.api/database"
)
func GetAllCustomer(c *gin.Context) {
	c.JSON(Constant.HTTP_STATUS_OKE, model.Response{
		Message: "Retrieve Success",
		StatusCode: Constant.HTTP_STATUS_OKE,
	})
}

func GetDetailCustomer(c *gin.Context) {
	name := c.Param("name")
	c.JSON(Constant.HTTP_STATUS_OKE , model.Response{
		Message: "Hello " + name,
		StatusCode: Constant.HTTP_STATUS_OKE,
	})
}

func CreateCustomer(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(Constant.HTTP_STATUS_BAD_REQUEST , model.Response{
			Message: "Bad Request",
			StatusCode: Constant.HTTP_STATUS_BAD_REQUEST,
		})
	}
	_,errorInsert := database.AddCustomer(customer)
	if(errorInsert != nil) {
		c.JSON(Constant.HTTP_STATUS_BAD_REQUEST , model.Response{
			Message: errorInsert.Error(),
			StatusCode: Constant.HTTP_STATUS_BAD_REQUEST,
		})
	}
	c.JSON(Constant.HTTP_STATUS_CREATED , model.Response{
		Message: "Add Success",
		StatusCode: Constant.HTTP_STATUS_CREATED,
	})
}