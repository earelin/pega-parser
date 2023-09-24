package domain

import "time"

type ProcesoElectoral struct {
	Id     int       `json:"id"`
	Data   time.Time `json:"data"`
	Tipo   int       `json:"tipo"`
	Ambito int       `json:"ambito"`
}

type ProcesoElectoralDetails struct {
	Id     int       `json:"id"`
	Data   time.Time `json:"data"`
	Tipo   int       `json:"tipo"`
	Ambito int       `json:"ambito"`
}

type ProcesosElectoraisRepository interface {
	FindAll() []ProcesoElectoral
	FindById(id int) (ProcesoElectoralDetails, bool)
}
