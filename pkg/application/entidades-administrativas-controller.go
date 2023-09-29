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
)

type EntidadesAdministrativasController struct {
	repository domain.EntidadesAdministrativasRepository
}

func NewEntidadesAdministrativasController(
	e *gin.Engine,
	repository domain.EntidadesAdministrativasRepository,
) {
	c := &EntidadesAdministrativasController{}
	c.repository = repository
	e.GET("/comunidades-autonomas", c.GetComunidadesAutonomas)
	e.GET("/comunidade-autonoma/:id/provincias", c.GetComunidadesAutonomaProvincias)
	e.GET("/provincias", c.GetProvincias)
	e.GET("/provincia/:id/concellos", c.GetConcellosProvincia)
	e.GET("/concellos/pescuda/:search", c.GetConcellosByName)
}

func (c *EntidadesAdministrativasController) GetComunidadesAutonomas(gc *gin.Context) {
	ca := c.repository.FindAllComunidadesAutonomas()
	gc.JSON(200, ca)
}

func (c *EntidadesAdministrativasController) GetProvincias(gc *gin.Context) {
	p := c.repository.FindAllProvincias()
	gc.JSON(200, p)
}

func (c *EntidadesAdministrativasController) GetComunidadesAutonomaProvincias(gc *gin.Context) {
	var id Id
	if err := gc.ShouldBindUri(&id); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}
	ps := c.repository.FindAllProvinciasByComunidadeAutonoma(id.Id)

	if len(ps) == 0 {
		gc.Status(404)
	} else {
		gc.JSON(200, ps)
	}
}

func (c *EntidadesAdministrativasController) GetConcellosProvincia(gc *gin.Context) {
	var id Id
	if err := gc.ShouldBindUri(&id); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	cs := c.repository.FindAllConcellosByProvincia(id.Id)

	if len(cs) == 0 {
		gc.Status(404)
	} else {
		gc.JSON(200, cs)
	}
}

func (c *EntidadesAdministrativasController) GetConcellosByName(gc *gin.Context) {
	var queryParams Search
	if err := gc.ShouldBindUri(&queryParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}
	cs := c.repository.FindAllConcellosByName(queryParams.Search)
	gc.JSON(200, cs)
}
