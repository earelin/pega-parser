package election

import (
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/archive_reader"
	"github.com/earelin/pega/tools/infoelectoral/pkg/file_reader"
	"io/fs"
	"log"
	"strconv"
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
	return fmt.Sprintf("Election file for: %s\n", date)
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

func buildFilenameGenerator(electionType int, month int, year int) func(bool, string) string {
	return func(exists bool, fileType string) string {
		var s string
		var yearString = strconv.Itoa(year)[2:]
		if exists {
			s = fmt.Sprintf("%s%02d%s%02d.DAT", fileType, electionType, yearString, month)
		}
		return s
	}
}

func buildCustomPrefixFilenameGenerator(month int, year int) func(bool, string) string {
	return func(exists bool, prefix string) string {
		var s string
		var yearString = strconv.Itoa(year)[2:]
		if exists {
			s = fmt.Sprintf("%s%s%02d.DAT", prefix, yearString, month)
		}
		return s
	}
}
