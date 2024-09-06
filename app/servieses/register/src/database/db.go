package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type User struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func AddUserToDB(u User) error{
	db, err := gorm.Open(sqlite.Open("storage/users.db"))
	if err != nil {
		return err
	}
	
}