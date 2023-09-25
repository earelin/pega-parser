package application

import (
	"github.com/earelin/pega/pkg/domain"
	"github.com/gin-gonic/gin"
)

type ResultadosController struct {
	repository domain.ResultadosRepository
}

func NewResultadosController(e *gin.Engine, repository domain.ResultadosRepository) {
	c := &ResultadosController{}
	c.repository = repository
}
