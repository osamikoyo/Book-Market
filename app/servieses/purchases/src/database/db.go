package database

import (
	"log/slog"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}


type Purshaes struct{
	Id string `json:"id"`
	Books string `json:"books"`
	Sale int	`json:"sale"`
	Username string `json:"username"`
}


func GetPursh(u User) ([]Purshaes, error){
	var Pershaes []Purshaes

	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	db, err := gorm.Open(sqlite.Open("purshes.db"))
	if err != nil {
		loger.Error(err.Error())
		return Pershaes, err
	}
	res := db.Where("username = ?", u.Username).Find(&Pershaes)
	if res.Error != nil {
		loger.Error(res.Error.Error())
		return Pershaes, res.Error
	}

	return Pershaes, nil
}