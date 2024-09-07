package database

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type User struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}
func EncodeStringToUser(data string) User {
	var res User

	result := strings.Split(data, "|")
	res.Username = result[0]
	res.Email = result[1]
	res.Password = result[2]
	res.Token = result[3]

	return res
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
func CreateMessage(u User) string {
	result := fmt.Sprintf("%s|%s|%s|%s", u.Username, u.Email, u.Password, u.Token)
	return result
}
func produceMessage(topic string, message string) {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()


	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key"), 
			Value: []byte(message),
		},
	)
	if err != nil {
		loger.Error( err.Error())
	}
	loger.Info("Сообщение отправлено:", message, nil)
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
	produceMessage("users", CreateMessage(u))

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
