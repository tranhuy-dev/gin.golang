package business

import (
	"github.com/gin-gonic/gin"
	"gin.api/model"
	"gin.api/constant"
)

func CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(Constant.HTTP_STATUS_BAD_REQUEST , model.Response{
			Message: "Bad Request",
			StatusCode: Constant.HTTP_STATUS_BAD_REQUEST,
		})
	}

	c.JSON(Constant.HTTP_STATUS_CREATED , model.Response{
		Message: "Add book " + book.Name + " success",
		StatusCode: Constant.HTTP_STATUS_CREATED,
	})
}

func GetAllBook(c *gin.Context) {
	c.JSON(Constant.HTTP_STATUS_OKE , model.Response{
		Message: "Get Books Success",
		StatusCode: Constant.HTTP_STATUS_OKE,
	})
}

func GetDetailBook(c *gin.Context) {
	name := c.Param("name")
	c.JSON(Constant.HTTP_STATUS_OKE , model.Response{
		Message: "Get " + name + " success",
		StatusCode: Constant.HTTP_STATUS_OKE,
	})
}