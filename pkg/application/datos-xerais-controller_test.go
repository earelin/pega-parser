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

func TestGetDatosXerais(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByProceso", 1).
		Return(datosXerais, true)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, datosXeraisResponse, w.Body.String())
}

func TestGetDatosXerais_Null(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByProceso", 1).
		Return(domain.DatosXerais{}, false)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetDatosXeraisComunidadeAutonoma(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByComunidadeAutonoma", 1, 2).
		Return(datosXerais, true)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/comunidade-autonoma/2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, datosXeraisResponse, w.Body.String())
}

func TestGetDatosXeraisComunidadeAutonoma_Null(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByComunidadeAutonoma", 1, 2).
		Return(domain.DatosXerais{}, false)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/comunidade-autonoma/2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetDatosXeraisProvincia(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByProvincia", 1, 2).
		Return(datosXerais, true)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/provincia/2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, datosXeraisResponse, w.Body.String())
}

func TestGetDatosXeraisProvincia_Null(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByProvincia", 1, 2).
		Return(domain.DatosXerais{}, false)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/provincia/2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetDatosXeraisConcello(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByConcello", 1, 1001).
		Return(datosXerais, true)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/concello/1001", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, datosXeraisResponse, w.Body.String())
}

func TestGetDatosXeraisConcello_Null(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByConcello", 1, 1001).
		Return(domain.DatosXerais{}, false)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/concello/1001", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetDatosXeraisDistrito(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByDistrito", 1, 1001, 1).
		Return(datosXerais, true)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/concello/1001/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, datosXeraisResponse, w.Body.String())
}

func TestGetDatosXeraisDistrito_Null(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByDistrito", 1, 1001, 1).
		Return(domain.DatosXerais{}, false)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/concello/1001/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetDatosXeraisSeccion(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindBySeccion", 1, 1001, 1, 1).
		Return(datosXerais, true)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/concello/1001/1/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, datosXeraisResponse, w.Body.String())
}

func TestGetDatosXeraisSeccion_Null(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindBySeccion", 1, 1001, 1, 1).
		Return(domain.DatosXerais{}, false)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/concello/1001/1/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetDatosXeraisMesa(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByMesa", 1, 1001, 1, 1, "U").
		Return(datosXerais, true)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/concello/1001/1/1/U", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, datosXeraisResponse, w.Body.String())
}

func TestGetDatosXeraisMesa_Null(t *testing.T) {
	router := gin.Default()
	repository := new(DatosXeraisRepositoryMock)
	NewDatosXeraisController(router, repository)

	repository.On("FindByMesa", 1, 1001, 1, 1, "U").
		Return(domain.DatosXerais{}, false)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/1/datos-xerais/concello/1001/1/1/U", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

type DatosXeraisRepositoryMock struct {
	mock.Mock
}

func (m *DatosXeraisRepositoryMock) FindByProceso(id int) (domain.DatosXerais, bool) {
	args := m.Called(id)
	return args.Get(0).(domain.DatosXerais), args.Bool(1)
}

func (m *DatosXeraisRepositoryMock) FindByComunidadeAutonoma(id int, comunidadeAutonomaId int) (domain.DatosXerais, bool) {
	args := m.Called(id, comunidadeAutonomaId)
	return args.Get(0).(domain.DatosXerais), args.Bool(1)
}

func (m *DatosXeraisRepositoryMock) FindByProvincia(id int, provinciaId int) (domain.DatosXerais, bool) {
	args := m.Called(id, provinciaId)
	return args.Get(0).(domain.DatosXerais), args.Bool(1)
}

func (m *DatosXeraisRepositoryMock) FindByConcello(id int, concelloId int) (domain.DatosXerais, bool) {
	args := m.Called(id, concelloId)
	return args.Get(0).(domain.DatosXerais), args.Bool(1)
}

func (m *DatosXeraisRepositoryMock) FindByDistrito(id int, concelloId int, distritoId int) (domain.DatosXerais, bool) {
	args := m.Called(id, concelloId, distritoId)
	return args.Get(0).(domain.DatosXerais), args.Bool(1)
}

func (m *DatosXeraisRepositoryMock) FindBySeccion(id int, concelloId int, distritoId int, seccionId int) (domain.DatosXerais, bool) {
	args := m.Called(id, concelloId, distritoId, seccionId)
	return args.Get(0).(domain.DatosXerais), args.Bool(1)
}

func (m *DatosXeraisRepositoryMock) FindByMesa(id int, concelloId int, distritoId int, seccionId int, codigoMesa string) (domain.DatosXerais, bool) {
	args := m.Called(id, concelloId, distritoId, seccionId, codigoMesa)
	return args.Get(0).(domain.DatosXerais), args.Bool(1)
}

var datosXerais = domain.DatosXerais{
	CensoIne:  35169399,
	CensoCera: 2112809,
}

var datosXeraisResponse = `{"censoIne":35169399,"censoCera":2112809}`
