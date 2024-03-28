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
)

type DatosXeraisSqlRepository struct {
	pool *sql.DB
}

func NewDatosXeraisSqlRepository(pool *sql.DB) *DatosXeraisSqlRepository {
	return &DatosXeraisSqlRepository{pool: pool}
}

func (r *DatosXeraisSqlRepository) FindByProceso(id int) (domain.DatosXerais, bool) {
	return r.find("SELECT censo_ine, censo_cera FROM datos_xerais WHERE id = ?", id)
}

func (r *DatosXeraisSqlRepository) FindByComunidadeAutonoma(
	id int, comunidadeAutonomaId int,
) (domain.DatosXerais, bool) {
	return r.find(`
		SELECT censo_ine, censo_cera
		FROM datos_xerais_autonomicos
		WHERE id = ? AND comunidade_autonoma_id = ?`,
		id, comunidadeAutonomaId)
}

func (r *DatosXeraisSqlRepository) FindByProvincia(
	id int, provinciaId int,
) (domain.DatosXerais, bool) {
	return r.find(`
		SELECT censo_ine, censo_cera
		FROM datos_xerais_provincias
		WHERE id = ? AND provincia_id = ?`,
		id, provinciaId)
}

func (r *DatosXeraisSqlRepository) FindByConcello(
	id int, concelloId int,
) (domain.DatosXerais, bool) {
	return r.find(`
		SELECT censo_ine, 0
		FROM datos_xerais_concellos
		WHERE id = ? AND concello_id = ?`,
		id, concelloId)
}

func (r *DatosXeraisSqlRepository) FindByDistrito(
	id int, concelloId int, distritoId int,
) (domain.DatosXerais, bool) {
	return r.find(`
		SELECT censo_ine, 0
		FROM datos_xerais_distritos
		WHERE id = ? AND concello_id = ? AND distrito = ?`,
		id, concelloId, distritoId)
}

func (r *DatosXeraisSqlRepository) FindBySeccion(
	id int, concelloId int, distritoId int, seccionId int,
) (domain.DatosXerais, bool) {
	return r.find(`
		SELECT censo_ine, 0
		FROM datos_xerais_seccions
		WHERE id = ? AND concello_id = ? AND distrito = ? AND seccion = ?`,
		id, concelloId, distritoId, seccionId)
}

func (r *DatosXeraisSqlRepository) FindByMesa(
	id int, concelloId int, distritoId int, seccionId int, codigoMesa string,
) (domain.DatosXerais, bool) {
	return r.find(`
		SELECT censo, 0
		FROM mesa_electoral
		WHERE proceso_electoral_id = ? AND concello_id = ? AND
		      distrito = ? AND seccion = ? AND codigo = ?`,
		id, concelloId, distritoId, seccionId, codigoMesa)
}

func (r *DatosXeraisSqlRepository) find(query string, args ...any) (domain.DatosXerais, bool) {
	var datosXerais domain.DatosXerais
	row := r.pool.QueryRow(query, args...)

	var err = row.Scan(&datosXerais.CensoIne, &datosXerais.CensoCera)
	if errors.Is(err, sql.ErrNoRows) {
		return datosXerais, false
	} else if err != nil {
		log.Printf("Error scanning procesos: %s", err)
	}

	return datosXerais, true
}
