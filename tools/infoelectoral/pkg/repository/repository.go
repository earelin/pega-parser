package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/election"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const inserirProcesoElectoral = "INSERT INTO procesos_electorais(tipo, ambito, data) VALUES (?, ?, ?)"
const inserirCandidatura = "INSERT INTO candidaturas(proceso_electoral_id, siglas, nome) VALUES (?, ?, ?)"
const inserirLista = "INSERT INTO listas(candidatura_id, ambito) VALUES (? , ?)"
const inserirCandidato = "INSERT INTO candidatos(lista_id, posicion, titular, nombre, apelidos) VALUES (?, ?, ?, ?, ?)"
const inserirMesaElectoral = "INSERT INTO mesas_electorais(proceso_electoral_id, concello_id, distrito, seccion, codigo, censo, votos_blanco, votos_nulos, votos_candidaturas) VALUES (?, ?, ?, ?, ?, ?, ?, ? ,?)"
const inserirCircunscripcionCera = "INSERT INTO circunscripcions_cera(proceso_electoral_id, provincia_id, censo, votos_blanco, votos_nulos, votos_candidaturas) VALUES (?, ?, ?, ?, ?, ?)"
const cargarReferenciasConcellos = "SELECT id, provincia_id, concello_ine FROM concellos"

type Repository struct {
	pool                 *sql.DB
	ctx                  context.Context
	referenciasConcellos map[string]int
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
		result, err = r.pool.ExecContext(r.ctx, inserirProcesoElectoral, e.Type, nil, e.Date)
	} else {
		result, err = r.pool.ExecContext(r.ctx, inserirProcesoElectoral, e.Type, e.Scope, e.Date)
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

func (r *Repository) CreateCandidaturas(procesoElectoral int64, candidatures []election.Candidatura) (map[int]int64, error) {
	var importedItems = make(map[int]int64)

	for _, c := range candidatures {
		var result, err = r.pool.ExecContext(r.ctx, inserirCandidatura, procesoElectoral, c.Siglas, c.Nome)
		if err != nil {
			return nil, fmt.Errorf("non foi posible gardar unha candidatura: %w", err)
		}

		var id int64
		id, err = result.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("no foi posible obter o id dunha candidatura gardada: %w", err)
		}

		importedItems[c.Codigo] = id
	}

	return importedItems, nil
}

func (r *Repository) CrearListasECandidatos(listaCandidatos []election.Candidate, candidaturasImportadas map[int]int64) error {
	var listasImportadas = make(map[string]int64)

	for _, c := range listaCandidatos {
		var candidaturaId = candidaturasImportadas[c.CodigoCandidatura]
		codigoCandidaturaEAmbitoTerritorial := fmt.Sprintf("%d_%d", c.CodigoCandidatura, c.AmbitoTerritorial)
		var listaId, importada = listasImportadas[codigoCandidaturaEAmbitoTerritorial]
		if !importada {
			var result, err = r.pool.ExecContext(r.ctx, inserirLista, candidaturaId, c.AmbitoTerritorial)
			if err != nil {
				return fmt.Errorf("no foi posible gardar unha lista: %w", err)
			}
			listaId, err = result.LastInsertId()
			if err != nil {
				return fmt.Errorf("non foi posible obter o id dunha lista gardada: %w", err)
			}
			listasImportadas[codigoCandidaturaEAmbitoTerritorial] = listaId
		}

		var _, err = r.pool.ExecContext(r.ctx, inserirCandidato, listaId, c.Posicion, c.Titular, c.Nome, c.Apelidos)
		if err != nil {
			return fmt.Errorf("non foi posible gardar un candidato: %w", err)
		}
	}

	return nil
}

func (r *Repository) CrearMesasElectorais(procesoElectoral int64, mesas []election.MesaElectoral) error {
	referenciasConcellos, err := r.cargarReferenciasConcellos()
	if err != nil {
		return fmt.Errorf("non se puideron gardar as mesas electorais: %w", err)
	}

	for _, m := range mesas {
		if m.CodigoProvincia == 99 {
			continue
		} else if m.CodigoConcello == 999 {
			_, err = r.pool.ExecContext(r.ctx, inserirCircunscripcionCera, procesoElectoral, m.CodigoProvincia,
				m.CensoIne, m.VotosBlanco, m.VotosNulos, m.VotosCandidaturas)
		} else {
			concelloId := referenciasConcellos[fmt.Sprintf("%d_%d", m.CodigoProvincia, m.CodigoConcello)]
			_, err = r.pool.ExecContext(r.ctx, inserirMesaElectoral, procesoElectoral, concelloId, m.Distrito,
				m.Seccion, m.CodigoMesa, m.CensoIne, m.VotosBlanco, m.VotosNulos, m.VotosCandidaturas)
		}

		if err != nil {
			return fmt.Errorf("non se puido gardar a mesa ou circunscripcion CERA %+v: %w", m, err)
		}
	}

	return nil
}

func (r *Repository) cargarReferenciasConcellos() (map[string]int, error) {
	if r.referenciasConcellos != nil {
		return r.referenciasConcellos, nil
	}

	rows, err := r.pool.QueryContext(r.ctx, cargarReferenciasConcellos)
	if err != nil {
		return nil, fmt.Errorf("non se puido obter o listado de concellos: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Panic("Non se puido pechar un cursor", err)
		}
	}(rows)

	var referenciasConcellos = make(map[string]int)

	for rows.Next() {
		var id, provinciaId, concelloIne int
		err = rows.Scan(&id, &provinciaId, &concelloIne)
		if err != nil {
			return nil, fmt.Errorf("non se puido ler o listado de concellos: %w", err)
		}
		referenciasConcellos[fmt.Sprintf("%d_%d", provinciaId, concelloIne)] = id
	}

	r.referenciasConcellos = referenciasConcellos

	return r.referenciasConcellos, nil
}
