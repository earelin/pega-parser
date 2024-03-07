/*
 * This program is free software: you can redistribute it and/or modify it under
 * the terms of the GNU General Public License as published by the Free Software
 * Foundation, either version 3 of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License
 * for more details.
 *
 * You should have received a copy of the GNU General Public License along with
 * this program. If not, see <https://www.gnu.org/licenses/>.
 */

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
	err = r.CreateCandidaturas(procesoElectoralId, candidatures)
	if err != nil {
		return fmt.Errorf("non se puideron gardar as candidaturas: %w", err)
	}

	listaCandidatos := e.CandidatesList()
	err = r.CrearListasECandidatos(procesoElectoralId, listaCandidatos)
	if err != nil {
		return fmt.Errorf("non se puideron gardar as listas e candidatos: %w", err)
	}

	mesas := e.MesasElectorais()
	err = r.CrearMesasElectorais(procesoElectoralId, mesas)
	if err != nil {
		return fmt.Errorf("non se puideron importar as mesas electorais: %w", err)
	}

	votosMesas := e.VotosMesasElectorais()
	err = r.CrearVotosEnMesasElectorais(procesoElectoralId, votosMesas)
	if err != nil {
		return fmt.Errorf("non se puideron importar os votos en mesa: %w", err)
	}

	return err
}
