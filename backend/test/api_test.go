package test

import (
	"net/http"
	"testing"
	"time"
    "fmt"

	"github.com/JorgeMG117/LolBets/backend/configs"
	"github.com/JorgeMG117/LolBets/backend/models"
	//"github.com/JorgeMG117/LolBets/backend/data"
)

/*
This tests if the data we get from the api and is correctly added to the database
*/

// Add games to the database
//func initializeDbGames(db, games) {

func TestDbEmpty(t *testing.T) {
	db := configs.ConnectDB()
    defer db.Close()

    // Initialize database

    // Empty database

    // Create returned games from the Api
    gamesInApi := []models.Game {
        {
            Id: 1,
            Team1: "G2",
            Team2: "FNC",
            League: "LEC",
            Time: time.Now(),
            Bets1: 0,
            Bets2: 0,
            BlockName: "Finals",
            Strategy: "best of 5",
        },
    }
    fmt.Println(gamesInApi)

    // Call function update games to database
    //data.UpdateGames(db, gamesInApi)

    // Get games from the database

    // Check if database has the correct values
    /*
    expectedGames := []models.Game {
    }
    */
}

func TestFetchGophers(t *testing.T) {
	_, err := http.NewRequest("GET", "/gophers", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
}
