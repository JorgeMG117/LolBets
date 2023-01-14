package models

import "database/sql"

type League struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
}

func GetLeaguesName(db *sql.DB) ([]string, error) {
	query := "SELECT Name FROM League"

	rows, err := db.Query(query)

	var leaguesNames []string

	if err != nil {
		return leaguesNames, err
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return leaguesNames, err
		}
		leaguesNames = append(leaguesNames, name)
	}

	return leaguesNames, err
}
