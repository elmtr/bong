package bong

import "go.mongodb.org/mongo-driver/bson"

type School struct {
	ID        string   `json:"id" bson:"id"`
	Name      string   `json:"name" bson:"name"`
	Intervals []string `json:"intervals" bson:"intervals"`
}

func GetSchools() ([]School, error) {
	var schools []School

	cursor, err := Schools.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &schools)
	return schools, err
}

func GetSchool(filter interface{}) (School, error) {
	var school School

	err := Schools.FindOne(ctx, filter).Decode(&school)

	return school, err
}