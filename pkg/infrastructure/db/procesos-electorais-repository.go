package db

import (
	"database/sql"
	"errors"
	"github.com/earelin/pega/pkg/domain"
	"log"
	"time"
)

type ProcesosElectoraisSqlRepository struct {
	pool *sql.DB
}

func NewProcesosElectoraisSqlRepository(pool *sql.DB) *ProcesosElectoraisSqlRepository {
	return &ProcesosElectoraisSqlRepository{pool: pool}
}

func (r *ProcesosElectoraisSqlRepository) FindAll() []domain.ProcesoElectoral {
	var procesos []domain.ProcesoElectoral
	rows, err := r.pool.Query("SELECT id, data, tipo, ambito FROM proceso_electoral ORDER BY data DESC")
	if err != nil {
		log.Printf("Error querying procesos: %s", err)
	}

	for rows.Next() {
		var proceso domain.ProcesoElectoral
		var dataRaw string
		var ambitoRaw sql.NullInt16
		err = rows.Scan(&proceso.Id, &dataRaw, &proceso.Tipo, &ambitoRaw)
		if err != nil {
			log.Printf("Error scanning procesos: %s", err)
		}
		proceso.Data, err = time.Parse("2006-01-02 15:04:05", dataRaw)
		if err != nil {
			log.Printf("Error parsing date: %s. %s", dataRaw, err)
		}
		proceso.Ambito = int(ambitoRaw.Int16)
		procesos = append(procesos, proceso)
	}

	return procesos
}

func (r *ProcesosElectoraisSqlRepository) FindDatosXeraisProcesoById(id int) (domain.DatosXerais, bool) {
	return r.findDatosXerais("SELECT censo_ine, censo_cera FROM datos_xerais WHERE id = ?", id)
}

func (r *ProcesosElectoraisSqlRepository) FindDatosXeraisByComunidadeAutonoma(id int, comunidadeAutonomaId int) (domain.DatosXerais, bool) {
	return r.findDatosXerais("SELECT censo_ine, censo_cera FROM datos_xerais_autonomicos WHERE id = ? AND comunidade_autonoma_id = ?", id, comunidadeAutonomaId)
}

func (r *ProcesosElectoraisSqlRepository) FindDatosXeraisByProvincia(id int, provinciaId int) (domain.DatosXerais, bool) {
	return r.findDatosXerais("SELECT censo_ine, censo_cera FROM datos_xerais_provincias WHERE id = ? AND provincia_id = ?", id, provinciaId)
}

func (r *ProcesosElectoraisSqlRepository) FindDatosXeraisByConcello(id int, concelloId int) (domain.DatosXerais, bool) {
	return r.findDatosXerais("SELECT censo_ine, 0 FROM datos_xerais_concellos WHERE id = ? AND concello_id = ?", id, concelloId)
}

func (r *ProcesosElectoraisSqlRepository) findDatosXerais(query string, args ...any) (domain.DatosXerais, bool) {
	var datosXerais domain.DatosXerais
	row := r.pool.QueryRow(query, args...)

	var err = row.Scan(&datosXerais.CensoIne, &datosXerais.CensoCera)
	if errors.Is(err, sql.ErrNoRows) {
		return datosXerais, false
	} else if err != nil {
		log.Printf("Error scanning procesos: %s", err)
	}

	return datosXerais, true
}
