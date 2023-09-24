package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Config struct {
	Host     string
	Database string
	User     string
	Password string
}

func (rc Config) toString() string {
	var userPassword string
	host := rc.Host

	if rc.User != "" {
		userPassword = rc.User
		host = "@" + host
		if rc.Password != "" {
			userPassword = userPassword + ":" + rc.Password
		}
	}

	return fmt.Sprintf("%s%s/%s", userPassword, host, rc.Database)
}

func (rc Config) BuildPool() *sql.DB {
	var pool, err = sql.Open("mysql", rc.toString())
	if err != nil {
		log.Panicf("Error connecting to database: %s", err)
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	return pool
}
