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
	var uriParams struct {
		Id int `uri:"id"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindByProceso(uriParams.Id)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}
