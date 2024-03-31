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

type ProcesosElectoraisController struct {
	repository domain.ProcesosElectoraisRepository
}

func BindProcesosElectoraisController(
	e *gin.Engine,
	procesosElectoraisRepository domain.ProcesosElectoraisRepository,
) {
	c := &ProcesosElectoraisController{}
	c.repository = procesosElectoraisRepository
	e.GET("/procesos-electorais", c.GetProcesosElectorais)
	e.GET("/procesos-electorais/tipos", c.GetProcesosElectoraisTipos)
	e.GET("/procesos-electorais/:id", c.GetProcesoElectoral)
}

func (c ProcesosElectoraisController) GetProcesosElectorais(gc *gin.Context) {
	procesosElectorais := c.repository.FindAll()
	gc.JSON(200, procesosElectorais)
}

func (c ProcesosElectoraisController) GetProcesoElectoral(gc *gin.Context) {
	var uriParams struct {
		Id int `uri:"id"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	procesoElectoral, ok := c.repository.FindById(uriParams.Id)

	if ok {
		gc.JSON(200, procesoElectoral)
	} else {
		gc.Status(404)
	}
}

func (c ProcesosElectoraisController) GetProcesosElectoraisTipos(gc *gin.Context) {
	procesosElectoraisTipos := c.repository.FindAllTipos()
	gc.JSON(200, procesosElectoraisTipos)
}
