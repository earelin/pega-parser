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
	"time"
)

func TestGetProcesosElectorais(t *testing.T) {
	router := gin.Default()
	repository := new(ProcesosElectoraisRepositoryMock)
	NewProcesosElectoraisController(router, repository)

	repository.On("FindAll").
		Return(procesosElectorais)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/procesos-electorais", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, procesosElectoraisResult, w.Body.String())
}

func TestGetProcesoElectoral(t *testing.T) {
	router := gin.Default()
	repository := new(ProcesosElectoraisRepositoryMock)
	NewProcesosElectoraisController(router, repository)

	repository.On("FindById", 2).
		Return(procesosElectorais[1], true)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/proceso-electoral/2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, procesoElectoralResult, w.Body.String())
}

func TestGetProcesosElectoraisTipos(t *testing.T) {
	router := gin.Default()
	repository := new(ProcesosElectoraisRepositoryMock)
	NewProcesosElectoraisController(router, repository)

	repository.On("FindAllTipos").
		Return(procesosElectoraisTipos)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/procesos-electorais/tipos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, procesosElectoraisTiposResult, w.Body.String())
}

type ProcesosElectoraisRepositoryMock struct {
	mock.Mock
}

func (m *ProcesosElectoraisRepositoryMock) FindAll() []domain.ProcesoElectoral {
	args := m.Called()
	return args.Get(0).([]domain.ProcesoElectoral)
}

func (m *ProcesosElectoraisRepositoryMock) FindById(id int) (domain.ProcesoElectoral, bool) {
	args := m.Called(id)
	return args.Get(0).(domain.ProcesoElectoral), args.Bool(1)
}

func (m *ProcesosElectoraisRepositoryMock) FindAllTipos() []domain.TipoProcesoElectoral {
	args := m.Called()
	return args.Get(0).([]domain.TipoProcesoElectoral)
}

var procesosElectorais = []domain.ProcesoElectoral{
	{
		Id:   1,
		Data: time.Date(2019, time.May, 26, 0, 0, 0, 0, time.UTC),
		Tipo: domain.TipoProcesoElectoral{
			Id:   7,
			Nome: "Autonómicas",
		},
		Ambito: 1,
	},
	{
		Id:   2,
		Data: time.Date(2018, time.February, 13, 0, 0, 0, 0, time.UTC),
		Tipo: domain.TipoProcesoElectoral{
			Id:   6,
			Nome: "Xerais",
		},
		Ambito: 0,
	},
}

var procesosElectoraisResult = `[{"id":1,"data":"2019-05-26T00:00:00Z","tipo":{"id":7,"nome":"Autonómicas"},"ambito":1},` +
	`{"id":2,"data":"2018-02-13T00:00:00Z","tipo":{"id":6,"nome":"Xerais"},"ambito":0}]`

var procesoElectoralResult = `{"id":2,"data":"2018-02-13T00:00:00Z","tipo":{"id":6,"nome":"Xerais"},"ambito":0}`

var procesosElectoraisTipos = []domain.TipoProcesoElectoral{
	{
		Id:   1,
		Nome: "Eleccións xerais",
	},
	{
		Id:   2,
		Nome: "Eleccións autonómicas",
	},
}

var procesosElectoraisTiposResult = `[{"id":1,"nome":"Eleccións xerais"},` +
	`{"id":2,"nome":"Eleccións autonómicas"}]`
