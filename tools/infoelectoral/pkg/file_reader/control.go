package file_reader

const (
	Referendum        = 1
	Congreso          = 2
	Senado            = 3
	Municipales       = 4
	Autonomicas       = 5
	Cabildos          = 6
	ParlamentoEuropeo = 7
	JuntasGenerales   = 15
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

type Control struct {
	ElectionType                                     int  `position:"0" length:"2"`
	Year                                             int  `position:"2" length:"4"`
	Month                                            int  `position:"6" length:"2"`
	Round                                            int  `position:"8" length:"1"`
	ControlFile                                      bool `position:"9" length:"1"`
	IdentificationFile                               bool `position:"10" length:"1"`
	CandidaturesFile                                 bool `position:"11" length:"1"`
	CandidatesListFile                               bool `position:"12" length:"1"`
	MunicipalitiesCommonDataFile                     bool `position:"13" length:"1"`
	MunicipalitiesCandidaturesDataFile               bool `position:"14" length:"1"`
	MunicipalitiesSuperiorScopeCommonDataFile        bool `position:"15" length:"1"`
	MunicipalitiesSuperiorScopeCandidaturesDataFile  bool `position:"16" length:"1"`
	TablesAndCeraCommonDataFile                      bool `position:"17" length:"1"`
	TablesAndCeraCandidaturesDataFile                bool `position:"18" length:"1"`
	MunicipalitiesSmallerThan250CommonDataFile       bool `position:"19" length:"1"`
	MunicipalitiesSmallerThan250CandidaturesDataFile bool `position:"20" length:"1"`
	JudicialDistrictCommonDataFile                   bool `position:"21" length:"1"`
	JudicialDistrictCandidaturesDataFile             bool `position:"22" length:"1"`
	ProvincialCouncilCommonDataFile                  bool `position:"23" length:"1"`
	ProvincialCouncilCandidaturesDataFile            bool `position:"24" length:"1"`
}
