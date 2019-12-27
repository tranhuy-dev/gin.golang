package database

import (
	"gin.api/model"
	"context"
	"errors"
	"gin.api/constant"
)

func AddCustomer(req model.Customer , sequence int) (interface{}, error) {
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

	customerCollection :=  Client.Database(DatabaseName).Collection("customers")
	_, errorQueryCustomer := customerCollection.InsertOne(context.TODO() , newCustomer)
	if errorQueryCustomer != nil {
		return nil ,errors.New("Add fail")
	}

	dataResponse := map[string]interface{}{}
	dataResponse["message"] = "SUCESS"

	return dataResponse, nil
}