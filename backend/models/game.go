package models

import (
	"database/sql"
	"fmt"
	//"encoding/json"
)

type Game struct {
	Id    string  `json:"id"`
	Team1 string  `json:"team1"`
	Team2 string  `json:"team2"`
	League string  `json:"league"`
	Odds  float64 `json:"odds"`
}

func GetGames(db *sql.DB, league string, team string) ([]Game, error) {
    query := "SELECT t1.Name, t2.Name, l.Name FROM Game g, Team t1, Team t2, League l WHERE t1.Id = g.Team_1 AND t2.Id = g.Team_2 AND l.Id = g.League"
    if league != "" {
        query = query + " AND l.Name = " + league
    }
    if team != "" {
        query = query + " AND (t1.Name = " + team + " OR t2.Name = " + team + ")"
    } 

    rows, err := db.Query(query)
    fmt.Println(rows)

    var games []Game

    if err != nil {
        return games, err
    }

    for rows.Next() {
        var game Game 
        err = rows.Scan(&game.Team1, &game.Team2, &game.League)
        if err != nil {
                return games, err
        }
        games = append(games, game)
    }

	return games, nil
}

func AddGame(db *sql.DB, newGame *Game) error {
    _, err := db.Exec("INSERT INTO Game(Team_1, Team_2, League) SELECT t1.Id, t2.Id, l.Id FROM Team t1, Team t2, League l WHERE t1.Name = ? AND t2.Name = ? AND l.Name = ?", newGame.Team1, newGame.Team2, newGame.League)
    return err
}
