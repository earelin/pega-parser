package application

import (
	"github.com/earelin/pega/pkg/domain"
	"github.com/gin-gonic/gin"
)

type ProcesosElectoraisController struct {
	repository domain.ProcesosElectoraisRepository
}

func (c ProcesosElectoraisController) GetProcesosElectorais(gc *gin.Context) {
	procesosElectorais := c.repository.FindAll()
	gc.JSON(200, procesosElectorais)
}

func NewProcesosElectoraisController(e *gin.Engine, procesosElectoraisRepository domain.ProcesosElectoraisRepository) {
	c := &ProcesosElectoraisController{}
	c.repository = procesosElectoraisRepository
	e.GET("/procesos-electorais", c.GetProcesosElectorais)
}
