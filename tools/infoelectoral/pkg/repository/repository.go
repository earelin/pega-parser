package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/election"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

const inserirCircunscripcionCera = "INSERT INTO circunscripcion_cera(proceso_electoral_id, provincia_id, censo, votos_blanco, votos_nulos, votos_candidaturas) VALUES "
const inserirLista = "INSERT INTO lista(candidatura_id, ambito) VALUES (? , ?)"
const inserirMesaElectoral = "INSERT INTO mesa_electoral(proceso_electoral_id, concello_id, distrito, seccion, codigo, censo, votos_blanco, votos_nulos, votos_candidaturas) VALUES "
const inserirVotosCircunscripcionCera = "INSERT INTO circunscripcion_cera_votos_candidatura(circuscripcion_cera_id, candidatura_id, posicion, votos) VALUES "
const inserirVotosMesaElectoral = "INSERT INTO mesa_electoral_votos_candidatura(mesa_electoral_id, candidatura_id, posicion, votos) VALUES "

type Repository struct {
	pool *sql.DB
	ctx  context.Context
}

func NewRepository(c Config, ctx context.Context) (*Repository, error) {
	var r Repository

	var pool, err = sql.Open("mysql", c.toString())
	if err != nil {
		return nil, err
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	r.pool = pool
	r.ctx = ctx

	return &r, nil
}

func (r *Repository) CheckConnection() error {
	var ctx, cancel = context.WithTimeout(r.ctx, 5*time.Second)
	defer cancel()
	return r.pool.PingContext(ctx)
}

func (r *Repository) CloseConnection() error {
	return r.pool.Close()
}

func (r *Repository) CreateProcesoElectoral(e election.Election) (int64, error) {
	const inserirProcesoElectoral = "INSERT INTO proceso_electoral(tipo, ambito, data) VALUES (?, ?, ?)"

	var result sql.Result
	var err error
	if e.Scope == 99 {
		result, err = r.pool.ExecContext(r.ctx, inserirProcesoElectoral, e.Type, nil, e.Date)
	} else {
		result, err = r.pool.ExecContext(r.ctx, inserirProcesoElectoral, e.Type, e.Scope, e.Date)
	}

	if err != nil {
		return 0, fmt.Errorf("no ha sido posible obtener el id de una candidatura guardada: %w", err)
	}

	var id int64
	id, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("no ha sido posible obtener el id de una candidatura guardada: %w", err)
	}

	return id, err
}

func (r *Repository) CreateCandidaturas(procesoElectoral int64, candidatures []election.Candidatura) (map[int]int64, error) {
	const inserirCandidatura = "INSERT INTO candidatura(proceso_electoral_id, siglas, nome) VALUES (?, ?, ?)"
	const actualizarConCabecerias = "UPDATE candidatura SET cabeceira_estatal = ?, cabeceira_autonomica = ?, cabeceira_provincial = ? WHERE id = ?"
	var importedItems = make(map[int]int64)

	for _, c := range candidatures {
		var result, err = r.pool.ExecContext(r.ctx, inserirCandidatura,
			procesoElectoral, c.Siglas, c.Nome)
		if err != nil {
			return nil, fmt.Errorf("non foi posible gardar unha candidatura: %w", err)
		}

		var id int64
		id, err = result.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("no foi posible obter o id dunha candidatura gardada: %w", err)
		}

		importedItems[c.Codigo] = id

		_, err = r.pool.ExecContext(r.ctx, actualizarConCabecerias,
			importedItems[c.CabeceiraEstatal],
			importedItems[c.CabeceiraAutonomica],
			importedItems[c.CabeceiraProvincial],
			id)
		if err != nil {
			return nil, fmt.Errorf("non foi posible actualizar unha candidatura coas cabeceiras: %w", err)
		}
	}

	return importedItems, nil
}

func (r *Repository) CrearListasECandidatos(listaCandidatos []election.Candidate, candidaturasImportadas map[int]int64) error {
	const inserirCandidato = "INSERT INTO candidato(lista_id, posicion, titular, nombre, apelidos) VALUES (?, ?, ?, ?, ?)"
	var listasImportadas = make(map[string]int64)

	for _, c := range listaCandidatos {
		var candidaturaId = candidaturasImportadas[c.CodigoCandidatura]
		codigoCandidaturaEAmbitoTerritorial := fmt.Sprintf("%d_%d", c.CodigoCandidatura, c.AmbitoTerritorial)
		var listaId, importada = listasImportadas[codigoCandidaturaEAmbitoTerritorial]
		if !importada {
			var result, err = r.pool.ExecContext(r.ctx, inserirLista, candidaturaId, c.AmbitoTerritorial)
			if err != nil {
				return fmt.Errorf("no foi posible gardar unha lista: %w", err)
			}
			listaId, err = result.LastInsertId()
			if err != nil {
				return fmt.Errorf("non foi posible obter o id dunha lista gardada: %w", err)
			}
			listasImportadas[codigoCandidaturaEAmbitoTerritorial] = listaId
		}

		var _, err = r.pool.ExecContext(r.ctx, inserirCandidato, listaId, c.Posicion, c.Titular, c.Nome, c.Apelidos)
		if err != nil {
			return fmt.Errorf("non foi posible gardar un candidato: %w", err)
		}
	}

	return nil
}

