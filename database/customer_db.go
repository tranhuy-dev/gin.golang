package database

import (
	"gin.api/model"
	"context"
	"errors"
	"log"
	"gin.api/constant"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)
func AddCustomer(req model.Customer , sequence int) (interface{}, error) {
	customerCollection :=  Client.Database(DatabaseName).Collection("customers")
	newCustomer := model.Customer{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Address: req.Address,
		Status: Constant.STATUS_OPEN,
		Reputation: 10,
		Image: req.Image,
		Books: make([]int , 0),
		Phone: req.Phone,
		Friends: make([]int , 0),
		Password: req.Password,
		Token: "",
		CustomerID: sequence,
	}

	_, errorQueryCustomer := customerCollection.InsertOne(context.TODO() , newCustomer)
	if errorQueryCustomer != nil {
		return nil ,errors.New("Add fail")
	}

	dataResponse := map[string]interface{}{}
	dataResponse["message"] = "SUCESS"

	return dataResponse, nil
}

func GetAllCustomer() (interface{} , error) {
	customerCollection :=  Client.Database(DatabaseName).Collection("customers")
	var customers []*model.Customer
	findOptions := options.Find()
	findOptions.SetLimit(100)

	cur , errQuery := customerCollection.Find(context.TODO() , bson.D{{}} ,findOptions)
	if errQuery != nil {
		log.Fatal(errQuery)
	}

	for cur.Next(context.TODO()) {
		var customerChild model.Customer
		errDecode := cur.Decode(&customerChild)
		if errDecode != nil {
			log.Fatal(errDecode)
		}

		customers = append(customers , &customerChild)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return customers , nil
}

func GetDetailCustomer(idCustomer int) (interface{} , error) {
	var customer model.Customer
	customerCollection :=  Client.Database(DatabaseName).Collection("customers")
	errQuery := customerCollection.FindOne(context.TODO() , bson.D{
		{"customer_id" , idCustomer},
	}).Decode(&customer)

	if errQuery != nil {
		log.Fatal(errQuery)
	}

	return customer , nil
}