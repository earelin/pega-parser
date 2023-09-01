package pkg

import "time"

type ProcesoElectoral struct {
	Id     int
	Tipo   int
	Data   time.Time
	Ambito int
}

type Organizacion struct {
	Id   int
	Nome string
}

type Candidatura struct {
	Ambito            int
	Organizacion      Organizacion
	PersoasCandidatas []PersoaCandidata
}

type PersoaCandidata struct {
	Nome     string
	Apelidos string
}
