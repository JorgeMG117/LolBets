package data

type ApiLeague struct {
	Data DataL `json:"data"`
}

type DataL struct {
	Leagues []LeagueL `json:"leagues"`
}

type LeagueL struct {
	ID              string          `json:"id"`
	Slug            string          `json:"slug"`
	Name            string          `json:"name"`
	Region          string          `json:"region"`
	Image           string          `json:"image"`
	Priority        int64           `json:"priority"`
	DisplayPriority DisplayPriority `json:"displayPriority"`
}

type DisplayPriority struct {
	Position int64  `json:"position"`
	Status   Status `json:"status"`
}

type Status string

const (
	ForceSelected Status = "force_selected"
	Hidden        Status = "hidden"
	NotSelected   Status = "not_selected"
	Selected      Status = "selected"
)

