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
	"database/sql"
	"errors"
	"github.com/earelin/pega/pkg/domain"
	"log"
	"time"
)

type ProcesosElectoraisSqlRepository struct {
	pool *sql.DB
}

func NewProcesosElectoraisSqlRepository(pool *sql.DB) *ProcesosElectoraisSqlRepository {
	return &ProcesosElectoraisSqlRepository{pool: pool}
}

func (r *ProcesosElectoraisSqlRepository) FindAll() []domain.ProcesoElectoral {
	var procesos []domain.ProcesoElectoral
	rows, err := r.pool.Query(`
		SELECT pe.id AS id, data, pe.tipo_id, tpe.nome AS tipo_nome, pe.ambito_id, ca.nome AS ambito_nome
		FROM proceso_electoral pe
        	LEFT JOIN tipo_proceso_electoral tpe ON tpe.id = pe.tipo_id
         	LEFT JOIN comunidade_autonoma ca ON ca.id = pe.ambito_id
		ORDER BY data DESC`)
	if err != nil {
		log.Printf("Error querying procesos: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var proceso domain.ProcesoElectoral
		var ambito domain.DivisionAdministrativa
		var dataRaw string

		err = rows.Scan(&proceso.Id, &dataRaw, &proceso.Tipo.Id, &proceso.Tipo.Nome, &ambito.Id, &ambito.Nome)
		if err != nil {
			log.Printf("Error scanning procesos: %s", err)
		}

		proceso.Data, err = time.Parse("2006-01-02T15:04:05Z", dataRaw)
		if err != nil {
			log.Printf("Error parsing date: %s. %s", dataRaw, err)
		}

		if ambito.Id != 0 {
			proceso.Ambito = &ambito
		}

		procesos = append(procesos, proceso)
	}

	return procesos
}

func (r *ProcesosElectoraisSqlRepository) FindById(id int) (domain.ProcesoElectoral, bool) {
	row := r.pool.QueryRow(`
		SELECT pe.id AS id, data, pe.tipo_id, tpe.nome AS tipo_nome, pe.ambito_id, ca.nome AS ambito_nome
		FROM proceso_electoral pe
        	LEFT JOIN tipo_proceso_electoral tpe ON tpe.id = pe.tipo_id
         	LEFT JOIN comunidade_autonoma ca ON ca.id = pe.ambito_id
		WHERE pe.id = ?`, id)

	var proceso domain.ProcesoElectoral
	var ambito domain.DivisionAdministrativa
	var dataRaw string

	var err = row.Scan(&proceso.Id, &dataRaw, &proceso.Tipo.Id, &proceso.Tipo.Nome, &ambito.Id, &ambito.Nome)
	if err != nil {
		log.Printf("Error scanning proceso: %s", err)
		if errors.Is(err, sql.ErrNoRows) {
			return proceso, false
		}
	}

	if ambito.Id != 0 {
		proceso.Ambito = &ambito
	}

	return proceso, true
}

func (r *ProcesosElectoraisSqlRepository) FindAllTipos() []domain.TipoProcesoElectoral {
	var tipos []domain.TipoProcesoElectoral
	rows, err := r.pool.Query(`
		SELECT id, nome
		FROM tipo_proceso_electoral`)
	if err != nil {
		log.Printf("Error querying tipos: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tipo domain.TipoProcesoElectoral
		err = rows.Scan(&tipo.Id, &tipo.Nome)
		if err != nil {
			log.Printf("Error scanning tipos: %s", err)
		}
		tipos = append(tipos, tipo)
	}

	return tipos
}
