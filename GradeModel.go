package bong

import "go.mongodb.org/mongo-driver/mongo/options"

type Grade struct {
	ID     string `json:"id" bson:"id"`
	SchoolID    string `json:"schoolID" bson:"schoolID"`
	YearFrom    int    `json:"yearFrom" bson:"yearFrom"`
	YearTo      int    `json:"yearTo" bson:"yearTo"`
	GradeNumber int    `json:"gradeNumber" bson:"gradeNumber"`
	GradeLetter string `json:"gradeLetter" bson:"gradeLetter"`
	Intervals   string `json:"intervals" bson:"intervals"`
}

func GetGrades(filter interface{}, optionsData interface{}) ([]Grade, error) {
	var grades []Grade

	options := options.Find()
	options.SetSort(optionsData)

	cursor, err := Grades.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &grades)
	if len(grades) == 0{
		grades = []Grade {}
	}
	return grades, err
}

func GetGrade(filter interface{}) (Grade, error) {
	var grade Grade

	err := Grades.FindOne(ctx, filter).Decode(&grade)

	return grade, err
}