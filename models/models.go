package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// Gorm Postgres library
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		log.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Print(err)
	}

	db = conn
	conn.Debug().AutoMigrate(&Admin{},
		&Client{},
		&Coins{},
		&Deposit{},
		&Event{},
		&Racing{},
		&Horse{},
		&Remates{},
		&Tablas{},
		&Retiro{})
}

// GetDB .
func GetDB() *gorm.DB {
	return db
}
