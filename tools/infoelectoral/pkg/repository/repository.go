package repository

import (
	"database/sql"
	"github.com/earelin/pega/tools/infoelectoral/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(c Config) (Repository, error) {
	var r Repository
	var err error
	r.db, err = sql.Open("mysql", c.toString())
	return r, err
}

func (r *Repository) CheckConnection(conf config.Config) {

}

func (r *Repository) CloseConnection() error {
	return r.db.Close()
}
