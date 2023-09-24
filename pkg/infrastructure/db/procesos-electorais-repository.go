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

func (r *ProcesosElectoraisSqlRepository) FindDatosXeraisById(id int) (domain.DatosXerais, bool) {
	var datosXerais domain.DatosXerais
	row := r.pool.QueryRow("SELECT id, data, tipo, ambito FROM proceso_electoral WHERE id = ?", id)

	var err = row.Scan()
	if errors.Is(err, sql.ErrNoRows) {
		return datosXerais, false
	} else if err != nil {
		log.Printf("Error scanning procesos: %s", err)
	}

	return datosXerais, true
}
