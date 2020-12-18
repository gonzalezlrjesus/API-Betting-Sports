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

	// DB config reading
	dbURI := os.Getenv("DATABASE_URL")
	if dbURI == "" {
		dbURI = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("db_host"), os.Getenv("db_user"), os.Getenv("db_name"), os.Getenv("db_pass"))
	}
	log.Println(dbURI)

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
