package bong

import (
	"context"
	"log"

	// mongodb
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"

	// redis
	"github.com/go-redis/redis/v8"
)

// collections
var AverageMarks *mongo.Collection
var FinalMarks *mongo.Collection
var Grades *mongo.Collection
var Marks *mongo.Collection
var Parents *mongo.Collection
var Periods *mongo.Collection
var Schools *mongo.Collection
var Students *mongo.Collection
var Subjects *mongo.Collection
var Teachers *mongo.Collection
var TermMarks *mongo.Collection
var Truancies *mongo.Collection

// sort types
var DateSort interface{} = bson.D{
  {Key: "dateMonth", Value: 1}, 
  {Key: "dateDay", Value: 1},
}
var TermSort interface{} = bson.D{}
var EmptySort interface{} = bson.D{
  {Key: "term", Value: 1},
}
var PeriodSort interface{} = bson.D{
  {Key: "day", Value: 1}, 
  {Key: "interval", Value: 1},
}
var GradeSort interface{} = bson.D{
  {Key: "grade.gradeNumber", Value: 1}, 
  {Key: "grade.gradeLetter", Value: 1},
}
var LastNameSort interface{} = bson.D{
  {Key: "lastName", Value: 1},
}

var ctx = context.Background()
var RDB *redis.Client

// initializing database
func InitDatabase(MongoURI string, RedisOptions *redis.Options) {
  // client
  client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoURI))
  if err != nil {
    log.Fatal(err)
  }

  // redis
  RDB = redis.NewClient(RedisOptions)

  // setting up collections
  AverageMarks = getCollection("averagemarks", client)
  FinalMarks = getCollection("finalmarks", client)
  Grades = getCollection("grades", client)
  Marks = getCollection("marks", client)
  Parents = getCollection("parents", client)
  Periods = getCollection("periods", client)
  Schools = getCollection("schools", client)
  Students = getCollection("students", client)
  Subjects = getCollection("subjects", client)
  Teachers = getCollection("teachers", client)
  TermMarks = getCollection("termmarks", client)
  Truancies = getCollection("truancies", client)
}

func getCollection(collectionName string, client *mongo.Client) (*mongo.Collection) {
  return client.Database("elmtree").Collection(collectionName)
}

func Set(key string, value string) error {
  err := RDB.Set(ctx, key, value, 0).Err()

  return err
}

func Get(key string) (string, error) {
  val, err := RDB.Get(ctx, key).Result()

  return val, err
}
