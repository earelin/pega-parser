package application

import (
	"github.com/earelin/pega/pkg/domain"
	"github.com/gin-gonic/gin"
)

type ProcesosElectoraisController struct {
	procesosElectoraisRepository domain.ProcesosElectoraisRepository
}

func (c ProcesosElectoraisController) GetProcesosElectorais(gc *gin.Context) {
	procesosElectorais := c.procesosElectoraisRepository.FindAll()
	gc.JSON(200, procesosElectorais)
}

func NewProcesosElectoraisController(e *gin.Engine, procesosElectoraisRepository domain.ProcesosElectoraisRepository) {
	c := &ProcesosElectoraisController{}
	c.procesosElectoraisRepository = procesosElectoraisRepository
	e.GET("/procesos-electorais", c.GetProcesosElectorais)
}
