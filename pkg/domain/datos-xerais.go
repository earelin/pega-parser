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

package domain

type DatosXerais struct {
	CensoIne  int `json:"censoIne"`
	CensoCera int `json:"censoCera"`
}

type DatosXeraisRepository interface {
	FindByProceso(id int) (DatosXerais, bool)
	FindByComunidadeAutonoma(id int, comunidadeAutonomaId int) (DatosXerais, bool)
	FindByProvincia(id int, provinciaId int) (DatosXerais, bool)
	FindByConcello(id int, concelloId int) (DatosXerais, bool)
	FindByDistrito(id int, concelloId int, distritoId int) (DatosXerais, bool)
	FindBySeccion(id int, concelloId int, distritoId int, seccionId int) (DatosXerais, bool)
	FindByMesa(id int, concelloId int, distritoId int, seccionId int, codigoMesa string) (DatosXerais, bool)
}
