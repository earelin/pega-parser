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

type DivisionsAdministrativasAdminController struct {
	repository domain.DivisionsAdministrativasRepository
}

func BindDivisionsAdministrativasAdminController(e *gin.Engine, repository domain.DivisionsAdministrativasRepository) {
	c := DivisionsAdministrativasAdminController{repository: repository}
	e.POST("/v1/concellos/import", c.ImportConcellos)
}

func (c DivisionsAdministrativasAdminController) ImportConcellos(context *gin.Context) {

}
