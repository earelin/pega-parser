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
	Id       int
	Nome     string
	Apelidos string
}

type MesaDeVotacion struct {
	Provincia         int
	Municipio         int
	DistritoMunicipal int
	Seccion           string
	Codigo            string
}
