package model

type Sequence struct {
	Type string `json:"type" bson:"type"`
	Count int `json:"count" bson:"count"`
}