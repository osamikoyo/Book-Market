package data

import (
	"auth/src/talking"
	"log/slog"
	"os"
	"strings"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}
func AddUserToDB(u User) error{
	db, err := gorm.Open(sqlite.Open("storage/users.db"))
	if err != nil {
		return err
	}
	

	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
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
func WorkWithMessage(wg *sync.WaitGroup) {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ch := make(chan string)
	chfe := make(chan error)


	go talking.GetMessage("users", ch, chfe)

	
	select{
	case result := <- ch:
		User := EncodeStringToUser(result)
		err := AddUserToDB(User)
		if err != nil {
			loger.Error(err.Error())
			wg.Done()
		}
		wg.Done()
	case err := <- chfe:
		if err != nil {
			loger.Error(err.Error())
			wg.Done()
		}
	}
}
func GetUserFromToken(token string) (User, error){
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	var u User


	db, err := gorm.Open(sqlite.Open("storage/users.db"))
	if err != nil {
		loger.Error(err.Error())
		return u, err
	}

	res := db.Where("token = ?", token).First(&u)
	if res.Error != nil {
		return u, err
	}

	return u, nil
}