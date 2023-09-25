package domain

import "time"

type ProcesoElectoral struct {
	Id     int       `json:"id"`
	Data   time.Time `json:"data"`
	Tipo   int       `json:"tipo"`
	Ambito int       `json:"ambito"`
}

type DatosXerais struct {
	CensoIne                    int       `json:"censoIne"`
	CensoCera                   int       `json:"censoCera"`
	PrimeiroAvanceParticipacion time.Time `json:"primeiroAvanceParticipacion"`
	SegundoAvanceParticipacion  time.Time `json:"segundoAvanceParticipacion"`
}

type ProcesosElectoraisRepository interface {
	FindAll() []ProcesoElectoral
	FindDatosXeraisProcesoById(id int) (DatosXerais, bool)
	FindDatosXeraisByComunidadeAutonoma(id int, comunidadeAutonomaId int) (DatosXerais, bool)
	FindDatosXeraisByProvincia(id int, provinciaId int) (DatosXerais, bool)
	FindDatosXeraisByConcello(id int, concelloId int) (DatosXerais, bool)
	FindDatosXeraisByDistrito(id int, concelloId int, distritoId int) (DatosXerais, bool)
	FindDatosXeraisBySeccion(id int, concelloId int, distritoId int, seccionId int) (DatosXerais, bool)
	FindDatosXeraisByMesa(id int, concelloId int, distritoId int, seccionId int, codigoMesa string) (DatosXerais, bool)
}
