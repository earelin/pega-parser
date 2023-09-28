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

package election

import (
	"github.com/earelin/pega/tools/infoelectoral/pkg/archive_reader"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewElection(t *testing.T) {
	zipFile, _ := archive_reader.NewZipFile("../../testdata/02201911_MESA.zip")
	electionDate, _ := time.Parse("2006-01-02", "2019-11-10")

	var e = NewElection(zipFile)

	assert.Equal(t, Election{
		zipFile: zipFile,
		Type:    2,
		Date:    electionDate,
		Scope:   99,
		files: dataFiles{
			IdentificationFile:                               "02021911.DAT",
			CandidaturesFile:                                 "03021911.DAT",
			CandidatesListFile:                               "04021911.DAT",
			MunicipalitiesCommonDataFile:                     "05021911.DAT",
			MunicipalitiesCandidaturesDataFile:               "06021911.DAT",
			MunicipalitiesSuperiorScopeCommonDataFile:        "07021911.DAT",
			MunicipalitiesSuperiorScopeCandidaturesDataFile:  "08021911.DAT",
			TablesAndCeraCommonDataFile:                      "09021911.DAT",
			TablesAndCeraCandidaturesDataFile:                "10021911.DAT",
			MunicipalitiesSmallerThan250CommonDataFile:       "",
			MunicipalitiesSmallerThan250CandidaturesDataFile: "",
			JudicialDistrictCommonDataFile:                   "",
			JudicialDistrictCandidaturesDataFile:             "",
			ProvincialCouncilCommonDataFile:                  "",
			ProvincialCouncilCandidaturesDataFile:            "",
		},
	}, e)
}
