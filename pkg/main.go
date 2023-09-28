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
	"github.com/earelin/pega/pkg/infrastructure/db"
	"github.com/gin-gonic/gin"
)

func ApplicationConfig(e *gin.Engine) {
	dbConfig := db.Config{
		Filename: "./database.sqlite",
	}
	var pool = dbConfig.BuildPool()
	var ear = db.NewEntidadesAdministrativasSqlRepository(pool)
	var per = db.NewProcesosElectoraisSqlRepository(pool)
	var dxr = db.NewDatosXeraisSqlRepository(pool)
	var rr = db.NewResultadosSqlRepository(pool)
	application.ConfigureApplicationLayer(e, ear, per, dxr, rr)
}
