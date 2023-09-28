/*
 * This program is free software: you can redistribute it and/or modify it under
 * the terms of the GNU General Public License as published by the Free Software
 * Foundation, either version 3 of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License
 * for more details.
 *
 * You should have received a copy of the GNU General Public License along with
 * this program. If not, see <https://www.gnu.org/licenses/>.
 */

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/election"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"time"
)

type Repository struct {
	pool *sql.DB
	ctx  context.Context
}

func NewRepository(c Config, ctx context.Context) (*Repository, error) {
	var r Repository

	var pool, err = sql.Open("sqlite3", c.Filename)
	if err != nil {
		return nil, err
	}

	pool.SetConnMaxLifetime(0)

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
	const inserirProcesoElectoral = "INSERT INTO proceso_electoral (tipo, ambito, data) VALUES (?, ?, ?)"

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

func (r *Repository) CreateCandidaturas(procesoElectoral int64, candidatures []election.Candidatura) error {
	const inserirCandidatura = `INSERT INTO candidatura (proceso_electoral_id, id, siglas, nome, cabeceira_estatal,
                        cabeceira_autonomica, cabeceira_provincial) VALUES (?, ?, ?, ?, ?, ?, ?)`

	for _, c := range candidatures {
		_, err := r.pool.ExecContext(r.ctx, inserirCandidatura, procesoElectoral, c.Codigo, c.Siglas, c.Nome,
			c.CabeceiraEstatal, c.CabeceiraAutonomica, c.CabeceiraProvincial)
		if err != nil {
			return fmt.Errorf("non foi posible gardar unha candidatura: %w", err)
		}
	}

	return nil
}

func (r *Repository) CrearListasECandidatos(procesoElectoral int64, listaCandidatos []election.Candidate) error {
	for _, c := range listaCandidatos {
		var _, err = r.pool.ExecContext(r.ctx,
			`INSERT INTO candidato (proceso_electoral_id, candidatura_id, ambito, posicion, titular,
                      					 nombre, apelidos, eleito)
				   VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			procesoElectoral, c.CodigoCandidatura, c.AmbitoTerritorial, c.Posicion, c.Titular, c.Nome, c.Apelidos,
			c.FoiEleito)
		if err != nil {
			return fmt.Errorf("non foi posible gardar un candidato: %w", err)
		}
	}

	return nil
}

func (r *Repository) CrearMesasElectorais(procesoElectoral int64, mesas []election.MesaElectoral) error {
	var err error

	var cInserts, mInserts []string
	var cValues, mValues []interface{}

	for i, m := range mesas {
		if i%1000 == 0 && i != 0 {
			err = r.inserirMultipleMesasECircunscripcionsCera(cInserts, mInserts, cValues, mValues)
			if err != nil {
				return fmt.Errorf("non se puideron gardar os dataos de mesas e circunscipcions: %w", err)
			}

			cInserts = []string{}
			cValues = []interface{}{}
			mInserts = []string{}
			mValues = []interface{}{}
		}

		if m.CodigoProvincia == 99 {
			continue
		} else if m.CodigoConcello == 999 {
			cInserts = append(cInserts, "(?, ?, ?, ?, ?, ?)")
			cValues = append(cValues, procesoElectoral, m.CodigoProvincia,
				m.CensoEscrutinioOCera, m.VotosBlanco, m.VotosNulos, m.VotosCandidaturas)
		} else {
			mInserts = append(mInserts, "(?, ?, ?, ?, ?, ?, ?, ?, ?)")
			concelloId := m.CodigoProvincia*1000 + m.CodigoConcello
			mValues = append(mValues, procesoElectoral, concelloId, m.Distrito,
				m.Seccion, m.CodigoMesa, m.CensoIne, m.VotosBlanco, m.VotosNulos, m.VotosCandidaturas)
		}
	}

	err = r.inserirMultipleMesasECircunscripcionsCera(cInserts, mInserts, cValues, mValues)
	if err != nil {
		return fmt.Errorf("non se puideron gardar os dataos de mesas e circunscipcions: %w", err)
	}

	return nil
}

func (r *Repository) inserirMultipleMesasECircunscripcionsCera(cInserts, mInserts []string, cValues, mValues []interface{}) error {
	const inserirCircunscripcionCera = `
		INSERT INTO circunscripcion_cera (proceso_electoral_id, provincia_id, censo, votos_blanco,
		                                 votos_nulos, votos_candidaturas)
		VALUES `
	const inserirMesaElectoral = `
		INSERT INTO mesa_electoral (proceso_electoral_id, concello_id, distrito, seccion, codigo, censo,
		                           votos_blanco, votos_nulos, votos_candidaturas)
		VALUES `

	var err error

	if len(cInserts) > 0 {
		sqlInsert := inserirCircunscripcionCera + strings.Join(cInserts, ",")
		_, err = r.pool.ExecContext(r.ctx, sqlInsert, cValues...)
		if err != nil {
			return fmt.Errorf("non se puideron gardar as circunscripcions CERA: %w", err)
		}
	}

	if len(mInserts) > 0 {
		sqlInsert := inserirMesaElectoral + strings.Join(mInserts, ",")
		_, err = r.pool.ExecContext(r.ctx, sqlInsert, mValues...)
		if err != nil {
			return fmt.Errorf("non se puideron gardar as mesas electorais: %w", err)
		}
	}

	return nil
}

func (r *Repository) CrearVotosEnMesasElectorais(procesoElectoral int64, votos []election.VotosMesaElectoral) error {

	var err error

	var cInserts, mInserts []string
	var cValues, mValues []interface{}

	for i, v := range votos {
		if i%1000 == 0 && i != 0 {
			err = r.insertirMultiplesVotos(cInserts, mInserts, cValues, mValues)
			if err != nil {
				return fmt.Errorf("non se puideron gardar os vortos dunha mesa ou circunscripcion: %w", err)
			}

			cInserts = []string{}
			cValues = []interface{}{}
			mInserts = []string{}
			mValues = []interface{}{}
		}

		if v.CodigoProvincia == 99 || v.Votos == 0 {
			continue
		} else if v.CodigoConcello == 999 {
			cInserts = append(cInserts, "(?, ?, ?, ?, ?)")
			cValues = append(cValues, procesoElectoral, v.CodigoProvincia, v.CandidaturaOuSenador, nil, v.Votos)
		} else {
			mInserts = append(mInserts, "(?, ?, ?, ?, ?, ?, ?, ?)")
			mValues = append(mValues, procesoElectoral, v.CodigoProvincia*1000+v.CodigoConcello, v.Distrito, v.Seccion,
				v.CodigoMesa, v.CandidaturaOuSenador, nil, v.Votos)
		}
	}

	err = r.insertirMultiplesVotos(cInserts, mInserts, cValues, mValues)
	if err != nil {
		return fmt.Errorf("non se puideron gardar os vortos dunha mesa ou circunscripcion: %w", err)
	}

	return nil
}

func (r *Repository) insertirMultiplesVotos(cInserts, mInserts []string, cValues, mValues []interface{}) error {
	const inserirVotosCircunscripcionCera = `
		INSERT INTO circunscripcion_cera_votos_candidatura (proceso_electoral_id, provincia_id, candidatura_id,
				    									   posicion, votos)
		VALUES `
	const inserirVotosMesaElectoral = `
		INSERT INTO mesa_electoral_votos_candidatura (proceso_electoral_id, concello_id, distrito, seccion,
		                                             codigo, candidatura_id, posicion, votos)
		VALUES `

	var err error

	if len(cInserts) > 0 {
		sqlInsert := inserirVotosCircunscripcionCera + strings.Join(cInserts, ",")
		_, err = r.pool.ExecContext(r.ctx, sqlInsert, cValues...)
		if err != nil {
			return fmt.Errorf("non se puideron gardar os votos das circunscripcions CERA: %w", err)
		}
	}

	if len(mInserts) > 0 {
		sqlInsert := inserirVotosMesaElectoral + strings.Join(mInserts, ",")
		_, err = r.pool.ExecContext(r.ctx, sqlInsert, mValues...)
		if err != nil {
			return fmt.Errorf("non se puideron gardar os votos das mesas electorais: %w", err)
		}
	}

	return nil
}
