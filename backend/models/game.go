package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
	//"encoding/json"
)

type Game struct {
	Id        int    `json:"id"`
	Team1     string `json:"team1"`
	Team2     string `json:"team2"`
	League    string `json:"league"`
	Time      string `json:"time"`
	Bets1     int    `json:"bets1"`
	Bets2     int    `json:"bets2"`
	Completed bool   `json:"completed"`
}

const MaxGames int = 50

// Active games
// var games [MaxGames]Game
// games := make([]Game, 0, MaxGames)
var games []Game
var indexOfGame []int

// TODO Moverlo a model.bet o algo asi
type Bet struct {
	Value int
	Team  bool
	//Usario???
}

// TODO
func BetController(chBets chan Bet, idxGame int) {
	out := false
	for !out {
		select {
		case bet := <-chBets:
			if bet.Team { //Team1
				games[idxGame].Bets1 += bet.Value
			} else {
				games[idxGame].Bets2 += bet.Value
			}
			fmt.Println(games)
		case <-time.After(10 * time.Minute):
			fmt.Println("Bet " + strconv.Itoa(idxGame) + " is over")
			out = true
		}
	}
}

func NumGames() int {
	return len(games)
}

func GetIdxOfGame(id int) int {
	for i := 0; i < len(indexOfGame); i++ {
		if indexOfGame[i] == id {
			return i
		}
	}
	return -1
}

func InitializeGames(db *sql.DB) error {
	games = make([]Game, 0, MaxGames)
	indexOfGame = make([]int, 0, MaxGames)

	//Cojemos como mucho maxGames partidos
	query := "SELECT g.Id, t1.Name, t2.Name, l.Name, g.Time, g.Bets_t1, g.Bets_t2 FROM Game g, Team t1, Team t2, League l WHERE t1.Id = g.Team_1 AND t2.Id = g.Team_2 AND l.Id = g.League"
	query = query + " LIMIT " + strconv.Itoa(MaxGames)

	rows, err := db.Query(query)

	if err != nil {
		return err
	}

	for rows.Next() {
		var game Game
		err = rows.Scan(&game.Id, &game.Team1, &game.Team2, &game.League, &game.Time, &game.Bets1, &game.Bets2)
		if err != nil {
			return err
		}
		games = append(games, game)
		indexOfGame = append(indexOfGame, game.Id)
	}

	fmt.Println(games)

	return nil
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
	result, err := db.Exec("INSERT INTO Game(Team_1, Team_2, League) SELECT t1.Id, t2.Id, l.Id FROM Team t1, Team t2, League l WHERE t1.Name = ? AND t2.Name = ? AND l.Name = ?", newGame.Team1, newGame.Team2, newGame.League)
	if val, _ := result.RowsAffected(); val != 1 {
		fmt.Println("No se ha insertado nada o se ha insertado mas de un valor")
		fmt.Println(newGame)
	}
	return err
}

func GetUnfinishedGames() []Game {
	return games
}
