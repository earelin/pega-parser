package db

import (
	"database/sql"
	"github.com/earelin/pega/pkg/domain"
	"log"
)

func NewResultadosSqlRepository(pool *sql.DB) *ResultadosSqlRepository {
	return &ResultadosSqlRepository{pool: pool}
}

type ResultadosSqlRepository struct {
	pool *sql.DB
}

func (r *ResultadosSqlRepository) FindByProceso(id int) (domain.Resultados, bool) {
	return r.find(`
SELECT
    FROM `, id)
}

func (r *ResultadosSqlRepository) FindByComunidadeAutonoma(id int, comunidadeAutonomaId int) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) FindByProvincia(id int, provinciaId int) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) FindByConcello(id int, concelloId int) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) FindByDistrito(id int, concelloId int, distritoId int) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) FindBySeccion(id int, concelloId int, distritoId int, seccionId int) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) FindByMesa(id int, concelloId int, distritoId int, seccionId int, codigoMesa string) (domain.Resultados, bool) {
	return domain.Resultados{}, false
}

func (r *ResultadosSqlRepository) find(query string, args ...any) (domain.Resultados, bool) {
	var resultados domain.Resultados

	rows, err := r.pool.Query(query, args...)
	if err != nil {
		log.Printf("Error querying resultados: %s", err)
	}
	defer rows.Close()

	var firstResult = true
	for rows.Next() {
		var votosBranco, votosNulos int
		var vc = domain.VotosCandidatura{}

		err = rows.Scan(votosBranco, votosNulos, &vc.Candidatura.Id, &vc.Candidatura.Name, &vc.Votos)
		if err != nil {
			log.Printf("Error scanning resultados: %s", err)
		}

		if firstResult {
			resultados.VotosBlanco = votosBranco
			resultados.VotosNulos = votosNulos
			firstResult = false
		}

		resultados.VotosCandidaturas = append(resultados.VotosCandidaturas, vc)
	}

	return resultados, true
}
