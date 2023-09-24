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
	const query = "SELECT id, nome FROM comunidade_autonoma"
	var comunidades []domain.EntidadeAdministrativa
	rows, err := r.pool.Query(query)
	if err != nil {
		log.Printf("Error querying comunidades autonomas: %s", err)
	}

	for rows.Next() {
		var comunidade domain.EntidadeAdministrativa
		err = rows.Scan(&comunidade.Id, &comunidade.Nome)
		if err != nil {
			log.Printf("Error scanning comunidades autonomas: %s", err)
		}
		comunidades = append(comunidades, comunidade)
	}

	return comunidades
}
