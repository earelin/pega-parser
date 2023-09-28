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
	assert.Equal(t, processosElectoraisResult, w.Body.String())
}

type ProcesosElectoraisRepositoryMock struct {
	mock.Mock
}

func (m *ProcesosElectoraisRepositoryMock) FindAll() []domain.ProcesoElectoral {
	args := m.Called()
	return args.Get(0).([]domain.ProcesoElectoral)
}

var procesosElectorais = []domain.ProcesoElectoral{
	{
		Id:     1,
		Data:   time.Date(2019, time.May, 26, 0, 0, 0, 0, time.UTC),
		Tipo:   7,
		Ambito: 0,
	},
	{
		Id:     2,
		Data:   time.Date(2018, time.February, 13, 0, 0, 0, 0, time.UTC),
		Tipo:   6,
		Ambito: 1,
	},
}

var processosElectoraisResult = `[{"id":1,"data":"2019-05-26T00:00:00Z","tipo":7,"ambito":0},{"id":2,"data":"2018-02-13T00:00:00Z","tipo":6,"ambito":1}]`
