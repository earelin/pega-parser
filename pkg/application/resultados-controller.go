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
