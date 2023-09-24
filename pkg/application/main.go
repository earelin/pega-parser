package application

import (
	"github.com/earelin/pega/pkg/application/base"
	"github.com/earelin/pega/pkg/domain"
	"github.com/gin-gonic/gin"
)

func ConfigureApplicationLayer(e *gin.Engine, ear domain.EntidadesAdministrativasRepository) {
	MonitoringConfig(e)
	base.NewEntidadesAdministrativasController(e, ear)
}
