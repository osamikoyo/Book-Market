package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct{
	Name string
	Author string
	Data string
}

func GetBooks() ([]Book, error){
	var Books []Book
	db, err := gorm.Open(sqlite.Open("storage/book.db"))
	if err != nil {
		return Books, err
	}
} 