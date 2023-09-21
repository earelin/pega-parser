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
		fmt.Println("Exporting candidatures")
		var c = e.Candidatures()
		err = fileExport(c, "candidatures")
		if err != nil {
			return fmt.Errorf("could not serialize candidatures: %w", err)
		}
	}

	if e.files.CandidatesListFile != "" {
		fmt.Println("Exporting candidates")
		var cl = e.CandidatesList()
		err = fileExport(cl, "candidates")
		if err != nil {
			return fmt.Errorf("could not serialize candidatures: %w", err)
		}
	}

	return nil
}

func (e Election) Candidatures() []Candidatura {
	fr := getFileReader[file_reader.CandidatureLine](e.zipFile, e.files.CandidaturesFile)
	defer fr.Close()

	var candidatures []Candidatura
	for {
		c, err := fr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic("Error reading candidatures files", err)
		}

		candidatures = append(candidatures, Candidatura{
			Codigo:              c.Code,
			Siglas:              c.Acronym,
			Nome:                c.Name,
			CabeceiraEstatal:    c.StateCode,
			CabeceiraAutonomica: c.AutonomicCode,
			CabeceiraProvincial: c.ProvincialCode,
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

		var ambitoTerritorial int
		if e.Type == Municipais {
			ambitoTerritorial = c.MunicipalCode
		} else {
			ambitoTerritorial = c.ProvinceCode
		}

		candidates = append(candidates, Candidate{
			AmbitoTerritorial: ambitoTerritorial,
			CodigoCandidatura: c.CandidatureCode,
			Posicion:          c.Position,
			Titular:           c.Type == "T",
			Nome:              c.Name,
			Apelidos:          strings.TrimSpace(fmt.Sprintf("%s %s", c.FirstSurname, c.SecondSurname)),
			FoiEleito:         c.Elected == "",
		})
	}

	return candidates
}

func (e Election) MesasElectorais() []MesaElectoral {
	fr := getFileReader[file_reader.DatosComunsDeMesasECera](e.zipFile, e.files.TablesAndCeraCommonDataFile)
	defer fr.Close()

	var mesas []MesaElectoral
	for {
		m, err := fr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic("Erro lendo as mesas electorais", err)
		}

		mesas = append(mesas, MesaElectoral{
			CodigoProvincia:   m.CodigoProvincia,
			CodigoConcello:    m.CodigoMunicipio,
			Distrito:          m.NumeroDistritoMunicipal,
			Seccion:           m.Seccion,
			CodigoMesa:        m.Mesa,
			CensoIne:          m.CensoIne,
			VotosBlanco:       m.VotosBlanco,
			VotosNulos:        m.VotosNulos,
			VotosCandidaturas: m.VotosACandidaturas,
		})
	}

	return mesas
}

func (e Election) VotosMesasElectorais() []VotosMesaElectoral {
	fr := getFileReader[file_reader.DatoCandidaturasDeMesasECera](e.zipFile, e.files.TablesAndCeraCandidaturesDataFile)
	defer fr.Close()

	var votosMesas []VotosMesaElectoral
	for {
		v, err := fr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic("Erro lendo os votos en mesas electorais", err)
		}

		votosMesas = append(votosMesas, VotosMesaElectoral{
			CodigoProvincia:      v.CodigoProvincia,
			CodigoConcello:       v.CodigoMunicipio,
			Distrito:             v.NumeroDistritoMunicipal,
			Seccion:              v.Seccion,
			CodigoMesa:           v.Mesa,
			CandidaturaOuSenador: v.CodigoCandidatura,
			Votos:                v.Votos,
		})
	}

	return votosMesas
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
