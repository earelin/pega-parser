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
		return fmt.Errorf("no se puido gardar o proceso electora na base de datos: %w", err)
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

	return nil
}
