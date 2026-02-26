package database

import (
	"GoGramTest/internal/model"
	"context"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, context.Context) {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		log.Fatal(err)
	}
	return db, ctx
}
