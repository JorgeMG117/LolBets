package models

import (
	"database/sql"
	"fmt"
)

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

func AddLeague(db *sql.DB, newLeague *League) error {
	result, err := db.Exec("INSERT INTO League(Name, Slug, Image) VALUES (?, ?, ?)", newLeague.Name, newLeague.Slug, newLeague.Image)
	if val, _ := result.RowsAffected(); val != 1 {
		fmt.Println("No se ha insertado nada o se ha insertado mas de un valor")
		fmt.Println(newLeague)
	}
	return err
}
