package election

import (
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/archive_reader"
	"github.com/earelin/pega/tools/infoelectoral/pkg/file_reader"
	"io"
	"io/fs"
	"log"
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

func (e Election) Candidatures() []Candidature {
	fr := getFileReader[file_reader.CandidacyLine](e.zipFile, e.files.CandidaturesFile)

	var candiatures []Candidature
	for {
		c, err := fr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic("Error reading candidatures files", err)
		}

		candiatures = append(candiatures, Candidature{
			Code:    c.Code,
			Acronym: c.Acronym,
			Name:    c.Name,
		})
	}

	return candiatures
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
