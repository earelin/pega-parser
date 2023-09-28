package domain

type Resultados struct {
	VotosBlanco       int `json:"votosBranco"`
	VotosNulos        int `json:"votosNulos"`
	VotosCandidaturas int `json:"votosCandidaturas"`
}

type ResultadoCandidatura struct {
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

type ResultadosCandidaturasRepository interface {
	FindByProceso(id int) ([]ResultadoCandidatura, bool)
	FindByComunidadeAutonoma(id int, comunidadeAutonomaId int) ([]ResultadoCandidatura, bool)
	FindByProvincia(id int, provinciaId int) ([]ResultadoCandidatura, bool)
	FindByConcello(id int, concelloId int) ([]ResultadoCandidatura, bool)
	FindByDistrito(id int, concelloId int, distritoId int) ([]ResultadoCandidatura, bool)
	FindBySeccion(id int, concelloId int, distritoId int, seccionId int) ([]ResultadoCandidatura, bool)
	FindByMesa(id int, concelloId int, distritoId int, seccionId int, codigoMesa string) ([]ResultadoCandidatura, bool)
}
