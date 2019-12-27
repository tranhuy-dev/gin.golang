package database

import (
	"gin.api/model"
	"context"
	"errors"
)

func AddCustomer(req model.Customer) (interface{}, error) {
	newCustomer := model.Customer{
		Name: req.Name,
	}

	customerCollection :=  Client.Database(DatabaseName).Collection("customers")
	_, errorQueryCustomer := customerCollection.InsertOne(context.TODO() , newCustomer)
	if errorQueryCustomer != nil {
		return nil ,errors.New("Add fail")
	}

	dataResponse := map[string]interface{}{}
	dataResponse["message"] = "SUCESS"

	return dataResponse, nil
}