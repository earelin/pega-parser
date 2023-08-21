package election

import (
	"github.com/earelin/pega/tools/infoelectoral/pkg/archive_reader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewElection(t *testing.T) {
	zipFile, _ := archive_reader.NewZipFile("../../testdata/02201911_MESA.zip")

	var e = NewElection(zipFile)

	assert.Equal(t, Election{
		zipFile: zipFile,
		Type:    2,
		Scope:   EstateScope,
		files: dataFiles{
			CandidaturesFile:                          "03021911.DAT",
			CandidatesListFile:                        "04021911.DAT",
			MunicipalitiesCommonDataFile:              "05021911.DAT",
			MunicipalitiesCandidaturesDataFile:        "06021911.DAT",
			MunicipalitiesSuperiorScopeCommonDataFile: "07021911.DAT",
		},
	}, e)
}

func Test_buildFilenameGenerator(t *testing.T) {
	var filenameGenerator = buildFilenameGenerator(2, 3, 2023)

	assert.Equal(t, "03022303.DAT", filenameGenerator(true, "03"))
	assert.Equal(t, "", filenameGenerator(false, "03"))
}
