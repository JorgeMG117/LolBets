package datastructures

import (
    "time"
    "strconv"
    "fmt"
    "database/sql"

	"github.com/JorgeMG117/LolBets/backend/models"
)

const (
    MAX_GAMES = 50
)

type ActiveGames struct {
    games []models.Game
    chBets []chan models.Bet
    idToIdx map[int]int

    existsGame []bool
    startIndex int
    numGames int

    activeBets map[int][]models.Bet

    chUpdateGame chan int
}

func (ag *ActiveGames) betController(gameId int) {
	out := false
    idxGame := ag.idToIdx[gameId]
    game := &ag.games[idxGame]
    chBets := ag.chBets[idxGame]

    timeLeft := time.Until(game.Time)
    /*
    fmt.Println("betController: ", timeLeft)
    fmt.Println("Game: ", game)
    fmt.Println("Game idx: ", idxGame)
    */

	for !out {
		select {
		case bet := <-chBets:
			if bet.Team { //Team1
				game.Bets1 += bet.Value
			} else {
				game.Bets2 += bet.Value
			}
			ag.activeBets[game.Id] = append(ag.activeBets[game.Id], bet)
            fmt.Println("Bet placed: ", bet)
            fmt.Println("Game: ", game)
        case <-time.After(timeLeft):
            fmt.Println("Time is up")
            fmt.Println("Game: ", game)
			out = true
            ag.chUpdateGame<-gameId
		}
	}
}

func (ag *ActiveGames) updateActiveGames(db *sql.DB) {
	out := false
	for !out {
        idGame := <-ag.chUpdateGame
        // Update to db activeBets[games[idxGame].Id]
        err := models.AddBets(db, ag.activeBets[idGame])
        if err != nil {
            fmt.Println("Error AddBets: ", err)
        }

        // Update game in db
        err = models.UpdateBetsOnGame(db, ag.GetGameById(idGame))
        if err != nil {
            fmt.Println("Error updateBetsOnGame: ", err)
        }

        fmt.Println("Removing game " + strconv.Itoa(idGame))
        ag.RemoveGame(idGame)
        if ag.numGames < MAX_GAMES - 10 {
            fmt.Println("Trying to fill games")
            // Fetch db to see if there are more games
            err := ag.addMoreActiveGames(db)
            if err != nil {
                fmt.Println("Error addMoreActiveGames: ", err)
            }
        }
        ag.PrintGames()
	}
}

// Creates a new ActiveGames struct
func newActiveGames() *ActiveGames {
    return &ActiveGames{
        games: make([]models.Game, MAX_GAMES),
        chBets: make([]chan models.Bet, MAX_GAMES),
        idToIdx: make(map[int]int),
        existsGame: make([]bool, MAX_GAMES),
        startIndex: 0,
        numGames: 0,
        activeBets: make(map[int][]models.Bet),
    }
}

func InitializeActiveGames(db *sql.DB) (*ActiveGames, error) {
    ag := newActiveGames()

    err := ag.addMoreActiveGames(db)
    ag.PrintGames()
    
    go ag.updateActiveGames(db)

    return ag, err
}


func (ag *ActiveGames) addMoreActiveGames(db *sql.DB) error {
    //Get the last game time
    lastTime := time.Now()
    for _, game := range ag.games {
        if game.Time.After(lastTime) {
            lastTime = game.Time
        }
    }

    games, err := models.GetActiveGames(db, lastTime, MAX_GAMES - ag.numGames)

    if err == nil {
        ag.AddMultipleGames(games)
    }

    return err
}

// Adds a new game to the list of active games
func (ag *ActiveGames) AddGame(game models.Game) {
    //TODO Revisar esto
    for i := 0; i < MAX_GAMES; i++ {
        idx := (i + ag.startIndex) % MAX_GAMES
        if !ag.existsGame[idx] {
            ag.existsGame[idx] = true
            ag.games[idx] = game
            ag.idToIdx[game.Id] = idx
            ag.numGames++
            ag.startIndex = idx + 1
            ag.chBets[idx] = make(chan models.Bet)
            go ag.betController(game.Id)
            return
        }
    }
}

func (ag *ActiveGames) AddMultipleGames(games []models.Game) {
    for _, game := range games {
        ag.AddGame(game)
    }
} 

// Removes a game
func (ag *ActiveGames) RemoveGame(id int) {
    idx := ag.idToIdx[id]
    delete(ag.idToIdx, id)
    ag.existsGame[idx] = false
}

// Returns the list of active games
func (ag *ActiveGames) GetGames(league string, team string) []models.Game {
    var response []models.Game


	for idx, exists := range ag.existsGame {
		if (exists) { 
            val := ag.games[idx]
            if (league == "" || val.League == league) && (team == "" || val.Team1 == team || val.Team2 == team) {
			    response = append(response, val)
            }
		}
	}

    return response
}

func (ag *ActiveGames) GetGameById(id int) models.Game {
    // TODO Check if game exists
    return ag.games[ag.idToIdx[id]]
}

// Place bet on a game
func (ag *ActiveGames) PlaceBet(bet models.Bet, idGame int) {
    /*
    fmt.Println("Placing bet: ", bet)
    fmt.Println("Game id: ", idGame)
    fmt.Println("Game: ", ag.games[ag.idToIdx[idGame]])
    fmt.Println("Game idx: ", ag.idToIdx[idGame])
    */
    ag.chBets[ag.idToIdx[idGame]] <- bet
}

// Print games
func (ag *ActiveGames) PrintGames() {
    //Print games and number of games
    for i := 0; i < ag.numGames; i++ {
        idx := (i + ag.startIndex) % ag.numGames
        fmt.Print("Index: ", idx)
        fmt.Print(", ")
        fmt.Println("Game: ", ag.games[idx])
    }
    fmt.Println("Number of games: ", ag.numGames)
}



