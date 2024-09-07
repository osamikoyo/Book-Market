package database

import (
	"log/slog"
	"os"
)

type User struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}


type Purshaes struct{
	Id string
	Boors string
	Sale int
	Username string
}


func GetPursh(u User) (Purshaes, error){
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	err


}