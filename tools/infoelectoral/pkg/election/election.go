package election

import (
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/archive_reader"
	"github.com/earelin/pega/tools/infoelectoral/pkg/file_reader"
	"io"
	"io/fs"
	"log"
	"strings"
	"time"
)

type dataFiles struct {
	IdentificationFile                               string
	CandidaturesFile                                 string
	CandidatesListFile                               string
	MunicipalitiesCommonDataFile                     string
	MunicipalitiesCandidaturesDataFile               string
	MunicipalitiesSuperiorScopeCommonDataFile        string
	MunicipalitiesSuperiorScopeCandidaturesDataFile  string
	TablesAndCeraCommonDataFile                      string
	TablesAndCeraCandidaturesDataFile                string
	MunicipalitiesSmallerThan250CommonDataFile       string
	MunicipalitiesSmallerThan250CandidaturesDataFile string
	JudicialDistrictCommonDataFile                   string
	JudicialDistrictCandidaturesDataFile             string
	ProvincialCouncilCommonDataFile                  string
	ProvincialCouncilCandidaturesDataFile            string
}

type Election struct {
	zipFile *archive_reader.ZipFile
	Type    int
	Scope   int
	Date    time.Time
	files   dataFiles
}

func (e Election) String() string {
	var date = e.Date.Format("02-01-2006")
	var electionType = ElectionTypeLabel[e.Type]
	return fmt.Sprintf("Election file for: %s %s\n", electionType, date)
}

func (e Election) ExportToFiles(fileExport func(interface{}, string) error) error {
	var err error

	if e.files.CandidaturesFile != "" {
		var c = e.Candidatures()
		err = fileExport(c, "candidatures")
		if err != nil {
			return fmt.Errorf("could not serialize candidatures: %w", err)
		}
	}

	if e.files.CandidatesListFile != "" {
		var cl = e.CandidatesList()
		err = fileExport(cl, "candidates")
		if err != nil {
			return fmt.Errorf("could not serialize candidatures: %w", err)
		}
	}

	return nil
}

func (e Election) Candidatures() []Candidature {
	fr := getFileReader[file_reader.CandidatureLine](e.zipFile, e.files.CandidaturesFile)
	defer fr.Close()

	var candidatures []Candidature
	for {
		c, err := fr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic("Error reading candidatures files", err)
		}

		candidatures = append(candidatures, Candidature{
			Code:    c.Code,
			Acronym: c.Acronym,
			Name:    c.Name,
		})
	}

	return candidatures
}

func (e Election) CandidatesList() []Candidate {
	fr := getFileReader[file_reader.CandidatesListLine](e.zipFile, e.files.CandidatesListFile)
	defer fr.Close()

	var candidates []Candidate
	for {
		c, err := fr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic("Error reading candidatures files", err)
		}

		var candidateType int
		if c.Type == "T" {
			candidateType = TitularCandidate
		} else {
			candidateType = AlternateCandidate
		}

		var birthdate time.Time
		if c.BirthdayDay != 0 && c.BirthdayMonth != 0 && c.BirthdayYear != 0 {
			birthdate, err = time.Parse("2006-01-02", fmt.Sprintf("%04d-%02d-%02d", c.BirthdayYear, c.BirthdayMonth, c.BirthdayDay))
			if err != nil {
				log.Panic("Error parsing date", err)
			}
		}

		var elected bool
		if c.Elected == "S" {
			elected = true
		} else {
			elected = false
		}

		candidates = append(candidates, Candidate{
			ProvinceCode:          c.ProvinceCode,
			ElectoralDistrictCode: c.ElectoralDistrict,
			MunicipalCode:         c.MunicipalCode,
			CandidatureCode:       c.CandidatureCode,
			Position:              c.Position,
			Type:                  candidateType,
			Name:                  c.Name,
			Surname:               strings.TrimSpace(fmt.Sprintf("%s %s", c.FirstSurname, c.SecondSurname)),
			Birthdate:             birthdate,
			NationalIdentityCard:  c.NationalIdentityCard,
			Elected:               elected,
		})
	}

	return candidates
}

func getFileReader[T any](archive *archive_reader.ZipFile, filename string) file_reader.FileReader[T] {
	var err error
	var file fs.File
	file, err = archive.Open(filename)
	if err != nil {
		log.Panic("Could not open file", filename, err)
	}

	var fileReader file_reader.FileReader[T]
	fileReader, err = file_reader.NewFileReader[T](file)
	if err != nil {
		log.Panic("Could not open reader for file", filename, err)
	}

	return fileReader
}
