package bong

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
  ID string `json:"id" bson:"id"`
  Email string `json:"email" bson:"email"`
  Password string `json:"password" bson:"password"`
}

type AdminClaims struct {
  ID string `json:"id" bson:"id"`
  Email string `json:"email" bson:"email"`
  Password string `json:"password" bson:"password"`

  jwt.StandardClaims
}

func GenAdminToken(admin Admin) (string, error) {
  expirationTime := time.Now().Add(8760 * time.Hour)

  claims := &AdminClaims {
    ID: admin.ID,
    Email: admin.Email,
    StandardClaims: jwt.StandardClaims {
      ExpiresAt: expirationTime.Unix(),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(JWTKey)
} 

func GetAdmin(email string) (Admin, error) {
  var admin Admin

  err := Admins.FindOne(ctx, 
    bson.M{
      "email": email,
    },
  ).Decode(&admin)
  return admin, err
}

func (admin *Admin) Insert() (interface {}, error) {
  if admin.ID == "" {
    admin.ID = GenID()
  }

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
  if err != nil {
    return nil, err
  }

  admin.Password = string(hashedPassword)

  return Admins.InsertOne(ctx, admin)
}