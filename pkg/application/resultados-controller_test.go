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

func TestGetResultadosByProceso(t *testing.T) {
	router := gin.Default()
	repository := new(ResultadosRepositoryMock)
	NewResultadosController(router, repository)

	repository.On("FindByProceso", 1).
		Return(resultadosGlobais, true)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/resultados", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resultadosGlobaisResponse, w.Body.String())
}

func TestGetResultadosByProceso_Null(t *testing.T) {
	router := gin.Default()
	repository := new(ResultadosRepositoryMock)
	NewResultadosController(router, repository)

	repository.On("FindByProceso", 1).
		Return(domain.Resultados{}, false)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/resultados", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

type ResultadosRepositoryMock struct {
	mock.Mock
}

func (m *ResultadosRepositoryMock) FindByProceso(id int) (domain.Resultados, bool) {
	args := m.Called(id)
	return args.Get(0).(domain.Resultados), args.Bool(1)
}

func (m *ResultadosRepositoryMock) FindByComunidadeAutonoma(
	id int, comunidadeAutonomaId int,
) (domain.Resultados, bool) {
	args := m.Called(id, comunidadeAutonomaId)
	return args.Get(0).(domain.Resultados), args.Bool(1)
}

func (m *ResultadosRepositoryMock) FindByProvincia(
	id int, provinciaId int,
) (domain.Resultados, bool) {
	args := m.Called(id, provinciaId)
	return args.Get(0).(domain.Resultados), args.Bool(1)
}

func (m *ResultadosRepositoryMock) FindByConcello(
	id int, concelloId int,
) (domain.Resultados, bool) {
	args := m.Called(id, concelloId)
	return args.Get(0).(domain.Resultados), args.Bool(1)
}

func (m *ResultadosRepositoryMock) FindByDistrito(
	id int, concelloId int, distritoId int,
) (domain.Resultados, bool) {
	args := m.Called(id, concelloId, distritoId)
	return args.Get(0).(domain.Resultados), args.Bool(1)
}

func (m *ResultadosRepositoryMock) FindBySeccion(
	id int, concelloId int, distritoId int, seccionId int,
) (domain.Resultados, bool) {
	args := m.Called(id, concelloId, distritoId, seccionId)
	return args.Get(0).(domain.Resultados), args.Bool(1)
}

func (m *ResultadosRepositoryMock) FindByMesa(
	id int, concelloId int, distritoId int,
	seccionId int, codigoMesa string,
) (domain.Resultados, bool) {
	args := m.Called(id, concelloId, distritoId, seccionId, codigoMesa)
	return args.Get(0).(domain.Resultados), args.Bool(1)
}

var resultadosGlobais = domain.Resultados{
	VotosBlanco:       217005,
	VotosNulos:        194345,
	VotosCandidaturas: 22258966,
}

var resultadosGlobaisResponse = `{"votosBranco":217005,"votosNulos":194345,"votosCandidaturas":22258966}`
