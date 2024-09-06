package database

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/golang-jwt/jwt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type User struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}
func GenerateKeyPair(bits int) (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", err
	}

	
	privPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	return string(privPEM), nil
}
func CreateToken(user User) (string, error) {
	jwtSecret, err := GenerateKeyPair(2048)
	if err != nil {
		return "", err
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
func AddUserToDB(u User) error{
	db, err := gorm.Open(sqlite.Open("storage/users.db"))
	if err != nil {
		return err
	}
	

	result := db.Create(&u)

	var errs error
	u.Token, errs = CreateToken(u)
	if errs != nil {
		return errs
	}


	if result.Error != nil {
		return result.Error
	}
	return nil
}
func GetAllUsers() ([]User, error) {
	var Users []User
	db, err := gorm.Open(sqlite.Open("storage/users.db"))
	if err != nil {
		return Users, err
	}

	db.Find(&Users)
	return Users, nil
}
