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
func BetController(chBets chan Bet, idxGame int, chUpdateGame chan int) {
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
            chUpdateGame<-idxGame
		}
	}
}

//updateActiveGames updates the games slice with the new games
//espero a que me avisen por un canal que game hay que quitar
//lo quito
//si hay hueco mayor que 10 
//llamo a funcion para intentar rellenar hueco

//Param 
func updateActiveGames(db *sql.DB, chUpdateGame chan int) {
	out := false
	for !out {
        idxGame := <-chUpdateGame
        // Update to db activeBets[games[idxGame].Id]
        err := AddBets(db, activeBets[games[idxGame].Id])
        if err != nil {
            fmt.Println("Error AddBets: ", err)
        }

        fmt.Println("Removing game " + strconv.Itoa(games[idxGame].Id))
        games = append(games[:idxGame], games[idxGame+1:]...)
        //TODO: Why do we need indexOfGame
        indexOfGame = append(indexOfGame[:idxGame], indexOfGame[idxGame+1:]...)
        fmt.Println("Num games: " + strconv.Itoa(len(games)))
        if len(games) < MaxGames - 10 {
            fmt.Println("Trying to fill games")
            // Fetch db to see if there are more games
            err := addMoreActiveGames(db)
            if err != nil {
                fmt.Println("Error addMoreActiveGames: ", err)
            }
        }
        for _, game := range games {
            fmt.Println(game)
        }
	}
}

func Scan_game(rows *sql.Rows) (Game, error) {
    var game Game
    var horario string
    err := rows.Scan(&game.Id, &game.Team1, &game.Team2, &game.League, &horario, &game.Bets1, &game.Bets2, &game.Completed, &game.BlockName, &game.Strategy)
    if err != nil {
        return game, err
    }
    game.Time, err = time.Parse("2006-01-02 15:04:05", horario)
    if err != nil {
        return game, err
    }
    return game, nil
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

//Adds more games to the active games slice
func addMoreActiveGames(db *sql.DB) error {
    //Get the last game time
    lastTime := time.Now()
    for _, game := range games {
        if game.Time.After(lastTime) {
            lastTime = game.Time
        }
    }

	query := "SELECT g.Id, t1.Name, t2.Name, l.Name, g.Time, g.Bets_t1, g.Bets_t2, g.Completed, g.BlockName, g.Strategy FROM Game g, Team t1, Team t2, League l WHERE t1.Code = g.Team_1 AND t2.Code = g.Team_2 AND l.Id = g.League AND g.Completed = 0"
    query = query + " AND g.Time > '" + lastTime.Format("2006-01-02 15:04:05") + "'"
    query = query + " ORDER BY g.Time ASC"
	query = query + " LIMIT " + strconv.Itoa(MaxGames - len(games))

	rows, err := db.Query(query)

	if err != nil {
		return err
	}

	for rows.Next() {
        game, err := Scan_game(rows)
		if err != nil {
			return err
		}

		games = append(games, game)
		indexOfGame = append(indexOfGame, game.Id)
	}

    return nil
}

func InitializeGames(db *sql.DB, chUpdateGames chan int) error {
	games = make([]Game, 0, MaxGames)
	indexOfGame = make([]int, 0, MaxGames)

	//Cojemos como mucho maxGames partidos
    //TODO Ver si habria que cojer los no completados cuya fecha de inicio sea superior al ultimo partido completado
	query := "SELECT g.Id, t1.Name, t2.Name, l.Name, g.Time, g.Bets_t1, g.Bets_t2, g.Completed, g.BlockName, g.Strategy FROM Game g, Team t1, Team t2, League l WHERE t1.Code = g.Team_1 AND t2.Code = g.Team_2 AND l.Id = g.League AND g.Completed = 0"
    query = query + " ORDER BY g.Time ASC"
	query = query + " LIMIT " + strconv.Itoa(MaxGames)

	rows, err := db.Query(query)

	if err != nil {
		return err
	}

	for rows.Next() {
        game, err := Scan_game(rows)
		if err != nil {
			return err
		}

		games = append(games, game)
		indexOfGame = append(indexOfGame, game.Id)
	}

	initializeActiveBets()

    fmt.Println("Games initialized")
    for _, game := range games {
	    fmt.Println(game)
    }
    fmt.Println("Num games: " + strconv.Itoa(len(games)))

    //TODO Find a way to update the game slice for when games are over
    go updateActiveGames(db, chUpdateGames)


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
		if (league == "" || val.League == league) && (team == "" || val.Team1 == team || val.Team2 == team) {
			response = append(response, val)
		}
	}
	return response, nil
}


func GetGamesDb(db *sql.DB, league string, team string) ([]Game, error) {
    query := "SELECT g.Id, t1.Name, t2.Name, l.Name, g.Time, g.Bets_t1, g.Bets_t2, g.Completed, g.BlockName, g.Strategy FROM Game g, Team t1, Team t2, League l WHERE t1.Code = g.Team_1 AND t2.Code = g.Team_2 AND l.Id = g.League"
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
        game, err := Scan_game(rows)
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

// TODO: Maybe just update those who have changed, completed <> 0
func UpdateMultipleGames(db *sql.DB, games []Game) error {
    for _, val := range games {
        if val.Completed == 0 {
            continue
        }
        err := UpdateGame(db, &val)
        if err != nil {
            return err
        }
    }
    return nil
}

func UpdateGame(db *sql.DB, game *Game) error {
    fmt.Println("Updating game: ", game)
    fmt.Println("Completed: ", game.Completed)
    fmt.Println("Id: ", game.Id)
    _, err := db.Exec("UPDATE Game SET Bets_t1 = ?, Bets_t2 = ?, Completed = ? WHERE Id = ?", game.Bets1, game.Bets2, game.Completed, game.Id)
    return err
}

func GetUnfinishedGames(db *sql.DB) ([]Game, error) {
	query := "SELECT g.Id, t1.Name, t2.Name, l.Name, g.Time, g.Bets_t1, g.Bets_t2, g.Completed, g.BlockName, g.Strategy FROM Game g, Team t1, Team t2, League l WHERE t1.Code = g.Team_1 AND t2.Code = g.Team_2 AND l.Id = g.League AND g.Completed = 0"

	rows, err := db.Query(query)

	var games []Game

	if err != nil {
		return games, err
	}

	for rows.Next() {
        game, err := Scan_game(rows)
        if err != nil {
            return games, err
        }
		games = append(games, game)
	}

	return games, nil
}

func GetLastCompletedGameTime(db *sql.DB) (time.Time, error) {
    var horario string
    var lastGameTime time.Time
    err := db.QueryRow("SELECT Time FROM Game WHERE Completed <> 0 ORDER BY Time DESC LIMIT 1").Scan(&horario)
    if err != nil {
        //Get the earliest game that is not completed
        err = db.QueryRow("SELECT Time FROM Game WHERE Completed = 0 ORDER BY Time ASC LIMIT 1").Scan(&horario)
        if err != nil {
            return time.Now(), nil 
        }
    }
    lastGameTime, err = time.Parse("2006-01-02 15:04:05", horario)
    if err != nil {
        return lastGameTime, err
    }
    return lastGameTime, nil
}


