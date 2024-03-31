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
)

type DivisionsAdministrativasSqlRepository struct {
	pool *sql.DB
}

func NewDivisionsAdministrativasSqlRepository(pool *sql.DB) *DivisionsAdministrativasSqlRepository {
	return &DivisionsAdministrativasSqlRepository{pool: pool}
}

func (r *DivisionsAdministrativasSqlRepository) FindAllComunidadesAutonomas() []domain.DivisionAdministrativa {
	return r.findEntidades("SELECT id, nome FROM comunidade_autonoma ORDER BY nome")
}

func (r *DivisionsAdministrativasSqlRepository) FindAllProvincias() []domain.DivisionAdministrativa {
	return r.findEntidades("SELECT id, nome FROM provincia ORDER BY nome")
}

func (r *DivisionsAdministrativasSqlRepository) FindAllProvinciasByComunidadeAutonoma(
	caId int,
) []domain.DivisionAdministrativa {
	return r.findEntidades(`
		SELECT id, nome
		FROM provincia
		WHERE comunidade_autonoma_id = ? ORDER BY nome`,
		caId)
}

func (r *DivisionsAdministrativasSqlRepository) FindAllConcellosByProvincia(
	pId int,
) []domain.DivisionAdministrativa {
	return r.findEntidades(`
		SELECT id, nome
		FROM concello
		WHERE provincia_id = ? ORDER BY nome`,
		pId)
}

func (r *DivisionsAdministrativasSqlRepository) FindAllConcellosByName(
	name string,
) []domain.DivisionAdministrativa {
	return r.findEntidades(`
		SELECT id, nome
		FROM concello
		WHERE nome LIKE ? ORDER BY nome`,
		"%"+name+"%")
}

func (r *DivisionsAdministrativasSqlRepository) findEntidades(
	sql string, args ...any,
) []domain.DivisionAdministrativa {
	var entidades []domain.DivisionAdministrativa
	rows, err := r.pool.Query(sql, args...)
	if err != nil {
		log.Printf("Error querying entidaes: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var provincia domain.DivisionAdministrativa
		err = rows.Scan(&provincia.Id, &provincia.Nome)
		if err != nil {
			log.Printf("Error scanning entidades: %s", err)
		}
		entidades = append(entidades, provincia)
	}

	return entidades
}
