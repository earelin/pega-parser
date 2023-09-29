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
	"fmt"
	"github.com/earelin/pega/tools/infoelectoral/pkg/archive_reader"
	"github.com/earelin/pega/tools/infoelectoral/pkg/file_reader"
	"io"
	"io/fs"
	"log"
	"strconv"
	"time"
)

func NewElection(zipFile *archive_reader.ZipFile) Election {
	var e = Election{
		zipFile: zipFile,
	}

	loadControlData(&e, zipFile)
	loadIdentificationData(&e, zipFile)

	return e
}

func loadControlData(e *Election, archive *archive_reader.ZipFile) {
	var err error
	var controlFile fs.File
	controlFile, err = archive.FindFileWithPrefix(FicheiroControlPrefixo)
	if err != nil {
		log.Panic("Could not find control file", err)
	}

	var controlFileReader file_reader.FileReader[file_reader.ControlLine]
	controlFileReader, err = file_reader.NewFileReader[file_reader.ControlLine](controlFile)
	if err != nil {
		log.Panic("Could not open reader for control file", err)
	}

	var control file_reader.ControlLine
	control, err = controlFileReader.Read()
	if err != nil && err != io.EOF {
		log.Panic("Could not read control file information", err)
	}

	e.Type = control.ElectionType
	var generateFileName = buildFilenameGenerator(control.ElectionType, control.Month, control.Year)
	var generateCustomPrefixFileName = buildCustomPrefixFilenameGenerator(
		control.Month, control.Year)

	e.files = dataFiles{
		IdentificationFile:                               generateFileName(true, FicheiroIdentificacionPrefixo),
		CandidaturesFile:                                 generateFileName(control.CandidaturesFile, FicheiroCandidaturasPrefixo),
		CandidatesListFile:                               generateFileName(control.CandidatesListFile, FicheiroListaCandidatosPrefix),
		MunicipalitiesCommonDataFile:                     generateFileName(control.MunicipalitiesCommonDataFile, FicheiroDatosComunsConcellosPrefixo),
		MunicipalitiesCandidaturesDataFile:               generateFileName(control.MunicipalitiesCandidaturesDataFile, MunicipalitiesCandidaturesDataFilePrefix),
		MunicipalitiesSuperiorScopeCommonDataFile:        generateFileName(control.MunicipalitiesSuperiorScopeCommonDataFile, MunicipalitiesSuperiorScopeCommonDataFilePrefix),
		MunicipalitiesSuperiorScopeCandidaturesDataFile:  generateFileName(control.MunicipalitiesSuperiorScopeCandidaturesDataFile, MunicipalitiesSuperiorScopeCandidaturesDataFile),
		TablesAndCeraCommonDataFile:                      generateFileName(control.TablesAndCeraCommonDataFile, TablesAndCeraCommonDataFilePrefix),
		TablesAndCeraCandidaturesDataFile:                generateFileName(control.TablesAndCeraCandidaturesDataFile, TablesAndCeraCandidaturesDataFilePrefix),
		MunicipalitiesSmallerThan250CommonDataFile:       generateFileName(control.MunicipalitiesSmallerThan250CommonDataFile, MunicipalitiesSmallerThan250CommonDataFilePrefix),
		MunicipalitiesSmallerThan250CandidaturesDataFile: generateFileName(control.MunicipalitiesSmallerThan250CandidaturesDataFile, MunicipalitiesSmallerThan250CommonDataFilePrefix),
		JudicialDistrictCommonDataFile:                   generateCustomPrefixFileName(control.JudicialDistrictCommonDataFile, JudicialDistrictCommonDataFilePrefix),
		JudicialDistrictCandidaturesDataFile:             generateCustomPrefixFileName(control.JudicialDistrictCandidaturesDataFile, JudicialDistrictCandidaturesDataFilePrefix),
		ProvincialCouncilCommonDataFile:                  generateCustomPrefixFileName(control.ProvincialCouncilCommonDataFile, ProvincialCouncilCommonDataFilePrefix),
		ProvincialCouncilCandidaturesDataFile:            generateCustomPrefixFileName(control.ProvincialCouncilCandidaturesDataFile, ProvincialCouncilCandidaturesDataFilePrefix),
	}
}

func loadIdentificationData(e *Election, archive *archive_reader.ZipFile) {
	var identificationFileReader = getFileReader[file_reader.IdentificationLine](
		archive, e.files.IdentificationFile,
	)

	var identification file_reader.IdentificationLine
	var err error
	identification, err = identificationFileReader.Read()
	if err != nil && err != io.EOF {
		log.Panic("Could not read control file information", err)
	}

	var electionDate = fmt.Sprintf("%04d-%02d-%02d",
		identification.CelebrationYear, identification.CelebrationMonth, identification.CelebrationDay)
	e.Date, err = time.Parse("2006-01-02", electionDate)
	if err != nil {
		log.Panic("Could not read election date", err)
	}
	e.Scope = identification.TerritorialScope
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
