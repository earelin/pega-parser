package importer

import (
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/election"
	"github.com/earelin/pega/tools/infoelectoral/pkg/repository"
)

func ImportElectionData(r *repository.Repository, e election.Election) error {
	var err error
	var procesoElectoralId int64

	procesoElectoralId, err = r.CreateProcesoElectoral(e)
	if err != nil {
		return fmt.Errorf("no se puido gardar o proceso electoral: %w", err)
	}

	candidatures := e.Candidatures()
	var importedCandidatures map[int]int64
	importedCandidatures, err = r.CreateCandidaturas(procesoElectoralId, candidatures)
	if err != nil {
		return fmt.Errorf("non se puideron gardar as candidaturas: %w", err)
	}

	listaCandidatos := e.CandidatesList()
	err = r.CrearListasECandidatos(listaCandidatos, importedCandidatures)
	if err != nil {
		return fmt.Errorf("non se puideron gardar as listas e candidatos: %w", err)
	}

	mesas := e.MesasElectorais()
	err = r.CrearMesasElectorais(procesoElectoralId, mesas)
	if err != nil {
		return fmt.Errorf("non se puideron crear as mesas electorais: %w", err)
	}

	return nil
}
