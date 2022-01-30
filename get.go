package bong

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMarks(subjectID string, studentID string) ([]Mark, error) {
	marksID := fmt.Sprintf("marks:%v:%v", subjectID, studentID)

  var marks []Mark

	marksJSON, err := Get(marksID)
  if err != nil {
    return nil, err
  }

  // if cache is there, unmarshal and return
  if marksJSON != "" {
    err := json.Unmarshal([]byte(marksJSON), &marks)
    if err != nil {
      return nil, err
    }
    return marks, nil
  } else {
    // but if it's not there, query the db
    options := options.Find()
    options.SetSort(DateSort)

    cursor, err := Marks.Find(ctx, bson.M{
      "subject.subjectID": subjectID,
      "studentID": studentID,
    }, options)
    if err != nil {
      return nil, err
    }
    err = cursor.All(ctx, &marks)
    if err != nil {
      return nil, err
    }


    // and update the cache
    marksJSON, err := json.Marshal(marks)
    if err != nil {
      return nil, err
    }
    err = Set(marksID, string(marksJSON))
    if err != nil {
      return nil, err
    }
    return marks, nil
  }
}