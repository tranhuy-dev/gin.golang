package database

import (
	"gin.api/model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSequence(typeSequence string) (int , error) {
	sequenceCollection := Client.Database(DatabaseName).Collection("sequence")
	var sequence model.Sequence
	filter := bson.D{
		{"type" , typeSequence}}

	err := sequenceCollection.FindOne(context.TODO() , filter).Decode(&sequence)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	updateSeq , _ := UpdateSequence(typeSequence , 1)
	if !updateSeq {
		return 0, errors.New("Update fail")
	}
	return sequence.Count , nil
}

func UpdateSequence(typeSequence string , value int) (bool , error) {
	sequenceCollection := Client.Database(DatabaseName).Collection("sequence")
	var sequence model.Sequence
	filter := bson.D{
		{"type" , typeSequence}}

	updateBody := bson.D{
		{"$inc" , bson.D{
			{"count",1},
		}},
	}

	err := sequenceCollection.FindOneAndUpdate(context.TODO() , filter , updateBody).Decode(&sequence)

	if err != nil {
		return false , errors.New("Update fail")
	}

	return true , nil
}