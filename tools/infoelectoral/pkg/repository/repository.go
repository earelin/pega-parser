package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/election"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const insertProcesoElectoral = "INSERT INTO procesos_electorais(tipo, ambito_ine, data) VALUES (?, ?, ?)"
const insertCandidatura = "INSERT INTO candidaturas()"

type Repository struct {
	pool *sql.DB
	ctx  context.Context
}

func NewRepository(c Config, ctx context.Context) (*Repository, error) {
	var r Repository

	var pool, err = sql.Open("mysql", c.toString())
	if err != nil {
		return nil, err
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	r.pool = pool
	r.ctx = ctx

	return &r, nil
}

func (r *Repository) CheckConnection() error {
	var ctx, cancel = context.WithTimeout(r.ctx, 5*time.Second)
	defer cancel()
	return r.pool.PingContext(ctx)
}

func (r *Repository) CloseConnection() error {
	return r.pool.Close()
}

func (r *Repository) CreateProcesoElectoral(e election.Election) error {
	var err error
	if e.Scope == 99 {
		_, err = r.pool.ExecContext(r.ctx, insertProcesoElectoral, e.Type, nil, e.Date)
	} else {
		_, err = r.pool.ExecContext(r.ctx, insertProcesoElectoral, e.Type, e.Scope, e.Date)
	}
	return err
}

func (r *Repository) CreateCandidaturas(candidatures []election.Candidature) (map[int]int64, error) {
	var importedItems map[int]int64

	for _, c := range candidatures {
		var result, err = r.pool.ExecContext(r.ctx, insertCandidatura, c.Acronym, c.Name)
		if err != nil {
			return nil, fmt.Errorf("no ha sido posible guardar una candidatura: %w", err)
		}

		var id int64
		id, err = result.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("no ha sido posible obtener el id de una candidatura guardada: %w", err)
		}

		importedItems[c.Code] = id
	}

	return importedItems, nil
}
