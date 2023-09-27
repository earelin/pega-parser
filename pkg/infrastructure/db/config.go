package db

import (
	"database/sql"
	"log"
)

type Config struct {
	Filename string
}

func (rc Config) BuildPool() *sql.DB {
	var pool, err = sql.Open("sqlite3", rc.Filename)
	if err != nil {
		log.Panicf("Error connecting to database: %s", err)
	}

	return pool
}
