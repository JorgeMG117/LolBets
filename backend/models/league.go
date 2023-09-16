package models

import (
	"database/sql"
	"fmt"
)

type League struct {
    ApiID  string `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
	Image  string `json:"image"`
}

// Print leagues from DB
func PrintLeagues(db *sql.DB) error {
    rows, err := db.Query("SELECT * FROM League")

    if err != nil {
        return err
    }

    for rows.Next() {
        var league League
        err = rows.Scan(&league.ApiID, &league.Name, &league.Region, &league.Image)
        if err != nil {
            return err
        }
        fmt.Println(league)
    }

    return nil
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
	result, err := db.Exec("INSERT INTO League(ApiID, Name, Region, Image) VALUES (?, ?, ?, ?)", newLeague.ApiID, newLeague.Name, newLeague.Region, newLeague.Image)

	if err != nil {
		return err
	}

	if val, _ := result.RowsAffected(); val != 1 {
		fmt.Println("No se ha insertado nada o se ha insertado mas de un valor")
		fmt.Println(newLeague.Name)
	}

	return err
}

func GetLeagues(db *sql.DB, league string, team string) ([]Game, error) {
	return nil, nil
	// query := "SELECT t1.Name, t2.Name, l.Name FROM Game g, Team t1, Team t2, League l WHERE t1.Id = g.Team_1 AND t2.Id = g.Team_2 AND l.Id = g.League"
	// if league != "" {
	// 	query = query + " AND l.Name = " + league
	// }
	// if team != "" {
	// 	query = query + " AND (t1.Name = " + team + " OR t2.Name = " + team + ")"
	// }

	// rows, err := db.Query(query)

	// var games []Game

	// if err != nil {
	// 	return games, err
	// }

	// for rows.Next() {
	// 	var game Game
	// 	err = rows.Scan(&game.Team1, &game.Team2, &game.League)
	// 	if err != nil {
	// 		return games, err
	// 	}
	// 	games = append(games, game)
	// }

	// return games, nil
}
