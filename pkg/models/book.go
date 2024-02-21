package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"github.com/simple_MYSQL_CRUD_API_With_Golang/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	BookName    string `gorm:"NOT NULL" json:"book_name"`
	Author      string `gorm:"NOT NULL" json:"author"`
	Publication string `gorm:"NOT NULL" json:"publication"`
}

func init() {
	db = config.GetConnect()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBookDB() {
	db.Create(b)
}

func DeleteBookById(id uint) bool {
	if !checkIdExisted(int64(id)) {
		return false
	}
	db.Where("id = ?", id).Delete(&Book{})
	return true
}

func GetAllBookDB() []Book {
	var allbook []Book
	db.Find(&allbook)
	return allbook
}

func GetBookByID(id uint) (Book, bool) {
	if !checkIdExisted(int64(id)) {
		return Book{}, false
	}
	var book Book
	db.Where("id = ?", id).Find(&book)
	return book, true
}

func UpdateById(id uint, book Book) bool {
	if !checkIdExisted(int64(id)) {
		return false
	}
	var curBook Book
	db.Where("id = ?", id).First(&curBook)
	if len(book.BookName) > 0 {
		curBook.BookName = book.BookName
	}
	if len(book.Author) > 0 {
		curBook.Author = book.Author
	}
	if len(book.Publication) > 0 {
		curBook.Publication = book.Publication
	}
	db.Save(&curBook)
	return true
}

func checkIdExisted(id int64) bool {
	var count int64
	db.Model(&Book{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		return false
	}
	return true
}
