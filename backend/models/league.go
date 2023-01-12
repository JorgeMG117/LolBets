package models

type League struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
}

func GetLeaguesName() []string {
	return []string{"LEC"}
}
