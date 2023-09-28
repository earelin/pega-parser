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

package db

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
	rows, err := r.pool.Query("SELECT id, data, tipo, ambito FROM proceso_electoral ORDER BY data DESC")
	if err != nil {
		log.Printf("Error querying procesos: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var proceso domain.ProcesoElectoral
		var dataRaw string
		var ambitoRaw sql.NullInt16
		err = rows.Scan(&proceso.Id, &dataRaw, &proceso.Tipo, &ambitoRaw)
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
