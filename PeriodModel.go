package bong

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Period struct {
  ID        string `json:"id" bson:"id"`

  Day       int    `json:"day" bson:"day"`
  Interval  int    `json:"interval" bson:"interval"`

  // modifiable
  Room      string `json:"room" bson:"room"`
  // indexable
  Subject Subject   `json:"subject" bson:"subject"`
}

func GetPeriods(filter interface {}, optionsData interface{}) ([]Period, error) {
  var periods []Period

	options := options.Find()
	options.SetSort(optionsData)

	cursor, err := Periods.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &periods)
	if len(periods) == 0{
		periods = []Period {}
	}
	return periods, err
}

func GetPeriod(filter interface {}) (Period, error) {
  var period Period

  err := Periods.FindOne(
    ctx,
    filter,
  ).Decode(&period)

  return period, err
}

func UpdatePeriod(id interface{}, subject Subject, room string) (Period, error) {
  var period Period
  err := Periods.FindOneAndUpdate(
    ctx,
    bson.M {
      "id": id,
    },
    bson.M{
      "$set": bson.M{
        "room": room,
        "subject": subject,
      },
    },
  ).Decode(&period)

  period.Room = room
  period.Subject = subject

  return period, err
}


func (period *Period) Insert() (interface{}, error) {
	if (period.ID == "") {
		period.ID = GenID()
	}
	
	res, err := Periods.InsertOne(ctx, period)

	return res.InsertedID, err
}

func DeletePeriod(filter interface{}) (error) {
  _, err := Periods.DeleteOne(
    ctx, 
    filter,
  )

  return err
}