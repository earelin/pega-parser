package db

import (
	"database/sql"
	"github.com/earelin/pega/pkg/domain"
	"log"
)

type EntidadesAdministrativasSqlRepository struct {
	pool *sql.DB
}

func NewEntidadesAdministrativasSqlRepository(pool *sql.DB) *EntidadesAdministrativasSqlRepository {
	return &EntidadesAdministrativasSqlRepository{pool: pool}
}

func (r *EntidadesAdministrativasSqlRepository) FindAllComunidadesAutonomas() []domain.EntidadeAdministrativa {
	return r.findEntidades("SELECT id, nome FROM comunidade_autonoma ORDER BY nome")
}

func (r *EntidadesAdministrativasSqlRepository) FindAllProvincias() []domain.EntidadeAdministrativa {
	return r.findEntidades("SELECT id, nome FROM provincia ORDER BY nome")
}

func (r *EntidadesAdministrativasSqlRepository) FindAllProvinciasByComunidadeAutonoma(caId int) []domain.EntidadeAdministrativa {
	return r.findEntidades("SELECT id, nome FROM provincia WHERE comunidade_autonoma_id = ? ORDER BY nome", caId)
}

func (r *EntidadesAdministrativasSqlRepository) FindAllConcellosByProvincia(pId int) []domain.EntidadeAdministrativa {
	return r.findEntidades("SELECT id, nome FROM concello WHERE provincia_id = ? ORDER BY nome", pId)
}

func (r *EntidadesAdministrativasSqlRepository) FindAllConcellosByName(name string) []domain.EntidadeAdministrativa {
	log.Printf("Searching for concellos with name: %s", name)
	return r.findEntidades("SELECT id, nome FROM concello WHERE nome LIKE ? ORDER BY nome", "%"+name+"%")
}

func (r *EntidadesAdministrativasSqlRepository) findEntidades(sql string, args ...any) []domain.EntidadeAdministrativa {
	var entidades []domain.EntidadeAdministrativa
	rows, err := r.pool.Query(sql, args...)
	if err != nil {
		log.Printf("Error querying entidaes: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var provincia domain.EntidadeAdministrativa
		err = rows.Scan(&provincia.Id, &provincia.Nome)
		if err != nil {
			log.Printf("Error scanning entidades: %s", err)
		}
		entidades = append(entidades, provincia)
	}

	return entidades
}
