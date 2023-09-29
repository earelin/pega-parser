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
)

func NewResultadosSqlRepository(pool *sql.DB) *ResultadosSqlRepository {
	return &ResultadosSqlRepository{pool: pool}
}

type ResultadosSqlRepository struct {
	pool *sql.DB
}

func (r *ResultadosSqlRepository) FindByProceso(id int) (domain.Resultados, bool) {
	return r.find(`
	SELECT SUM(votos_branco)   AS votos_branco,
       SUM(votos_nulos)        AS votos_nulos,
       SUM(votos_candidaturas) AS votos_candidaturas
	FROM (SELECT SUM(votos_blanco)       AS votos_branco,
             SUM(votos_nulos)        AS votos_nulos,
             SUM(votos_candidaturas) AS votos_candidaturas
      FROM mesa_electoral me
      WHERE me.proceso_electoral_id = ?
      UNION
      SELECT SUM(votos_blanco)       AS votos_branco,
             SUM(votos_nulos)        AS votos_nulos,
             SUM(votos_candidaturas) AS votos_candidaturas
      FROM circunscripcion_cera cc
      WHERE cc.proceso_electoral_id = ?) AS votos`, id, id)
}

func (r *ResultadosSqlRepository) FindByComunidadeAutonoma(
	id int, comunidadeAutonomaId int,
) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) FindByProvincia(
	id int, provinciaId int,
) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) FindByConcello(
	id int, concelloId int,
) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) FindByDistrito(
	id int, concelloId int, distritoId int,
) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) FindBySeccion(
	id int, concelloId int, distritoId int, seccionId int,
) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) FindByMesa(
	id int, concelloId int, distritoId int, seccionId int, codigoMesa string,
) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) find(query string, args ...any) (domain.Resultados, bool) {
	var resultados domain.Resultados

	row := r.pool.QueryRow(query, args...)
	err := row.Scan(&resultados.VotosBlanco, &resultados.VotosNulos, &resultados.VotosCandidaturas)
	if err != nil {
		log.Printf("Error scanning resultados: %s", err)
		return resultados, false
	}

	return resultados, true
}
