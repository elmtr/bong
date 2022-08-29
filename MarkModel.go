package bong

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mark struct {
	ID        string `json:"id" bson:"id"`
	Value     int    `json:"value" bson:"value"`
	DateDay   string `json:"dateDay" bson:"dateDay"`
	DateMonth string `json:"dateMonth" bson:"dateMonth"`
	SubjectID string `json:"subjectID" bson:"subjectID"`
	StudentID string `json:"studentID" bson:"studentID"`
}

func GetMarks(filter interface{}, optionsData interface{}) ([]Mark, error) {
	var marks []Mark

	options := options.Find()
	options.SetSort(optionsData)

	cursor, err := Marks.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &marks)
	return marks, err
}

func (mark *Mark) Insert() (interface{}, error) {
	if (mark.ID == "") {
		mark.ID = GenID()
	}
	
	res, err := Marks.InsertOne(ctx, mark)

	return res.InsertedID, err
}