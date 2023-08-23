package election

const (
	EstateScope              = 0
	AutonomousCommunityScope = 1
)

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
	Referendum            = 1
	Congress              = 2
	Senate                = 3
	Municipals            = 4
	AutonomousCommunities = 5
	Cabildos              = 6
	EuropeanParliament    = 7
	JuntasGenerales       = 15
)

const (
	TitularCandidate   = 1
	AlternateCandidate = 2
)

var ElectionTypeLabel = map[int]string{
	Referendum:            "Referéndum",
	Congress:              "Congreso",
	Senate:                "Senado",
	Municipals:            "Municipais",
	AutonomousCommunities: "Comunidades Autónomas",
	Cabildos:              "Cabildos",
	EuropeanParliament:    "Parlamento Europeo",
	JuntasGenerales:       "Xuntas Xeráis",
}

type Candidature struct {
	Code    int    `json:"id"`
	Acronym string `json:"acronym"`
	Name    string `json:"name"`
}
