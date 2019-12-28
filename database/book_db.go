package database

import (
	"gin.api/model"
	"gin.api/constant"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func CreateBook(req model.Book , idBook int) (interface{} , error) {
	bookCollection := Client.Database(DatabaseName).Collection("books")
	newBooks := model.Book{
		Name: req.Name,
		Type: req.Type,
		BookID: idBook,
		Status: Constant.STATUS_OPEN,
		Quality: 10,
		TotalBorrow: 0,
		Description: req.Description,
		Author: req.Author,
		Image: req.Image,
	}

	_ , errAdd := bookCollection.InsertOne(context.TODO() , newBooks)
	if errAdd != nil {
		return nil , errAdd
	}

	return "Success" , nil
}

func GetAllBook() (interface{} , error) {
	bookCollection := Client.Database(DatabaseName).Collection("books")
	var books []*model.Book
	findOption := options.Find()
	findOption.SetLimit(100)

	cur, errQuery := bookCollection.Find(context.TODO() , bson.D{{}} , findOption)
	if errQuery != nil {
		log.Fatal(errQuery)
	}

	for cur.Next(context.TODO()) {
		var bookChild model.Book
		err := cur.Decode(&bookChild)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books , &bookChild)
	}
	
	return books , nil
}

func GetDetailBook(id int) (interface{} , error) {
	bookCollection := Client.Database(DatabaseName).Collection("books")
	var book model.Book
	filter := bson.D{
		{"book_id" , id}}

	err := bookCollection.FindOne(context.TODO() , filter).Decode(&book)

	if err != nil {
		log.Fatal(err)
	}

	return book, nil
}