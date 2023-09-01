package file_reader

import (
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"testing"
)

func TestControlRead(t *testing.T) {
	var f fs.File
	f, _ = os.Open("testdata/control.DAT")
	fr, _ := NewFileReader[ControlLine](f)

	control := ControlLine{
		ElectionType:                       2,
		Year:                               2019,
		Month:                              11,
		Round:                              1,
		ControlFile:                        true,
		IdentificationFile:                 true,
		CandidaturesFile:                   true,
		CandidatesListFile:                 true,
		MunicipalitiesCommonDataFile:       true,
		MunicipalitiesCandidaturesDataFile: true,
		MunicipalitiesSuperiorScopeCommonDataFile:        true,
		MunicipalitiesSuperiorScopeCandidaturesDataFile:  true,
		TablesAndCeraCommonDataFile:                      true,
		TablesAndCeraCandidaturesDataFile:                true,
		MunicipalitiesSmallerThan250CommonDataFile:       false,
		MunicipalitiesSmallerThan250CandidaturesDataFile: false,
		JudicialDistrictCommonDataFile:                   false,
		JudicialDistrictCandidaturesDataFile:             false,
		ProvincialCouncilCommonDataFile:                  false,
		ProvincialCouncilCandidaturesDataFile:            false,
	}

	c, _ := fr.Read()
	assert.Equal(t, control, c)
}

func TestIdentificationRead(t *testing.T) {
	var f fs.File
	f, _ = os.Open("testdata/identification.DAT")
	fr, _ := NewFileReader[IdentificationLine](f)

	identification := IdentificationLine{
		Type:                           2,
		Year:                           2019,
		Month:                          11,
		Round:                          1,
		ScopeType:                      "N",
		TerritorialScope:               99,
		CelebrationDay:                 10,
		CelebrationMonth:               11,
		CelebrationYear:                2019,
		PollStationOpeningTime:         "09:00",
		PollStationClosingTime:         "20:00",
		FirstParticipationAdvanceTime:  "14:00",
		SecondParticipationAdvanceTime: "18:00",
	}

	i, _ := fr.Read()
	assert.Equal(t, identification, i)
}

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
