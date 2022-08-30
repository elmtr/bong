package bong

import (
	"context"
	"fmt"
	"log"

	// mongodb
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"

	// redis
	"github.com/go-redis/redis/v8"
)

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

// administrative collections
var Schools *mongo.Collection
var Grades *mongo.Collection
var Subjects *mongo.Collection
var Periods *mongo.Collection

// main collections
var Marks *mongo.Collection
var Truancies *mongo.Collection
var DraftMarks *mongo.Collection
var PointsCollection *mongo.Collection
var AverageMarks *mongo.Collection

// accounts collections
var Teachers *mongo.Collection
var Students *mongo.Collection
var Parents *mongo.Collection
var Admins *mongo.Collection

var ctx = context.Background()
var Client *mongo.Client
var RDB *redis.Client

// initializing database
func InitDB(MongoURI string) {
  var err error

  Client, err = mongo.Connect(
    context.Background(),
    options.Client().ApplyURI(MongoURI),
  )

  if err != nil {
    log.Fatal(err)
  }

  // loading collections
  Schools = GetCollection("schools", Client)
  Grades = GetCollection("grades", Client)
  Subjects = GetCollection("subjects", Client)
  Periods = GetCollection("periods", Client)
  Marks = GetCollection("marks", Client)
  Truancies = GetCollection("truancies", Client)
  DraftMarks = GetCollection("draftmarks", Client)
  AverageMarks = GetCollection("averagemarks", Client)
  PointsCollection = GetCollection("points", Client)
  Teachers = GetCollection("teachers", Client)
  Students = GetCollection("students", Client)
  Parents = GetCollection("parents", Client)
  Admins = GetCollection("admins", Client)

  fmt.Println("connected to MongoDB")
}

func GetCollection(collectionName string, client *mongo.Client) (*mongo.Collection) {
  return client.Database("dev").Collection(collectionName)
}

func InitCache(RedisOptions *redis.Options) {
  RDB = redis.NewClient(RedisOptions)
  
  pong, _ := RDB.Ping(context.Background()).Result()
  if pong == "PONG" {
    fmt.Println("connected to Redis")
  } else {
    fmt.Println("not connected to redis")
  }
}


func Set(key string, value string) error {
  err := RDB.Set(ctx, key, value, 0).Err()

  return err
}

func Get(key string) (string, error) {
  val, err := RDB.Get(ctx, key).Result()

  return val, err
}

func Del(key string) error {
  _, err := RDB.Del(ctx, key).Result()

  return err
}
