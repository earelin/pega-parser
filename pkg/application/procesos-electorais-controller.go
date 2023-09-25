package application

import (
	"github.com/earelin/pega/pkg/domain"
	"github.com/gin-gonic/gin"
)

type ProcesosElectoraisController struct {
	repository domain.ProcesosElectoraisRepository
}

type IdAndComunidadeAutonomaId struct {
	Id                   int `uri:"id"`
	ComunidadeAutonomaId int `uri:"comunidadeAutonomaId"`
}

func (c ProcesosElectoraisController) GetProcesosElectorais(gc *gin.Context) {
	procesosElectorais := c.repository.FindAll()
	gc.JSON(200, procesosElectorais)
}

func (c ProcesosElectoraisController) GetDatosXerais(gc *gin.Context) {
	var id Id
	if err := gc.ShouldBindUri(&id); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindDatosXeraisProcesoById(id.Id)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c ProcesosElectoraisController) GetDatosXeraisComunidadeAutonoma(gc *gin.Context) {
	var uriParams struct {
		Id                   int `uri:"id"`
		ComunidadeAutonomaId int `uri:"comunidadeAutonomaId"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindDatosXeraisByComunidadeAutonoma(uriParams.Id, uriParams.ComunidadeAutonomaId)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c ProcesosElectoraisController) GetDatosXeraisProvincia(gc *gin.Context) {
	var uriParams struct {
		Id          int `uri:"id"`
		ProvinciaId int `uri:"provinciaId"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindDatosXeraisByProvincia(uriParams.Id, uriParams.ProvinciaId)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c ProcesosElectoraisController) GetDatosXeraisConcello(gc *gin.Context) {
	var uriParams struct {
		Id         int `uri:"id"`
		ConcelloId int `uri:"concelloId"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindDatosXeraisByConcello(uriParams.Id, uriParams.ConcelloId)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c ProcesosElectoraisController) GetDatosXeraisDistrito(gc *gin.Context) {
	var uriParams struct {
		Id         int `uri:"id"`
		ConcelloId int `uri:"concelloId"`
		DistritoId int `uri:"distritoId"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindDatosXeraisByDistrito(uriParams.Id, uriParams.ConcelloId, uriParams.DistritoId)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c ProcesosElectoraisController) GetDatosXeraisSeccion(gc *gin.Context) {
	var uriParams struct {
		Id         int `uri:"id"`
		ConcelloId int `uri:"concelloId"`
		DistritoId int `uri:"distritoId"`
		SeccionId  int `uri:"seccionId"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindDatosXeraisBySeccion(uriParams.Id, uriParams.ConcelloId, uriParams.DistritoId, uriParams.SeccionId)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func (c ProcesosElectoraisController) GetDatosXeraisMesa(gc *gin.Context) {
	var uriParams struct {
		Id         int    `uri:"id"`
		ConcelloId int    `uri:"concelloId"`
		DistritoId int    `uri:"distritoId"`
		SeccionId  int    `uri:"seccionId"`
		CodigoMesa string `uri:"codigoMesa"`
	}
	if err := gc.ShouldBindUri(&uriParams); err != nil {
		gc.JSON(400, gin.H{"msg": err})
		return
	}

	ps, ok := c.repository.FindDatosXeraisByMesa(uriParams.Id, uriParams.ConcelloId, uriParams.DistritoId, uriParams.SeccionId, uriParams.CodigoMesa)

	if ok {
		gc.JSON(200, ps)
	} else {
		gc.Status(404)
	}
}

func NewProcesosElectoraisController(e *gin.Engine, procesosElectoraisRepository domain.ProcesosElectoraisRepository) {
	c := &ProcesosElectoraisController{}
	c.repository = procesosElectoraisRepository
	e.GET("/procesos-electorais", c.GetProcesosElectorais)
	e.GET("/proceso-electoral/:id/datos-xerais", c.GetDatosXerais)
	e.GET("/proceso-electoral/:id/datos-xerais/comunidade-autonoma/:comunidadeAutonomaId", c.GetDatosXeraisComunidadeAutonoma)
	e.GET("/proceso-electoral/:id/datos-xerais/provincia/:provinciaId", c.GetDatosXeraisProvincia)
	e.GET("/proceso-electoral/:id/datos-xerais/concello/:concelloId", c.GetDatosXeraisConcello)
	e.GET("/proceso-electoral/:id/datos-xerais/concello/:concelloId/:distritoId", c.GetDatosXeraisDistrito)
	e.GET("/proceso-electoral/:id/datos-xerais/concello/:concelloId/:distritoId/:seccionId", c.GetDatosXeraisSeccion)
	e.GET("/proceso-electoral/:id/datos-xerais/concello/:concelloId/:distritoId/:seccionId/:codigoMesa", c.GetDatosXeraisMesa)
}
