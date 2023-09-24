package domain

type EntidadeAdministrativa struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

type EntidadesAdministrativasRepository interface {
	FindAllComunidadesAutonomas() []EntidadeAdministrativa
	FindAllProvincias() []EntidadeAdministrativa
	FindAllProvinciasByComunidadeAutonoma(caId int) []EntidadeAdministrativa
	FindAllConcellosByProvincia(pId int) []EntidadeAdministrativa
	FindAllConcellosByName(name string) []EntidadeAdministrativa
}
