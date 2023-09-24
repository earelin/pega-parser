package base

import (
	"github.com/earelin/pega/pkg/domain"
	"github.com/gin-gonic/gin"
)

type EntidadesAdministrativasController struct {
	repository domain.EntidadesAdministrativasRepository
}

func NewEntidadesAdministrativasController(e *gin.Engine, repository domain.EntidadesAdministrativasRepository) {
	c := &EntidadesAdministrativasController{}
	c.repository = repository
	e.GET("/comunidades-autonomas", c.GetComunidadesAutonomas)
}

func (c *EntidadesAdministrativasController) GetComunidadesAutonomas(gc *gin.Context) {
	ca := c.repository.FindAllComunidadesAutonomas()
	gc.JSON(200, ca)
}
