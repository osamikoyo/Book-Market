package database

import (
	"log/slog"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct{
	Name string `json:"name"`
	Author string `json:"author"`
	Data string `json:"data"`
}

func GetBooksFromDB() ([]Book, error){
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	var books []Book
	db, err := gorm.Open(sqlite.Open("storage/book.db"))
	if err != nil {
		return books, err
	}
	result := db.Find(&books)
	if result.Error != nil {
		loger.Error(result.Error.Error())
		return books, result.Error
	}
	return books, nil
} 