package domain

type Resultados struct {
	VotosBlanco       int
	VotosNulos        int
	VotosCandidaturas []VotosCandidatura
}

type VotosCandidatura struct {
	Candidatura Candidatura
	Votos       int
}

type ResultadosRepository interface {
	FindByProceso(id int) (Resultados, bool)
	FindByComunidadeAutonoma(id int, comunidadeAutonomaId int) (Resultados, bool)
	FindByProvincia(id int, provinciaId int) (Resultados, bool)
	FindByConcello(id int, concelloId int) (Resultados, bool)
	FindByDistrito(id int, concelloId int, distritoId int) (Resultados, bool)
	FindBySeccion(id int, concelloId int, distritoId int, seccionId int) (Resultados, bool)
	FindByMesa(id int, concelloId int, distritoId int, seccionId int, codigoMesa string) (Resultados, bool)
}
