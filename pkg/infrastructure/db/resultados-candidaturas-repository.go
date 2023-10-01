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

func NewResultadosCandidaturasSqlRepository(pool *sql.DB) *ResultadosCandidaturasSqlRepository {
	return &ResultadosCandidaturasSqlRepository{pool: pool}
}

type ResultadosCandidaturasSqlRepository struct {
	pool *sql.DB
}

func (r *ResultadosCandidaturasSqlRepository) FindByProceso(id int) ([]domain.ResultadoCandidatura, bool) {
	return r.find(`
	SELECT id, nome, SUM(votos) AS votos
	FROM (SELECT ce.id AS id, ce.nome, SUM(votos) AS votos
		FROM mesa_electoral_votos_candidatura mevc
		LEFT JOIN candidatura c ON c.id = mevc.candidatura_id
		LEFT JOIN candidatura ce ON ce.id = c.cabeceira_estatal
		WHERE c.proceso_electoral_id = ?
		GROUP BY ce.id
	UNION SELECT ce.id AS id, ce.nome, SUM(votos) AS votos
		FROM circunscripcion_cera_votos_candidatura ccvc
		LEFT JOIN candidatura c ON c.id = ccvc.candidatura_id
		LEFT JOIN candidatura ce ON ce.id = c.cabeceira_estatal
		WHERE c.proceso_electoral_id = ?
		GROUP BY ce.id) AS votos
	GROUP BY id, nome
	ORDER BY nome`, id, id)
}

func (r *ResultadosCandidaturasSqlRepository) FindByComunidadeAutonoma(
	id int, comunidadeAutonomaId int,
) ([]domain.ResultadoCandidatura, bool) {
	return []domain.ResultadoCandidatura{}, false
}

func (r *ResultadosCandidaturasSqlRepository) FindByProvincia(
	id int, provinciaId int,
) ([]domain.ResultadoCandidatura, bool) {
	return []domain.ResultadoCandidatura{}, false
}

func (r *ResultadosCandidaturasSqlRepository) FindByConcello(
	id int, concelloId int,
) ([]domain.ResultadoCandidatura, bool) {
	return []domain.ResultadoCandidatura{}, false
}

func (r *ResultadosCandidaturasSqlRepository) FindByDistrito(
	id int, concelloId int, distritoId int,
) ([]domain.ResultadoCandidatura, bool) {
	return []domain.ResultadoCandidatura{}, false
}

func (r *ResultadosCandidaturasSqlRepository) FindBySeccion(
	id int, concelloId int, distritoId int, seccionId int,
) ([]domain.ResultadoCandidatura, bool) {
	return []domain.ResultadoCandidatura{}, false
}

func (r *ResultadosCandidaturasSqlRepository) FindByMesa(
	id int, concelloId int, distritoId int, seccionId int, codigoMesa string,
) ([]domain.ResultadoCandidatura, bool) {
	return []domain.ResultadoCandidatura{}, false
}

func (r *ResultadosCandidaturasSqlRepository) find(
	query string, args ...any,
) ([]domain.ResultadoCandidatura, bool) {
	var resultados []domain.ResultadoCandidatura
	rows, err := r.pool.Query(query, args...)
	if err != nil {
		log.Printf("Error querying resultados: %s", err)
		return resultados, false
	}
	defer rows.Close()

	for rows.Next() {
		var resultado domain.ResultadoCandidatura
		err = rows.Scan(&resultado.Candidatura.Id, &resultado.Candidatura.Nome, &resultado.Votos)
		if err != nil {
			log.Printf("Error scanning resultados: %s", err)
		}
		resultados = append(resultados, resultado)
	}

	return resultados, true
}
