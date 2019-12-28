package business

import (
	Constant "gin.api/constant"
	"gin.api/database"
	"gin.api/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"log"
)

func GetAllCustomer(c *gin.Context) {
	customersData, errQuery := database.GetAllCustomer()
	if errQuery != nil {
		c.JSON(Constant.HTTP_STATUS_INTERNAL_ERROR, model.Response{
			Message:    "Retrieve Fail",
			StatusCode: Constant.HTTP_STATUS_INTERNAL_ERROR,
		})
	}

	c.JSON(Constant.HTTP_STATUS_OKE, model.Response{
		Message:    "Retrieve Success",
		StatusCode: Constant.HTTP_STATUS_OKE,
		Data:       customersData,
	})
}

func GetDetailCustomer(c *gin.Context) {
	id := c.Param("id")
	customerID , errParse := strconv.Atoi(id)
	if errParse != nil {
		log.Fatal(errParse)
	}
	dataCustomer , errQuery := database.GetDetailCustomer(customerID)
	if errQuery != nil {
		c.JSON(Constant.HTTP_STATUS_INTERNAL_ERROR, model.Response{
			Message:    "Get Fail ",
			StatusCode: Constant.HTTP_STATUS_INTERNAL_ERROR,
		})
	}
	c.JSON(Constant.HTTP_STATUS_OKE, model.Response{
		Message:    "Get Success ",
		StatusCode: Constant.HTTP_STATUS_OKE,
		Data: dataCustomer,
	})
}

func CreateCustomer(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(Constant.HTTP_STATUS_BAD_REQUEST, model.Response{
			Message:    "Bad Request",
			StatusCode: Constant.HTTP_STATUS_BAD_REQUEST,
		})
	}

	seqCount, errSeqCount := database.GetSequence("customer")
	if errSeqCount != nil {
		c.JSON(Constant.HTTP_STATUS_BAD_REQUEST, model.Response{
			Message:    errSeqCount.Error(),
			StatusCode: Constant.HTTP_STATUS_BAD_REQUEST,
		})
	}

	_, errorInsert := database.AddCustomer(customer, seqCount)
	if errorInsert != nil {
		c.JSON(Constant.HTTP_STATUS_BAD_REQUEST, model.Response{
			Message:    errorInsert.Error(),
			StatusCode: Constant.HTTP_STATUS_BAD_REQUEST,
		})
	}
	c.JSON(Constant.HTTP_STATUS_CREATED, model.Response{
		Message:    "Add Success",
		StatusCode: Constant.HTTP_STATUS_CREATED,
	})
}
