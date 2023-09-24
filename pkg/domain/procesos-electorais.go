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
}
