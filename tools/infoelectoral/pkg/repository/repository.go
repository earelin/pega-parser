package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/election"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const insertarProcesoElectoral = "INSERT INTO procesos_electorais(tipo, ambito_ine, data) VALUES (?, ?, ?)"
const insertarCandidatura = "INSERT INTO candidaturas(proceso_electoral_id, siglas, nome) VALUES (?, ?, ?)"
const insertarLista = "INSERT INTO listas(candidatura_id, ambito_ine) VALUES (? , ?)"
const insertarCandidato = "INSERT INTO candidatos(candidatura_id, orden, titular, nombre, apelidos) VALUES (?, ?, ?, ?, ?)"

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

func (r *Repository) CreateProcesoElectoral(e election.Election) (int64, error) {
	var result sql.Result
	var err error
	if e.Scope == 99 {
		result, err = r.pool.ExecContext(r.ctx, insertarProcesoElectoral, e.Type, nil, e.Date)
	} else {
		result, err = r.pool.ExecContext(r.ctx, insertarProcesoElectoral, e.Type, e.Scope, e.Date)
	}

	if err != nil {
		return 0, fmt.Errorf("no ha sido posible obtener el id de una candidatura guardada: %w", err)
	}

	var id int64
	id, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("no ha sido posible obtener el id de una candidatura guardada: %w", err)
	}

	return id, err
}

func (r *Repository) CreateCandidaturas(procesoElectoral int64, candidatures []election.Candidature) (map[int]int64, error) {
	var importedItems = make(map[int]int64)

	for _, c := range candidatures {
		var result, err = r.pool.ExecContext(r.ctx, insertarCandidatura, procesoElectoral, c.Acronym, c.Name)
		if err != nil {
			return nil, fmt.Errorf("non foi posible gardar unha candidatura: %w", err)
		}

		var id int64
		id, err = result.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("no foi posible obter o id dunha candidatura gardada: %w", err)
		}

		importedItems[c.Code] = id
	}

	return importedItems, nil
}

func (r *Repository) CrearListasECandidatos(listaCandidatos []election.Candidate, candidaturasImportadas map[int]int64) error {
	var listasImportadas = make(map[string]int)
	for _, c := range listaCandidatos {
		listCodeAndPosition := fmt.Sprintf("%d_%d", c.CandidatureCode, c.Position)
		var listaId, listaImportada = listasImportadas[listCodeAndPosition]
		if !listaImportada {
			var result, err = r.pool.ExecContext(r.ctx, insertarLista)
		}
	}
	return errors.New("not implemented yet")
}
