package pkg

import (
	"github.com/earelin/pega/pkg/application"
	"github.com/earelin/pega/pkg/infrastructure/db"
	"github.com/gin-gonic/gin"
)

func ApplicationConfig(e *gin.Engine) {
	dbConfig := db.Config{
		Database: "pega",
		User:     "root",
	}
	var pool = dbConfig.BuildPool()
	var ear = db.NewEntidadesAdministrativasSqlRepository(pool)
	var per = db.NewProcesosElectoraisSqlRepository(pool)
	var dxr = db.NewDatosXeraisSqlRepository(pool)
	var rr = db.NewResultadosSqlRepository(pool)
	application.ConfigureApplicationLayer(e, ear, per, dxr, rr)
}
