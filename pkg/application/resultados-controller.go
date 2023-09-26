package application

import (
	"github.com/earelin/pega/pkg/domain"
	"github.com/gin-gonic/gin"
)

func NewResultadosController(e *gin.Engine, repository domain.ResultadosRepository) {
	c := &ResultadosController{repository: repository}
	e.GET("/proceso-electoral/:id/resultados", c.GetResultadosByProceso)
}

type ResultadosController struct {
	repository domain.ResultadosRepository
}

func (c ResultadosController) GetResultadosByProceso(gc *gin.Context) {

}
