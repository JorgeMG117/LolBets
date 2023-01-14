package models

type Bet struct {
	Value  int  `json:"value"`
	Team   bool `json:"team"`
	UserId int  `json:"userId"`
	GameId int  `json:"gameId"`
}

var activeBets [][]Bet

func initializeActiveBets() {

}
