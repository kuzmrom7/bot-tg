package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"

)

var (
	db *sql.DB
)

func New() {
	var (
		DB   = "bot"
		User = "bot"
		Pass = "bot"
		Host = "localhost"
		Port = 5432
		SSL  = "disable"
	)
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		Host, Port, User, Pass, DB, SSL))
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	log.Printf("DB connection on %Port", Port)
}
