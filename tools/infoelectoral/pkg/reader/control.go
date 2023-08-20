package reader

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

type Control struct {
	ElectionType                                     int
	Year                                             int
	Month                                            int
	Round                                            int
	ControlFile                                      bool
	ElectoralProcessIdentificationFile               bool
	CandidaturesFile                                 bool
	ListOfCandidatesFile                             bool
	CommonMunicipalitiesDataFile                     bool
	MunicipalitiesCandidaturesFile                   bool
	DataOfScopeSuperiorToMunicipalityFile            bool
	CandidaturesScopeSuperiorToMunicipalityFile      bool
	CommonDataTablesAndCeraFile                      bool
	CandidaturesDataTablesAndCeraFile                bool
	CommonDataMunicipalitiesSmallerThan250File       bool
	CandidaturesDataMunicipalitiesSmallerThan250File bool
}
