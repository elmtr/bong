package bong

import "go.mongodb.org/mongo-driver/mongo/options"

type AverageMark struct {
	ID        string `json:"id,omitempty" bson:"id,omitempty"`
	Value     int    `json:"value,omitempty" bson:"value,omitempty"`
	SubjectID string `json:"subjectID,omitempty" bson:"subjectID,omitempty"`
	StudentID string `json:"studentID,omitempty" bson:"studentID,omitempty"`
}

func GetAverageMarks(filter interface{}, optionsData interface{}) ([]AverageMark, error) {
	var averageMarks []AverageMark

	options := options.Find()
	options.SetSort(optionsData)

	cursor, err := AverageMarks.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &averageMarks)
	return averageMarks, err
}

func (averageMark *AverageMark) Insert() (interface{}, error) {
	averageMark.ID = GenID()

	res, err := AverageMarks.InsertOne(ctx, averageMark)
	
	return res.InsertedID, err
}