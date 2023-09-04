package election

const (
	ControlFilePrefix                                      = "01"
	IdentificationFilePrefix                               = "02"
	CandidaturesFilePrefix                                 = "03"
	CandidatesListFilePrefix                               = "04"
	MunicipalitiesCommonDataFilePrefix                     = "05"
	MunicipalitiesCandidaturesDataFilePrefix               = "06"
	MunicipalitiesSuperiorScopeCommonDataFilePrefix        = "07"
	MunicipalitiesSuperiorScopeCandidaturesDataFile        = "08"
	TablesAndCeraCommonDataFilePrefix                      = "09"
	TablesAndCeraCandidaturesDataFilePrefix                = "10"
	MunicipalitiesSmallerThan250CommonDataFilePrefix       = "1104"
	MunicipalitiesSmallerThan250CandidaturesDataFilePrefix = "1204"
	JudicialDistrictCommonDataFilePrefix                   = "0510"
	JudicialDistrictCandidaturesDataFilePrefix             = "0610"
	ProvincialCouncilCommonDataFilePrefix                  = "0710"
	ProvincialCouncilCandidaturesDataFilePrefix            = "0810"
)

const (
	Referendum           = 1
	Congreso             = 2
	Senado               = 3
	Municipais           = 4
	ComunidadesAutonomas = 5
	Cabildos             = 6
	ParlamentoEuropeo    = 7
	XuntasXerais         = 15
)

var ElectionTypeLabel = map[int]string{
	Referendum:           "Referéndum",
	Congreso:             "Congreso",
	Senado:               "Senado",
	Municipais:           "Municipais",
	ComunidadesAutonomas: "Comunidades Autónomas",
	Cabildos:             "Cabildos",
	ParlamentoEuropeo:    "Parlamento Europeo",
	XuntasXerais:         "Xuntas Xeráis",
}

type Candidature struct {
	Code    int
	Acronym string
	Name    string
}

type Candidate struct {
	AmbitoTerritorial int
	CandidatureCode   int
	Position          int
	Titular           bool
	Name              string
	Surname           string
}
