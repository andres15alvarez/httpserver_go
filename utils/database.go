package utils

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

func ConnectDB() *pg.DB {
	options, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db := pg.Connect(options)
	return db
}
