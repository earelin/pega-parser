package domain

type EntidadeAdministrativa struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

type EntidadesAdministrativasRepository interface {
	FindAllComunidadesAutonomas() []EntidadeAdministrativa
}
