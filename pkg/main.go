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

package pkg

import (
	"github.com/earelin/pega/pkg/application"
	"github.com/earelin/pega/pkg/repository"
	"github.com/gin-gonic/gin"
)

func ApplicationConfig(e *gin.Engine) {
	dbConfig := repository.Config{
		Filename: "./database.sqlite",
	}
	var pool = dbConfig.BuildPool()
	var ear = repository.NewDivisionsAdministrativasSqlRepository(pool)
	var per = repository.NewProcesosElectoraisSqlRepository(pool)
	application.ConfigureApplicationLayer(e, ear, per)
}
