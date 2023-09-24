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

func (c ProcesosElectoraisController) GetDatosXerais(gc *gin.Context) {
	var id Id
	if err := gc.ShouldBindUri(&id); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindDatosXeraisById(id.Id)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func NewProcesosElectoraisController(e *gin.Engine, procesosElectoraisRepository domain.ProcesosElectoraisRepository) {
	c := &ProcesosElectoraisController{}
	c.repository = procesosElectoraisRepository
	e.GET("/procesos-electorais", c.GetProcesosElectorais)
	e.GET("/proceso-electoral/:id/datos-xerais", c.GetDatosXerais)
}
