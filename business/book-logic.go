package business

import (
	"github.com/gin-gonic/gin"
	"gin.api/model"
	"gin.api/constant"
	"gin.api/database"
	"strconv"
)

func CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(Constant.HTTP_STATUS_BAD_REQUEST , model.Response{
			Message: "Bad Request",
			StatusCode: Constant.HTTP_STATUS_BAD_REQUEST,
		})
	}
	seqCount , errSeq := database.GetSequence("book")
	if errSeq != nil {
		c.JSON(Constant.HTTP_STATUS_INTERNAL_ERROR , model.Response{
			Message: "get sequence fail",
			StatusCode: Constant.HTTP_STATUS_INTERNAL_ERROR,
		})
	}

	_ , errAdd := database.CreateBook(book , seqCount)
	if errAdd != nil {
		c.JSON(Constant.HTTP_STATUS_INTERNAL_ERROR , model.Response{
			Message: "Add book fail",
			StatusCode: Constant.HTTP_STATUS_INTERNAL_ERROR,
		})
	}
	c.JSON(Constant.HTTP_STATUS_CREATED , model.Response{
		Message: "Add book " + book.Name + " success",
		StatusCode: Constant.HTTP_STATUS_CREATED,
	})
}


func GetAllBook(c *gin.Context) {
	dataBooks , errQuery := database.GetAllBook()
	if errQuery != nil {
		c.JSON(Constant.HTTP_STATUS_INTERNAL_ERROR , model.Response{
			Message: "Get Books Fail",
			StatusCode: Constant.HTTP_STATUS_INTERNAL_ERROR,
		})
	}
	c.JSON(Constant.HTTP_STATUS_OKE , model.Response{
		Message: "Get Books Success",
		StatusCode: Constant.HTTP_STATUS_OKE,
		Data: dataBooks,
	})
}

func GetDetailBook(c *gin.Context) {
	id := c.Param("id")
	idBooks , errBook := strconv.Atoi(id)

	if errBook != nil {
		c.JSON(Constant.HTTP_STATUS_BAD_REQUEST , model.Response{
			Message: "Bad request",
			StatusCode: Constant.HTTP_STATUS_BAD_REQUEST,
		})
	}

	dataBook , errQuery := database.GetDetailBook(idBooks)
	if errQuery != nil {
		c.JSON(Constant.HTTP_STATUS_INTERNAL_ERROR , model.Response{
			Message: "get book fail",
			StatusCode: Constant.HTTP_STATUS_INTERNAL_ERROR,
		})
	}

	c.JSON(Constant.HTTP_STATUS_OKE , model.Response{
		Message: "Get Success",
		StatusCode: Constant.HTTP_STATUS_OKE,
		Data: dataBook,
	})
}