package file_reader

type ControlLine struct {
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

type IdentificationLine struct {
	Type                           int    `position:"0" length:"2"`
	Year                           int    `position:"2" length:"4"`
	Month                          int    `position:"6" length:"2"`
	Round                          int    `position:"8" length:"1"`
	ScopeType                      string `position:"9" length:"1"` // N/A
	TerritorialScope               int    `position:"10" length:"2"`
	CelebrationDay                 int    `position:"12" length:"2"`
	CelebrationMonth               int    `position:"14" length:"2"`
	CelebrationYear                int    `position:"16" length:"4"`
	PollStationOpeningTime         string `position:"20" length:"5"`
	PollStationClosingTime         string `position:"25" length:"5"`
	FirstParticipationAdvanceTime  string `position:"30" length:"5"`
	SecondParticipationAdvanceTime string `position:"35" length:"5"`
}

type CandidatureLine struct {
	ElectionType   int    `position:"0" length:"2"`
	Year           int    `position:"2" length:"4"`
	Month          int    `position:"6" length:"2"`
	Code           int    `position:"8" length:"6"`
	Acronym        string `position:"14" length:"50"`
	Name           string `position:"64" length:"150"`
	ProvincialCode int    `position:"214" length:"6"`
	AutonomicCode  int    `position:"220" length:"6"`
	StateCode      int    `position:"226" length:"6"`
}

type CandidatesListLine struct {
	ElectionType         int    `position:"0" length:"2"`
	Year                 int    `position:"2" length:"4"`
	Month                int    `position:"6" length:"2"`
	Round                int    `position:"8" length:"1"`
	ProvinceCode         int    `position:"9" length:"2"`
	ElectoralDistrict    int    `position:"11" length:"1"`
	MunicipalCode        int    `position:"12" length:"3"`
	CandidatureCode      int    `position:"15" length:"6"`
	Position             int    `position:"21" length:"3"`
	Type                 string `position:"24" length:"1"` // T/S
	Name                 string `position:"25" length:"25"`
	FirstSurname         string `position:"50" length:"25"`
	SecondSurname        string `position:"75" length:"25"`
	Genre                string `position:"100" length:"1"` // M/F
	BirthdayDay          int    `position:"101" length:"2"`
	BirthdayMonth        int    `position:"103" length:"2"`
	BirthdayYear         int    `position:"105" length:"4"`
	NationalIdentityCard string `position:"109" length:"10"`
	Elected              string `position:"119" length:"1"` // S/N
}

type MunicipalitiesCommonDataLine struct {
	ElectionType                     int    `position:"0" length:"2"`
	Year                             int    `position:"2" length:"4"`
	Month                            int    `position:"6" length:"2"`
	Round                            int    `position:"8" length:"1"`
	AutonomousCommunity              int    `position:"9" length:"2"`
	Province                         int    `position:"11" length:"2"`
	Municipality                     int    `position:"13" length:"3"`
	MunicipalDistrict                int    `position:"16" length:"2"`
	Name                             string `position:"18" length:"100"`
	ElectoralDistrict                int    `position:"118" length:"1"`
	JudicialDistrict                 int    `position:"119" length:"3"`
	ProvincialCouncil                int    `position:"122" length:"3"`
	Comarca                          int    `position:"125" length:"3"`
	PoblacionDeDerecho               int    `position:"128" length:"8"`
	NumberOfTables                   int    `position:"136" length:"5"`
	CensoIne                         int    `position:"141" length:"8"`
	CensoDeEscrutinio                int    `position:"149" length:"8"`
	CenseCereEscrutinio              int    `position:"157" length:"8"`
	TotalVotantesCere                int    `position:"165" length:"8"`
	FirstParticipationAdvanceVoters  int    `position:"173" length:"8"`
	SecondParticipationAdvanceVoters int    `position:"181" length:"8"`
	BlankVotes                       int    `position:"189" length:"8"`
	NullVotes                        int    `position:"197" length:"8"`
	ValidVotes                       int    `position:"205" length:"8"`
	NumberOfSeats                    int    `position:"213" length:"3"`
}

type MunicipalitiesCandidaturesDataLine struct {
	ElectionType      int `position:"0" length:"2"`
	Year              int `position:"2" length:"4"`
	Month             int `position:"6" length:"2"`
	Round             int `position:"8" length:"1"`
	Province          int `position:"9" length:"2"`
	Municipality      int `position:"11" length:"3"`
	MunicipalDistrict int `position:"14" length:"2"`
	Candidature       int `position:"16" length:"6"`
}
