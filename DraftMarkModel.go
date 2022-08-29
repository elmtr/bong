package bong

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DraftMark struct {
  ID        string `json:"id" bson:"id"`
	Value     int    `json:"value" bson:"value"`
	DateDay   string `json:"dateDay" bson:"dateDay"`
	DateMonth string `json:"dateMonth" bson:"dateMonth"`
	SubjectID string `json:"subjectID" bson:"subjectID"`
	StudentID string `json:"studentID" bson:"studentID"`
}

func GetDraftMarks(filter interface{}, optionsData interface{}) ([]DraftMark, error) {
	var draftMarks []DraftMark

	options := options.Find()
	options.SetSort(optionsData)

	cursor, err := DraftMarks.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &draftMarks)
	return draftMarks, err
}

func (draftMark *DraftMark) Insert() (interface{}, error) {
	draftMark.ID = GenID()
	res, err := DraftMarks.InsertOne(ctx, draftMark)

	return res.InsertedID, err
}

func UpdateDraftMark(filter interface{}, value int, dateDay string, dateMonth string) (DraftMark, error) {
  var draftMark DraftMark

  err := DraftMarks.FindOneAndUpdate(ctx, filter, bson.M{
    "$set": bson.M{
      "value": value,
      "dateDay": dateDay,
      "dateMonth": dateMonth,
    },
  }).Decode(&draftMark)

  draftMark.Value = value
  draftMark.DateDay = dateDay
  draftMark.DateMonth = dateMonth

  return draftMark, err
}

func DefinitivateDraftMark(filter interface{}) (Mark, error) {
  var draftMark DraftMark

  err := DraftMarks.FindOneAndDelete(ctx, filter).Decode(&draftMark)
  if err != nil {
    return Mark {}, err
  }

  mark := Mark(draftMark)
  _, err = mark.Insert()
  return mark, err
}