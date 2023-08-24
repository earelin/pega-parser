package file_reader

import (
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"testing"
)

func TestCandidatureRead(t *testing.T) {
	var f fs.File
	f, _ = os.Open("testdata/candidatures.DAT")
	fr, _ := NewFileReader[CandidatureLine](f)

	candidatures := []CandidatureLine{
		{
			ElectionType:   2,
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
			ElectionType:   2,
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
			ElectionType:   2,
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

func TestCandidatesListRead(t *testing.T) {
	var f fs.File
	f, _ = os.Open("testdata/candidates.DAT")
	fr, _ := NewFileReader[CandidatesListLine](f)

	candidates := []CandidatesListLine{
		{
			ElectionType:         2,
			Year:                 2019,
			Month:                11,
			Round:                1,
			ProvinceCode:         7,
			ElectoralDistrict:    9,
			MunicipalCode:        999,
			CandidatureCode:      55,
			Position:             6,
			Type:                 "T",
			Name:                 "Emilio",
			FirstSurname:         "Hernández",
			SecondSurname:        "Martín",
			Genre:                "M",
			BirthdayDay:          0,
			BirthdayMonth:        0,
			BirthdayYear:         0,
			NationalIdentityCard: "0000000000",
			Elected:              "N",
		},
		{
			ElectionType:         2,
			Year:                 2019,
			Month:                11,
			Round:                1,
			ProvinceCode:         15,
			ElectoralDistrict:    9,
			MunicipalCode:        999,
			CandidatureCode:      91,
			Position:             6,
			Type:                 "T",
			Name:                 "Jesús",
			FirstSurname:         "Fernández",
			SecondSurname:        "Diez",
			Genre:                "M",
			BirthdayDay:          0,
			BirthdayMonth:        0,
			BirthdayYear:         0,
			NationalIdentityCard: "0000000000",
			Elected:              "N",
		},
		{
			ElectionType:         2,
			Year:                 2019,
			Month:                11,
			Round:                1,
			ProvinceCode:         41,
			ElectoralDistrict:    9,
			MunicipalCode:        999,
			CandidatureCode:      44,
			Position:             1,
			Type:                 "T",
			Name:                 "Esperanza",
			FirstSurname:         "Gómez",
			SecondSurname:        "Corona",
			Genre:                "F",
			BirthdayDay:          0,
			BirthdayMonth:        0,
			BirthdayYear:         0,
			NationalIdentityCard: "0000000000",
			Elected:              "S",
		},
	}

	for _, candidate := range candidates {
		c, err := fr.Read()
		assert.Nil(t, err)
		assert.Equal(t, candidate, c)
	}
}
