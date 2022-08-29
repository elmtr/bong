package bong

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

type Parent struct {
  ID string `json:"id" bson:"id"`
  FirstName string `json:"firstName" bson:"firstName"`
  LastName string `json:"lastName" bson:"lastName"`
  Phone string `json:"phone" bson:"phone"`
  Students []ParentStudent `json:"students" bson:"students"`
  Password string `json:"password" bson:"password"`
  Passcode string `json:"passcode" bson:"passcode"`
}

type ParentStudent struct {
  ID string `json:"id" bson:"id"`
  FirstName string `json:"firstName" bson:"firstName"`
  LastName string `json:"lastName" bson:"lastName"`
}

type ParentClaims struct {
  ID string `json:"id"`
  FirstName string `json:"firstName" bson:"firstName"`
  LastName string `json:"lastName" bson:"lastName"`
  Phone string `json:"phone" bson:"phone"`
  Students []ParentStudent `json:"students"`

  jwt.StandardClaims
}

func GenParentToken(parent Parent) (string, error) {
  expirationTime := time.Now().Add(8760 * time.Hour)

  claims := &ParentClaims {
    ID: parent.ID,
    FirstName: parent.FirstName,
    LastName: parent.LastName,
    Phone: parent.Phone,
    Students: parent.Students,
    StandardClaims: jwt.StandardClaims {
      ExpiresAt: expirationTime.Unix(),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(JWTKey)
}

func GetParent(filter interface{}) (Parent, error) {
  var parent Parent 

  err := Parents.FindOne(ctx, filter).Decode(&parent)
  return parent, err
}

func AddParentStudent(id string, students []ParentStudent, student ParentStudent) ([]ParentStudent, error) {
  students = append(students, student)

  _, err := Parents.UpdateOne(ctx, bson.M{
    "id": id,
  }, bson.M{
    "$set": bson.M{
      "students": students,
    },
  })

  return students, err
}

func (parent *Parent) Insert() (error) {
  parent.ID = GenID()

  _, err := Parents.InsertOne(ctx, parent)
  return err
}

func UpdateParent(id interface {}, update interface {}) (error) {
  _, err := Parents.UpdateOne(
    ctx,
    bson.M{"id": id},
    bson.M{
      "$set": update,
    },
  )

  return err
}
