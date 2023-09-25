package domain

import "time"

type DatosXerais struct {
	CensoIne                    int       `json:"censoIne"`
	CensoCera                   int       `json:"censoCera"`
	PrimeiroAvanceParticipacion time.Time `json:"primeiroAvanceParticipacion"`
	SegundoAvanceParticipacion  time.Time `json:"segundoAvanceParticipacion"`
}

type DatosXeraisRepository interface {
	FindByProceso(id int) (DatosXerais, bool)
	FindByComunidadeAutonoma(id int, comunidadeAutonomaId int) (DatosXerais, bool)
	FindByProvincia(id int, provinciaId int) (DatosXerais, bool)
	FindByConcello(id int, concelloId int) (DatosXerais, bool)
	FindByDistrito(id int, concelloId int, distritoId int) (DatosXerais, bool)
	FindBySeccion(id int, concelloId int, distritoId int, seccionId int) (DatosXerais, bool)
	FindByMesa(id int, concelloId int, distritoId int, seccionId int, codigoMesa string) (DatosXerais, bool)
}
