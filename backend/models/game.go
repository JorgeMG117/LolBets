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

func UpdateBetsOnGame(db *sql.DB, game Game) error {
    fmt.Println("Updating game: ", game)
    _, err := db.Exec("UPDATE Game SET Bets_t1 = ?, Bets_t2 = ? WHERE Id = ?", game.Bets1, game.Bets2, game.Id)
    return err
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
    _, err := db.Exec("UPDATE Game SET Bets_t1 = ?, Bets_t2 = ?, Completed = ? WHERE Id = ?", game.Bets1, game.Bets2, game.Completed, game.Id)
    if err != nil {
        return err
    }

    // Update coins of users
    query := `
        UPDATE User u
        JOIN (
            SELECT b.UserId, SUM(b.Value * b.Odds) AS TotalWinnings
            FROM Bet b
            INNER JOIN Game g ON b.GameId = g.Id
            WHERE b.GameId = ?
            AND (
                (b.Team = 0 AND g.Completed = 1) OR 
                (b.Team = 1 AND g.Completed = 2)
            )
            GROUP BY b.UserId
        ) AS WinningBets ON u.Id = WinningBets.UserId
        SET u.Coins = u.Coins + WinningBets.TotalWinnings;
        `
    _, err = db.Exec(query, game.Id)

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


func GetActiveGames(db *sql.DB, lastTime time.Time, numGames int) ([]Game, error) {
    query := "SELECT g.Id, t1.Name, t2.Name, l.Name, g.Time, g.Bets_t1, g.Bets_t2, g.Completed, g.BlockName, g.Strategy FROM Game g, Team t1, Team t2, League l WHERE t1.Code = g.Team_1 AND t2.Code = g.Team_2 AND l.Id = g.League AND g.Completed = 0"
    query = query + " AND g.Time > '" + lastTime.Format("2006-01-02 15:04:05") + "'"
    query = query + " ORDER BY g.Time ASC"
    query = query + " LIMIT " + strconv.Itoa(numGames)

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

