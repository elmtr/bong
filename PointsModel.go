package bong

import "go.mongodb.org/mongo-driver/bson"

type Points struct {
  ID        string `json:"id" bson:"id"`
	Value     int    `json:"value" bson:"value"`
	SubjectID string `json:"subjectID" bson:"subjectID"`
	StudentID string `json:"studentID" bson:"studentID"`
}

func ModifyPoints(filter interface{}, amount int) (Points, error) {
  var points Points

  err := PointsCollection.FindOneAndUpdate(ctx, filter, bson.M{
    "$inc": bson.M{
      "value": amount,
    },
  }).Decode(&points)

  points.Value += amount
  return points, err
}

func IncreasePoints(filter interface{}, ) (Points, error) {
  return ModifyPoints(filter, 1)
}

func DecreasePoints(filter interface{}, ) (Points, error) {
  return ModifyPoints(filter, -1)
}