package bong

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

type Teacher struct {
  ID string `json:"id" bson:"id"`
  LastName string `json:"lastName" bson:"lastName"`
  FirstName string `json:"firstName" bson:"firstName"`
  Phone string `json:"phone" bson:"phone"`
  Homeroom Grade `json:"homeroom" bson:"homeroom"`
  Subjects []Subject `json:"subjects" bson:"subjects"`
  Password string `json:"password" bson:"password"`
  Passcode string `json:"passcode" bson:"passcode"`
}

type TeacherClaims struct {
  ID string `json:"id" bson:"id"`
  LastName string `json:"lastName" bson:"lastName"`
  FirstName string `json:"firstName" bson:"firstName"`
  Phone string `json:"phone" bson:"phone"`
  Homeroom Grade `json:"homeroom" bson:"homeroom"`
  Subjects []Subject `json:"subjects" bson:"subjects"`
  
  jwt.StandardClaims
}

func GenTeacherToken(teacher Teacher) (string, error) {
  expirationTime := time.Now().Add(8760 * time.Hour)

  claims := &TeacherClaims {
    ID: teacher.ID,
    FirstName: teacher.FirstName,
    LastName: teacher.LastName,
    Phone: teacher.Phone,
    Homeroom: teacher.Homeroom,
    Subjects: teacher.Subjects,
    StandardClaims: jwt.StandardClaims {
      ExpiresAt: expirationTime.Unix(),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(JWTKey)
}


func GetTeacher(filter interface{}) (Teacher, error) {
  var teacher Teacher

  err := Teachers.FindOne(ctx, filter).Decode(&teacher)
  return teacher, err
}

func AddTeacherSubject(id string, subjects []Subject, subject Subject) ([]Subject, error) {
  subjects = append(subjects, subject)

  _, err := Teachers.UpdateOne(ctx, bson.M{
    "id": id,
  }, bson.M{
    "$set": bson.M{
      "subjects": subjects,
    },
  })

  return subjects, err
}

func RemoveTeacherSubject(id string, subjects []Subject, oldSubject Subject) ([]Subject, error) {
  var newSubjects []Subject 
  for _, subject := range subjects {
    if (subject.ID != oldSubject.ID) {
      newSubjects = append(newSubjects, subject)
    }
  }

  _, err := Teachers.UpdateOne(ctx, bson.M{
    "id": id,
  }, bson.M{
    "$set": bson.M{
      "subjects": newSubjects,
    },
  })

  return newSubjects, err
}

func UpdateTeacherHomeroom(id string, homeroom Grade) (Teacher, error) {
  var teacher Teacher 

  err := Teachers.FindOneAndUpdate(
    ctx, 
    bson.M{
      "id": id,
    },
    bson.M{
      "$set": bson.M{
        "homeroom": homeroom,
      },
    },
  ).Decode(&teacher)

  teacher.Homeroom = homeroom

  return teacher, err
}

func (teacher *Teacher) Insert() (error) {
  teacher.ID = GenID()
  teacher.Homeroom = Grade {}
  teacher.Subjects = []Subject {}
  
  _, err := Teachers.InsertOne(ctx, teacher)
  return err
}

func UpdateTeacher(id interface {}, update interface {}) (error) {
  _, err := Teachers.UpdateOne(
    ctx,
    bson.M{"id": id},
    bson.M{
      "$set": update,
    },
  )

  return err
}
