package models

import (
    "database/sql"
)

type Bet struct {
	Value  int  `json:"value"`
	Team   bool `json:"team"`
	UserId int  `json:"userId"`
	GameId int  `json:"gameId"`
    Odds   float32 `json:"odds"`
}

func AddBet(db *sql.DB, bet Bet) error {
    _, err := db.Exec("INSERT INTO Bet(GameId, UserId, Value, Team, Odds) VALUES (?, ?, ?, ?, ?)", bet.GameId, bet.UserId, bet.Value, bet.Team, bet.Odds)
    if err != nil {
        return err
    }
    return nil
}

func AddBets(db *sql.DB, bets []Bet) error {
    for _, bet := range bets {
        err := AddBet(db, bet)
        if err != nil {
            return err
        }
    }
    return nil 

}

//Get completed bets of a user
func GetBetsOfUser(db *sql.DB, userId int) (*sql.Rows, error) {
	query := `
		SELECT Bet.GameId, Bet.UserId, Bet.Value, Bet.Team, Bet.Odds,
		       Game.Team_1 AS Team1, Game.Team_2 AS Team2,
		       League.Name AS LeagueName, Game.Completed
		FROM Bet
		JOIN Game ON Bet.GameId = Game.Id
		JOIN League ON Game.League = League.Id
		WHERE Bet.UserId = ?
		  AND Game.Completed > 0
		ORDER BY Game.Time ASC
		LIMIT 5;
        `

    rows, err := db.Query(query, userId)
    if err != nil {
        return nil, err
    }
    //defer rows.Close()

    return rows, err
}
