package importer

import (
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/election"
	"github.com/earelin/pega/tools/infoelectoral/pkg/repository"
)

func ImportElectionData(r *repository.Repository, e election.Election) error {
	var err error

	err = r.CreateProcesoElectoral(e)
	if err != nil {
		return fmt.Errorf("no se puido gardar o proceso electora na base de datos: %w", err)
	}

	candidatures := e.Candidatures()
	//var importedCandidatures map[int]int64
	_, err = r.CreateCandidaturas(candidatures)
	if err != nil {
		return fmt.Errorf("non se puideron gardar as candidaturas: %w", err)
	}

	return nil
}
