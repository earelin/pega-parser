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
		SELECT pe.id AS id, data, tpe.id AS tipo_id, tpe.nome AS tipo_nome, ambito
		FROM proceso_electoral pe
		LEFT JOIN tipo_proceso_electoral tpe ON tpe.id = pe.tipo
		ORDER BY data DESC`)
	if err != nil {
		log.Printf("Error querying procesos: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var proceso domain.ProcesoElectoral
		var dataRaw string
		var ambitoRaw sql.NullInt16
		err = rows.Scan(&proceso.Id, &dataRaw, &proceso.Tipo.Id, &proceso.Tipo.Nome, &ambitoRaw)
		if err != nil {
			log.Printf("Error scanning procesos: %s", err)
		}
		proceso.Data, err = time.Parse("2006-01-02T15:04:05Z", dataRaw)
		if err != nil {
			log.Printf("Error parsing date: %s. %s", dataRaw, err)
		}
		proceso.Ambito = int(ambitoRaw.Int16)
		procesos = append(procesos, proceso)
	}

	return procesos
}

func (r *ProcesosElectoraisSqlRepository) FindById(id int) (domain.ProcesoElectoral, bool) {
	var proceso domain.ProcesoElectoral
	row := r.pool.QueryRow(`
		SELECT pe.id AS id, data, tpe.id AS tipo_id, tpe.nome AS tipo_nome, ambito
		FROM proceso_electoral pe
		LEFT JOIN tipo_proceso_electoral tpe ON tpe.id = pe.tipo
		WHERE pe.id = ?`, id)

	var ambitoRaw sql.NullInt16
	var err = row.Scan(&proceso.Id, &proceso.Data, &proceso.Tipo.Id, &proceso.Tipo.Nome, &ambitoRaw)
	if err != nil {
		log.Printf("Error scanning proceso: %s", err)
		return proceso, false
	}
	proceso.Ambito = int(ambitoRaw.Int16)

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
