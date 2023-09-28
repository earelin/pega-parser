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

type Resultados struct {
	VotosBlanco       int `json:"votosBranco"`
	VotosNulos        int `json:"votosNulos"`
	VotosCandidaturas int `json:"votosCandidaturas"`
}

type ResultadoCandidatura struct {
	Candidatura Candidatura
	Votos       int
}

type ResultadosRepository interface {
	FindByProceso(id int) (Resultados, bool)
	FindByComunidadeAutonoma(id int, comunidadeAutonomaId int) (Resultados, bool)
	FindByProvincia(id int, provinciaId int) (Resultados, bool)
	FindByConcello(id int, concelloId int) (Resultados, bool)
	FindByDistrito(id int, concelloId int, distritoId int) (Resultados, bool)
	FindBySeccion(id int, concelloId int, distritoId int, seccionId int) (Resultados, bool)
	FindByMesa(id int, concelloId int, distritoId int, seccionId int, codigoMesa string) (Resultados, bool)
}

type ResultadosCandidaturasRepository interface {
	FindByProceso(id int) ([]ResultadoCandidatura, bool)
	FindByComunidadeAutonoma(id int, comunidadeAutonomaId int) ([]ResultadoCandidatura, bool)
	FindByProvincia(id int, provinciaId int) ([]ResultadoCandidatura, bool)
	FindByConcello(id int, concelloId int) ([]ResultadoCandidatura, bool)
	FindByDistrito(id int, concelloId int, distritoId int) ([]ResultadoCandidatura, bool)
	FindBySeccion(id int, concelloId int, distritoId int, seccionId int) ([]ResultadoCandidatura, bool)
	FindByMesa(id int, concelloId int, distritoId int, seccionId int, codigoMesa string) ([]ResultadoCandidatura, bool)
}
