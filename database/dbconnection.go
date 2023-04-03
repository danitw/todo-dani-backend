package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connection() *gorm.DB {
	connection := "host=0.0.0.0 port=5432 user=postgres dbname=todo-dani password=docker sslmode=disable" // TODO: make .env

	db, err := gorm.Open("postgres", connection)

	if err != nil {
		fmt.Println("Erro: ")
		fmt.Println(err)
	}

	return db
}
