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
		Scope:   EstateScope,
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

func Test_buildFilenameGenerator(t *testing.T) {
	var filenameGenerator = buildFilenameGenerator(2, 3, 2023)

	assert.Equal(t, "03022303.DAT", filenameGenerator(true, "03"))
	assert.Equal(t, "", filenameGenerator(false, "03"))
}

func Test_buildCustomPrefixFilenameGenerator(t *testing.T) {
	var filenameGenerator = buildCustomPrefixFilenameGenerator(3, 2023)

	assert.Equal(t, "05012303.DAT", filenameGenerator(true, "0501"))
	assert.Equal(t, "", filenameGenerator(false, "0501"))
}