func (r *Repository) CrearMesasElectorais(procesoElectoral int64, mesas []election.MesaElectoral) (map[string]int64, error) {
	var err error
	var mesasImportadas = make(map[string]int64)

	var cInserts, mInserts, cHash, mHash [][]string
	var cValues, mValues [][]interface{}
	var j = -1

	for i, m := range mesas {
		if i%1000 == 0 {
			j++
			cInserts = append(cInserts, []string{})
			cHash = append(cHash, []string{})
			cValues = append(cValues, []interface{}{})
			mInserts = append(mInserts, []string{})
			mHash = append(mHash, []string{})
			mValues = append(mValues, []interface{}{})
		}

		hash := fmt.Sprintf("%d_%d_%d_%d_%s", m.CodigoProvincia, m.CodigoConcello, m.Distrito, m.Seccion, m.CodigoMesa)

		if m.CodigoProvincia == 99 {
			continue
		} else if m.CodigoConcello == 999 {
			cInserts[j] = append(cInserts[j], "(?, ?, ?, ?, ?, ?)")
			cValues[j] = append(cValues[j], procesoElectoral, m.CodigoProvincia,
				m.CensoIne, m.VotosBlanco, m.VotosNulos, m.VotosCandidaturas)
			cHash[j] = append(cHash[j], hash)
		} else {
			mInserts[j] = append(mInserts[j], "(?, ?, ?, ?, ?, ?, ?, ?, ?)")
			concelloId := m.CodigoProvincia*1000 + m.CodigoConcello
			mValues[j] = append(mValues[j], procesoElectoral, concelloId, m.Distrito,
				m.Seccion, m.CodigoMesa, m.CensoIne, m.VotosBlanco, m.VotosNulos, m.VotosCandidaturas)
			mHash[j] = append(mHash[j], hash)
		}
	}

	var sqlResult sql.Result

	for i := 0; i < len(cInserts); i++ {
		if len(cInserts[i]) == 0 {
			continue
		}

		sqlInsert := inserirCircunscripcionCera + strings.Join(cInserts[i], ",")
		sqlResult, err = r.pool.ExecContext(r.ctx, sqlInsert, cValues[i]...)
		if err != nil {
			return nil, fmt.Errorf("non se puideron gardar as circunscripcions CERA: %w", err)
		}

		var insertedId int64
		insertedId, err = sqlResult.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("non se puido obter o id da ultima circunscripcions CERA inserida: %w", err)
		}
		for j := len(cHash[i]) - 1; j >= 0; j-- {
			mesasImportadas[cHash[i][j]] = insertedId
			insertedId -= 1
		}
	}

	for i := 0; i < len(mInserts); i++ {
		if len(mInserts[i]) == 0 {
			continue
		}

		sqlInsert := inserirMesaElectoral + strings.Join(mInserts[i], ",")
		sqlResult, err = r.pool.ExecContext(r.ctx, sqlInsert, mValues[i]...)
		if err != nil {
			return nil, fmt.Errorf("non se puideron gardar as mesas electorais: %w", err)
		}

		var insertedId int64
		insertedId, err = sqlResult.LastInsertId()
		if err != nil {
			return nil, fmt.Errorf("non se puido obter o id da mesa electoral inserida: %w", err)
		}
		for j := len(mHash[i]) - 1; j >= 0; j-- {
			mesasImportadas[mHash[i][j]] = insertedId
			insertedId -= 1
		}
	}

	return mesasImportadas, nil
}

func (r *Repository) CrearVotosEnMesasElectorais(candidaturasImportadas map[int]int64, mesasImportadas map[string]int64, votos []election.VotosMesaElectoral) error {
	var err error

	var cInserts, mInserts [][]string
	var cValues, mValues [][]interface{}
	var j = -1

	for i, v := range votos {
		if i%1000 == 0 {
			j++
			cInserts = append(cInserts, []string{})
			cValues = append(cValues, []interface{}{})
			mInserts = append(mInserts, []string{})
			mValues = append(mValues, []interface{}{})
		}

		hashCircunscripcionOuMesa := fmt.Sprintf("%d_%d_%d_%d_%s", v.CodigoProvincia, v.CodigoConcello, v.Distrito, v.Seccion, v.CodigoMesa)
		circunscripcionOuMesa := mesasImportadas[hashCircunscripcionOuMesa]
		candidatura := candidaturasImportadas[v.CandidaturaOuSenador]

		if v.CodigoProvincia == 99 || v.Votos == 0 {
			continue
		} else if v.CodigoConcello == 999 {
			cInserts[j] = append(cInserts[j], "(?, ?, ?, ?)")
			cValues[j] = append(cValues[j], circunscripcionOuMesa, candidatura, nil, v.Votos)
		} else {
			mInserts[j] = append(mInserts[j], "(?, ?, ?, ?)")
			mValues[j] = append(mValues[j], circunscripcionOuMesa, candidatura, nil, v.Votos)
		}

		if err != nil {
			return fmt.Errorf("non se puideron insertar os votos dunha candidatura: %w", err)
		}
	}

	for i := 0; i < len(cInserts); i++ {
		if len(cInserts[i]) == 0 {
			continue
		}

		sqlInsert := inserirVotosCircunscripcionCera + strings.Join(cInserts[i], ",")
		_, err = r.pool.ExecContext(r.ctx, sqlInsert, cValues[i]...)
		if err != nil {
			return fmt.Errorf("non se puideron gardar os votos de circunscripcion CERA: %w", err)
		}
	}

	for i := 0; i < len(mInserts); i++ {
		if len(mInserts[i]) == 0 {
			continue
		}

		sqlInsert := inserirVotosMesaElectoral + strings.Join(mInserts[i], ",")
		_, err = r.pool.ExecContext(r.ctx, sqlInsert, mValues[i]...)
		if err != nil {
			return fmt.Errorf("non se puideron gardar os votos de mesa electoral: %w", err)
		}
	}

	return nil
}
