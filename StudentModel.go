package bong

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
  ID string `json:"id"`
  FirstName string `json:"firstName" bson:"firstName"`
  LastName string `json:"lastName" bson:"lastName"`
  Phone string `json:"phone" bson:"phone"`
  Grade Grade `json:"grade"`
  Subjects []ShortSubject `json:"subjects"`
  Password string `json:"password" bson:"password"`
  Passcode string `json:"passcode" bson:"passcode"`
}

type ShortSubject struct {
  ID string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type StudentClaims struct {
  ID string `json:"id"`
  FirstName string `json:"firstName" bson:"firstName"`
  LastName string `json:"lastName" bson:"lastName"`
  Phone string `json:"phone" bson:"phone"`
  Grade Grade `json:"grade"`
  Subjects []ShortSubject `json:"subjects"`
  
  jwt.StandardClaims
}

func GenStudentToken(student Student) (tokenString string,err error) {
  expirationTime := time.Now().Add(8760 * time.Hour)

  claims := &StudentClaims {
    ID: student.ID,
    FirstName: student.FirstName,
    LastName: student.LastName,
    Phone: student.Phone,
    Grade: student.Grade,
    Subjects: student.Subjects,
    StandardClaims: jwt.StandardClaims {
      ExpiresAt: expirationTime.Unix(),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(JWTKey)
}

func GetStudent(filter interface{}) (Student, error) {
  var student Student

  err := Students.FindOne(ctx, filter).Decode(&student)
  return student, err
}

func GetStudents(filter interface{}) ([]Student, error) {
  var students []Student

  options := options.Find()
  options.SetSort(EmptySort)

  cursor, err := Students.Find(ctx, filter, options)
  if err != nil {
    return nil, err
  }
  err = cursor.All(ctx, &students)
  return students, err
}

func AddStudentSubject(id string, subjects []ShortSubject, subject ShortSubject) ([]ShortSubject, error) {
  subjects = append(subjects, subject)

  _, err := Students.UpdateOne(ctx, bson.M{
    "id": id,
  }, bson.M{
    "$set": bson.M{
      "subjects": subjects,
    },
  })

  return subjects, err
}

func UpdateStudentGrade(filter interface{}, grade Grade) (Student, error) {
  var student Student 

  err := Students.FindOneAndUpdate(ctx, filter, bson.M{
    "$set": bson.M{
      "grade": grade,
    },
  }).Decode(&student)

  student.Grade = grade

  return student, err
}

func (student *Student) Insert() (error) {
  student.ID = GenID()
  student.Grade = Grade {}
  student.Subjects = []ShortSubject {}

  _, err := Students.InsertOne(ctx, student)
  return err
}

func UpdateStudent(id interface {}, update interface {}) (error) {
  _, err := Students.UpdateOne(
    ctx,
    bson.M{"id": id},
    bson.M{
      "$set": update,
    },
  )

  return err
}
