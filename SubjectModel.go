package bong

import "go.mongodb.org/mongo-driver/mongo/options"

type Subject struct {
	ID string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	Grade     Grade  `json:"grade" bson:"grade"`
}

func GetSubject(filter interface{}) (Subject, error) {
	var subject Subject

	err := Subjects.FindOne(ctx, filter).Decode(&subject)

	return subject, err
}

func GetSubjects(filter interface{}, optionsData interface{}) ([]Subject, error) {
	var subjects []Subject

	options := options.Find()
	options.SetSort(optionsData)

	cursor, err := Subjects.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &subjects)
	return subjects, err
}
