package application

import (
	"github.com/earelin/pega/pkg/domain"
	"github.com/gin-gonic/gin"
)

func ConfigureApplicationLayer(e *gin.Engine,
	ear domain.EntidadesAdministrativasRepository,
	per domain.ProcesosElectoraisRepository) {
	MonitoringConfig(e)
	NewEntidadesAdministrativasController(e, ear)
	NewProcesosElectoraisController(e, per)
}
