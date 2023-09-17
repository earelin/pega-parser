package election

const (
	FicheiroControlPrefixo                                 = "01"
	FicheiroIdentificacionPrefixo                          = "02"
	FicheiroCandidaturasPrefixo                            = "03"
	FicheiroListaCandidatosPrefix                          = "04"
	FicheiroDatosComunsConcellosPrefixo                    = "05"
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

type Candidatura struct {
	Codigo int
	Siglas string
	Nome   string
}

type Candidate struct {
	AmbitoTerritorial int
	CodigoCandidatura int
	Posicion          int
	Titular           bool
	Nome              string
	Apelidos          string
	FoiEleito         bool
}

type MesaElectoral struct {
	CodigoProvincia   int
	CodigoConcello    int
	Distrito          int
	Seccion           int
	CodigoMesa        string
	CensoIne          int
	VotosBlanco       int
	VotosNulos        int
	VotosCandidaturas int
}

type VotosMesaElectoral struct {
	CodigoProvincia      int
	CodigoConcello       int
	Distrito             int
	Seccion              int
	CodigoMesa           string
	CandidaturaOuSenador int
	Votos                int
}
