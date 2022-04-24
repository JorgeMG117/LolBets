package models

type Game struct {
	Id    string  `json:"id"`
	Team1 string  `json:"team1"`
	Team2 string  `json:"team2"`
	Odds  float64 `json:"price"`
}

var games = []Game{
	{Id: "1", Team1: "G2", Team2: "Koi", Odds: 56.99},
	{Id: "2", Team1: "T1", Team2: "Fnatic", Odds: 56.99},
}

func GetGames() []Game {
	return games
}

func AddGame(newGame Game) {
	games = append(games, newGame)
}
