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
	t.Skip("Not implemented")
}

func TestGetDatosXeraisProvincia(t *testing.T) {
	t.Skip("Not implemented")
}

func TestGetDatosXeraisConcello(t *testing.T) {
	t.Skip("Not implemented")
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
