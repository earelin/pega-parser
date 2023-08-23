package file_reader

import (
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"testing"
)

func TestCandidacyRead(t *testing.T) {
	var f fs.File
	f, _ = os.Open("testdata/candidatures.DAT")
	fr, _ := NewFileReader[CandidatureLine](f)

	candidatures := []CandidatureLine{
		{
			Type:           2,
			Year:           2019,
			Month:          11,
			Code:           2,
			Acronym:        "AHORA CANARIAS",
			Name:           "AHORA CANARIAS: Alternativa Nacionalista Canaria (ANC) y Unidad del Pueblo",
			ProvincialCode: 2,
			AutonomicCode:  2,
			StateCode:      2,
		},
		{
			Type:           2,
			Year:           2019,
			Month:          11,
			Code:           3,
			Acronym:        "ANDECHA",
			Name:           "ANDECHA ASTUR",
			ProvincialCode: 3,
			AutonomicCode:  3,
			StateCode:      3,
		},
		{
			Type:           2,
			Year:           2019,
			Month:          11,
			Code:           5,
			Acronym:        "AUNACV",
			Name:           "AUNA COMUNITAT VALENCIANA",
			ProvincialCode: 5,
			AutonomicCode:  5,
			StateCode:      5,
		},
	}

	for _, candidature := range candidatures {
		c, err := fr.Read()
		assert.Nil(t, err)
		assert.Equal(t, candidature, c)
	}
}
