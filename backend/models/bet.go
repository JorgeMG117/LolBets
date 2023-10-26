package models

import (
    "database/sql"
)

type Bet struct {
	Value  int  `json:"value"`
	Team   bool `json:"team"`
	UserId int  `json:"userId"`
	GameId int  `json:"gameId"`
}

var activeBets map[int][]Bet

func initializeActiveBets() {
	activeBets = make(map[int][]Bet)
}

func AddBet(db *sql.DB, bet Bet) error {
    _, err := db.Exec("INSERT INTO Bet(GameId, UserId, Value, Team) VALUES (?, ?, ?, ?)", bet.GameId, bet.UserId, bet.Value, bet.Team)
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

