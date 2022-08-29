package bong

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Truancy struct {
	ID        string `json:"id" bson:"id"`
	Motivated bool   `json:"motivated" bson:"motivated"`
	DateDay   string `json:"dateDay" bson:"dateDay"`
	DateMonth string `json:"dateMonth" bson:"dateMonth"`
	SubjectID string `json:"subjectID" bson:"subjectID"`
	StudentID string `json:"studentID" bson:"studentID"`
}

func GetTruancies(filter interface{}, optionsData interface{}) ([]Truancy, error) {
	var truancies []Truancy

	options := options.Find()
	options.SetSort(optionsData)

	cursor, err := Truancies.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &truancies)
	return truancies, err
}

func (truancy *Truancy) Insert() (interface{}, error) {
	truancy.ID = GenID()
	truancy.Motivated = false
	res, err := Truancies.InsertOne(ctx, truancy)

	return res.InsertedID, err
}

func MotivateTruancy(filter interface{}) (Truancy, error) {
	var truancy Truancy

	err := Truancies.FindOneAndUpdate(ctx, filter, bson.M{
		"$set": bson.M{"motivated": true},
	}).Decode(&truancy)

	truancy.Motivated = true

	return truancy, err
}