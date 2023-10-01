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

package application

import (
	"github.com/earelin/pega/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetResultadosCandidaturasByProceso(t *testing.T) {
	router := gin.Default()
	repository := new(ResultadosCandidaturasRepositoryMock)
	NewResultadosCandidaturasController(router, repository)

	repository.On("FindByProceso", 1).
		Return(resultadosCandidaturas, true)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/resultados/candidaturas", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resultadosCandidaturasResponse, w.Body.String())
}

type ResultadosCandidaturasRepositoryMock struct {
	mock.Mock
}

func (m *ResultadosCandidaturasRepositoryMock) FindByProceso(
	id int,
) ([]domain.ResultadoCandidatura, bool) {
	args := m.Called(id)
	return args.Get(0).([]domain.ResultadoCandidatura), args.Bool(1)
}

func (m *ResultadosCandidaturasRepositoryMock) FindByComunidadeAutonoma(
	id int, comunidadeAutonomaId int,
) ([]domain.ResultadoCandidatura, bool) {
	args := m.Called(id, comunidadeAutonomaId)
	return args.Get(0).([]domain.ResultadoCandidatura), args.Bool(1)
}

func (m *ResultadosCandidaturasRepositoryMock) FindByProvincia(
	id int, provinciaId int,
) ([]domain.ResultadoCandidatura, bool) {
	args := m.Called(id, provinciaId)
	return args.Get(0).([]domain.ResultadoCandidatura), args.Bool(1)
}

func (m *ResultadosCandidaturasRepositoryMock) FindByConcello(
	id int, concelloId int,
) ([]domain.ResultadoCandidatura, bool) {
	args := m.Called(id, concelloId)
	return args.Get(0).([]domain.ResultadoCandidatura), args.Bool(1)
}

func (m *ResultadosCandidaturasRepositoryMock) FindByDistrito(
	id int, concelloId int, distritoId int,
) ([]domain.ResultadoCandidatura, bool) {
	args := m.Called(id, concelloId, distritoId)
	return args.Get(0).([]domain.ResultadoCandidatura), args.Bool(1)
}

func (m *ResultadosCandidaturasRepositoryMock) FindBySeccion(
	id int, concelloId int, distritoId int, seccionId int,
) ([]domain.ResultadoCandidatura, bool) {
	args := m.Called(id, concelloId, distritoId, seccionId)
	return args.Get(0).([]domain.ResultadoCandidatura), args.Bool(1)
}

func (m *ResultadosCandidaturasRepositoryMock) FindByMesa(
	id int, concelloId int, distritoId int,
	seccionId int, codigoMesa string,
) ([]domain.ResultadoCandidatura, bool) {
	args := m.Called(id, concelloId, distritoId, seccionId, codigoMesa)
	return args.Get(0).([]domain.ResultadoCandidatura), args.Bool(1)
}

var resultadosCandidaturas = []domain.ResultadoCandidatura{
	{
		Candidatura: domain.Candidatura{
			Id:   1,
			Nome: "PARTIDO POPULAR",
		},
		Votos: 1022323,
	},
	{
		Candidatura: domain.Candidatura{
			Id:   2,
			Nome: "PARTIDO DOS SOCIALISTAS DE GALICIA-PSOE",
		},
		Votos: 1000000,
	},
	{
		Candidatura: domain.Candidatura{
			Id:   3,
			Nome: "BLOQUE NACIONALISTA GALEGO",
		},
		Votos: 900000,
	},
}

var resultadosCandidaturasResponse = `[{"candidatura":{"id":1,"nome":"PARTIDO POPULAR"},"votos":1022323},` +
	`{"candidatura":{"id":2,"nome":"PARTIDO DOS SOCIALISTAS DE GALICIA-PSOE"},"votos":1000000},` +
	`{"candidatura":{"id":3,"nome":"BLOQUE NACIONALISTA GALEGO"},"votos":900000}]`
