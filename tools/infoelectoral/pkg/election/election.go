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
	return "Election file for: "
}

func NewElection(zipFile *archive_reader.ZipFile) Election {
	var e = Election{
		zipFile: zipFile,
	}

	loadControlData(&e, zipFile)

	return e
}

func loadControlData(e *Election, archive *archive_reader.ZipFile) {
	var err error
	var controlFile fs.File
	controlFile, err = archive.FindFileWithPrefix("01")
	if err != nil {
		log.Panic("Could not find control file", err)
	}

	var controlFileReader file_reader.FileReader[file_reader.Control]
	controlFileReader, err = file_reader.NewFileReader[file_reader.Control](controlFile)
	if err != nil {
		log.Panic("Could not open reader for control file", err)
	}

	var control file_reader.Control
	control, err = controlFileReader.Read()
	if err != nil {
		log.Panic("Could not read control file information", err)
	}

	e.Type = control.ElectionType
	var generateFileName = buildFilenameGenerator(control.ElectionType, control.Month, control.Year)

	e.files = dataFiles{
		CandidaturesFile:                          generateFileName(control.CandidaturesFile, CandidaturesFilePrefix),
		CandidatesListFile:                        generateFileName(control.CandidatesListFile, CandidatesListFilePrefix),
		MunicipalitiesCommonDataFile:              generateFileName(control.MunicipalitiesCommonDataFile, MunicipalitiesCommonDataFilePrefix),
		MunicipalitiesCandidaturesDataFile:        generateFileName(control.MunicipalitiesCandidaturesDataFile, MunicipalitiesCandidaturesDataFilePrefix),
		MunicipalitiesSuperiorScopeCommonDataFile: generateFileName(control.MunicipalitiesSuperiorScopeCommonDataFile, MunicipalitiesSuperiorScopeCommonDataFilePrefix),
		//MunicipalitiesSuperiorScopeCandidaturesDataFile
		//TablesAndCeraCommonDataFile
		//TablesAndCeraCandidaturesDataFile
		//MunicipalitiesSmallerThan250CommonDataFile
		//MunicipalitiesSmallerThan250CandidaturesDataFile
		//JudicialDistrictCommonDataFile
		//JudicialDistrictCandidaturesDataFile
		//ProvincialCouncilCommonDataFile
		//ProvincialCouncilCandidaturesDataFile
	}
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
