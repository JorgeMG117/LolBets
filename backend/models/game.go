package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
	//"encoding/json"
)

type Game struct {
	Id        int       `json:"id"`
	Team1     string    `json:"team1"`
	Team2     string    `json:"team2"`
	League    string    `json:"league"`
	Time      time.Time `json:"time"`
	Bets1     int       `json:"bets1"`
	Bets2     int       `json:"bets2"`
	Completed int       `json:"completed"`
	BlockName string    `json:"blockName"`
    Strategy  string    `json:"strategy"`
}

const MaxGames int = 50

// Active games
// var games [MaxGames]Game
// games := make([]Game, 0, MaxGames)
var games []Game
var indexOfGame []int

// TODO
func BetController(chBets chan Bet, idxGame int) {
	out := false
	timeLeft := time.Until(games[idxGame].Time)
    /*
    fmt.Println("Game " + games[idxGame].Team1 + " vs " + games[idxGame].Team2)
    fmt.Println(idxGame)
    fmt.Println("Time left: " + strconv.Itoa(int(timeLeft.Minutes())) + " minutes")
    */
	for !out {
		select {
		case bet := <-chBets:
			if bet.Team { //Team1
				games[idxGame].Bets1 += bet.Value
			} else {
				games[idxGame].Bets2 += bet.Value
			}
			activeBets[games[idxGame].Id] = append(activeBets[games[idxGame].Id], bet)
			fmt.Println(games)
        case <-time.After(timeLeft):
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
	query := "SELECT g.Id, t1.Name, t2.Name, l.Name, g.Time, g.Bets_t1, g.Bets_t2, g.BlockName, g.Strategy FROM Game g, Team t1, Team t2, League l WHERE t1.Code = g.Team_1 AND t2.Code = g.Team_2 AND l.Id = g.League AND g.Completed = 0"
    query = query + " ORDER BY g.Time ASC"
	query = query + " LIMIT " + strconv.Itoa(MaxGames)

	rows, err := db.Query(query)

	if err != nil {
		return err
	}

	for rows.Next() {
		var game Game
        var horario string
		err = rows.Scan(&game.Id, &game.Team1, &game.Team2, &game.League, &horario, &game.Bets1, &game.Bets2, &game.BlockName, &game.Strategy)
		if err != nil {
			return err
		}
        game.Time, err = time.Parse("2006-01-02 15:04:05", horario)
		if err != nil {
			return err
		}

		games = append(games, game)
		indexOfGame = append(indexOfGame, game.Id)
	}

	initializeActiveBets()

    fmt.Println("Games initialized")
	fmt.Println(games)

	return nil
}

// TODO: No se si habria que devolver mejor los guardados en memoria
func GetGames(db *sql.DB, league string, team string) ([]Game, error) {
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

	var response []Game

	for _, val := range games {
		if val.League == league && (val.Team1 == team || val.Team2 == team) {
			response = append(response, val)
		}
	}
	return response, nil
}


func GetGamesDb(db *sql.DB, league string, team string) ([]Game, error) {
    query := "SELECT t1.Name, t2.Name, l.Name, g.Time, g.Completed, g.BlockName, g.Strategy FROM Game g, Team t1, Team t2, League l WHERE t1.Code = g.Team_1 AND t2.Code = g.Team_2 AND l.Id = g.League"
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
        var horario string
        err = rows.Scan(&game.Team1, &game.Team2, &game.League, &horario, &game.Completed, &game.BlockName, &game.Strategy)
        if err != nil {
            return games, err
        }
        game.Time, err = time.Parse("2006-01-02 15:04:05", horario)
        if err != nil {
            return games, err
        }
        games = append(games, game)
    }

    return games, nil
}


/*
func AddGame(db *sql.DB, newGame *Game) error {
    result, err := db.Exec("INSERT INTO Game(Team_1, Team_2, League, Time, Bets_t1, Bets_t2, Completed, BlockName) SELECT t1.Code, t2.Code, l.Id, ?, ?, ?, ?, ? FROM Team t1, Team t2, League l WHERE t1.Name = ? AND t2.Name = ? AND l.Name = ?", newGame.Time, newGame.Bets1, newGame.Bets2, newGame.Completed, newGame.BlockName, newGame.Team1, newGame.Team2, newGame.League)
    if err != nil {
        return err
    }
    if val, _ := result.RowsAffected(); val != 1 {
        fmt.Println("No se ha insertado nada o se ha insertado mas de un valor")
        fmt.Println(newGame)
    }
    return nil
}
*/
func AddGame(db *sql.DB, newGame *Game) error {
    result, err := db.Exec("INSERT INTO Game(Team_1, Team_2, League, Time, Bets_t1, Bets_t2, Completed, BlockName, Strategy) SELECT t1.Code, t2.Code, l.Id, ?, ?, ?, ?, ?, ? FROM Team t1, Team t2, League l WHERE t1.Name = ? AND t2.Name = ? AND l.Name = ?", newGame.Time, newGame.Bets1, newGame.Bets2, newGame.Completed, newGame.BlockName, newGame.Strategy, newGame.Team1, newGame.Team2, newGame.League)
    if err != nil {
        return err
    }
    if val, _ := result.RowsAffected(); val != 1 {
        fmt.Println("No se ha insertado nada o se ha insertado mas de un valor")
        fmt.Println(newGame)
    }
    return nil
}



func AddMultipleGames(db *sql.DB, newGames []Game) error {
    for _, val := range newGames {
        err := AddGame(db, &val)
        if err != nil {
            return err
        }
    }
    return nil
}

func UpdateMultipleGames(db *sql.DB, games []Game) error {
    for _, val := range games {
        err := UpdateGame(db, &val)
        if err != nil {
            return err
        }
    }
    return nil
}

func UpdateGame(db *sql.DB, game *Game) error {
    _, err := db.Exec("UPDATE Game SET Bets_t1 = ?, Bets_t2 = ?, Completed = ? WHERE Id = ?", game.Bets1, game.Bets2, game.Completed, game.Id)
    return err
}

func GetUnfinishedGames(db *sql.DB) ([]Game, error) {
	query := "SELECT t1.Name, t2.Name, l.Name, g.Time, g.Completed, g.BlockName FROM Game g, Team t1, Team t2, League l WHERE t1.Code = g.Team_1 AND t2.Code = g.Team_2 AND l.Id = g.League AND g.Completed = 0"

	rows, err := db.Query(query)

	var games []Game

	if err != nil {
		return games, err
	}

	for rows.Next() {
		var game Game
		err = rows.Scan(&game.Team1, &game.Team2, &game.League, &game.Time, &game.Completed, &game.BlockName)
		if err != nil {
			return games, err
		}
		games = append(games, game)
	}

	return games, nil
}

func GetLastCompletedGameTime(db *sql.DB) (time.Time, error) {
    var lastGameTime time.Time
    err := db.QueryRow("SELECT Time FROM Game WHERE Completed = 1 ORDER BY Time DESC LIMIT 1").Scan(&lastGameTime)
    if err != nil {
        return time.Now(), err
    }
    return lastGameTime, nil
}















