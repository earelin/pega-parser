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

func (c ProcesosElectoraisController) GetProcesosElectorais(gc *gin.Context) {
	procesosElectorais := c.repository.FindAll()
	gc.JSON(200, procesosElectorais)
}

func NewProcesosElectoraisController(
	e *gin.Engine,
	procesosElectoraisRepository domain.ProcesosElectoraisRepository,
) {
	c := &ProcesosElectoraisController{}
	c.repository = procesosElectoraisRepository
	e.GET("/procesos-electorais", c.GetProcesosElectorais)
}
