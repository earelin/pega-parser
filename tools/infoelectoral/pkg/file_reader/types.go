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

type DatosComunsDeMesasECera struct {
	TipoEleccion                       int    `position:"0" length:"2"`
	Ano                                int    `position:"2" length:"4"`
	Mes                                int    `position:"6" length:"2"`
	NumeroDeVolta                      int    `position:"8" length:"1"`
	CodigoComunidadeAutonoma           int    `position:"9" length:"2"`
	CodigoProvincia                    int    `position:"11" length:"2"`
	CodigoMunicipio                    int    `position:"13" length:"3"`
	NumeroDistritoMunicipal            int    `position:"16" length:"2"`
	Seccion                            int    `position:"18" length:"4"`
	Mesa                               string `position:"22" length:"1"`
	CensoIne                           int    `position:"23" length:"7"`
	CensoEscrutinioOCera               int    `position:"30" length:"7"`
	CensoCereEnEscrutinio              int    `position:"37" length:"7"`
	TotalVotantesCere                  int    `position:"44" length:"7"`
	VotantesPrimerAvanceParticipacion  int    `position:"51" length:"7"`
	VotantesSecundoAvanceParticipacion int    `position:"58" length:"7"`
	VotosBlanco                        int    `position:"65" length:"7"`
	VotosNulos                         int    `position:"72" length:"7"`
	VotosACandidaturas                 int    `position:"79" length:"7"`
	VotosAfirmativosReferendum         int    `position:"86" length:"7"`
	VotosNegativosReferendum           int    `position:"93" length:"7"`
	DatosOficiales                     string `position:"100" length:"1"`
}

type DatoCandidaturasDeMesasECera struct {
	TipoEleccion             int    `position:"0" length:"2"`
	Ano                      int    `position:"2" length:"4"`
	Mes                      int    `position:"6" length:"2"`
	NumeroDeVolta            int    `position:"8" length:"1"`
	CodigoComunidadeAutonoma int    `position:"9" length:"2"`
	CodigoProvincia          int    `position:"11" length:"2"`
	CodigoMunicipio          int    `position:"13" length:"3"`
	NumeroDistritoMunicipal  int    `position:"16" length:"2"`
	Seccion                  int    `position:"18" length:"4"`
	Mesa                     string `position:"22" length:"1"`
	CodigoCandidatura        int    `position:"23" length:"6"`
	Votos                    int    `position:"29" length:"7"`
}
