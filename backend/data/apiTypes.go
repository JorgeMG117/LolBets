package data

/*
type Data struct {
	Data schedule `json:"data"`
}

type schedule struct {
	Schedule events `json:"schedule"`
}

type events struct {
	Events []game `json:"events"`
}

type match struct {
	Teams [2]models.Team `json:"teams"`
}

type game struct {
	Id int `json:"id"`

	StartTime time.Time     `json:"startTime"`
	BlockName string        `json:"blockName"`
	State     string        `json:"state"`
	Type      string        `json:"type"`
	Match     match         `json:"match"`
	League    models.League `json:"league"`
}

type LeaguesData struct {
	LeaguesData leagues `json:"data"`
}

type leagues struct {
	Leagues []models.League `json:"leagues"`
}
*/

type ApiSchedule struct {
	Data Data `json:"data"`
}

type Data struct {
	Schedule Schedule `json:"schedule"`
}

type Schedule struct {
	Pages  Pages   `json:"pages"`
	Events []Event `json:"events"`
}

type Event struct {
	StartTime string     `json:"startTime"`
	State     State      `json:"state"`
	Type      EventType  `json:"type"`
	BlockName BlockName  `json:"blockName"`
	League    League     `json:"league"`
	Match     MatchClass `json:"match"`
}

type League struct {
	Name Name `json:"name"`
	Slug Slug `json:"slug"`
}

type MatchClass struct {
	ID       string   `json:"id"`
	Flags    []Flag   `json:"flags"`
	Teams    []Team   `json:"teams"`
	Strategy Strategy `json:"strategy"`
}

type Strategy struct {
	Type  StrategyType `json:"type"`
	Count int64        `json:"count"`
}

type Team struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Image  string `json:"image"`
	Result Result `json:"result"`
	Record Record `json:"record"`
}

type Record struct {
	WINS   int64 `json:"wins"`
	Losses int64 `json:"losses"`
}

type Result struct {
	Outcome  *Outcome `json:"outcome"`
	GameWINS int64    `json:"gameWins"`
}

type Pages struct {
	Older string      `json:"older"`
	Newer interface{} `json:"newer"`
}

type BlockName string

const (
	CuartosDeFinal                  BlockName = "Cuartos de final"
	EliminatoriasDeLaFaseDeApertura BlockName = "Eliminatorias de la fase de apertura"
	Final                           BlockName = "Final"
	Grupos                          BlockName = "Grupos"
	GruposDeLaFaseDeApertura        BlockName = "Grupos de la fase de apertura"
	Semifinales                     BlockName = "Semifinales"
)

type Name string

const (
	Mundial                Name = "Mundial"
	WorldsQualifyingSeries Name = "Worlds Qualifying Series"
)

type Slug string

const (
	Worlds Slug = "worlds"
	Wqs    Slug = "wqs"
)

type Flag string

const (
	HasVOD    Flag = "hasVod"
	IsSpoiler Flag = "isSpoiler"
)

type StrategyType string

const (
	BestOf StrategyType = "bestOf"
)

type Outcome string

const (
	Loss Outcome = "loss"
	Win  Outcome = "win"
)

type State string

const (
	Completed State = "completed"
	Unstarted State = "unstarted"
)

type EventType string

const (
	Match EventType = "match"
)

