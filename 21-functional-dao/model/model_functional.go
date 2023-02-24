package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Option func(*gorm.DB) *gorm.DB

func withCategory(category int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("category = ?", category)
	}
}
func withPrice(price int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price = ?", price)
	}
}
func withPriceRange(min, max int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price >= ? AND price <= ?", min, max)
	}
}
func TableName(tableName string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(tableName)
	}
}

func GetBookByOptions(options ...Option) ([]Book, error) {
	var books []Book
	db := GetDB()
	for _, option := range options {
		db = option(db)
	}
	result := db.Find(&books)
	return books, result.Error
}

// client code
func ClientExample2() {
	books, _ := GetBookByOptions(
		withPrice(200),
		withCategory(2),
		TableName("books"),
		//..
	)
	fmt.Println(books)
}
