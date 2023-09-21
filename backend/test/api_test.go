package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/JorgeMG117/LolBets/backend/configs"
	"github.com/JorgeMG117/LolBets/backend/data"
)

/*
This tests if the data we get from the api and is correctly added to the database
*/

// Add games to the database
//func initializeDbGames(db, games) {

func TestDbEmpty(t *testing.T) {
	// Create connection to DB
	db := configs.ConnectDB()
	defer db.Close()

	// Initialize database
	data.InitializeDatabase(db)

	// Create returned games from the Api
	gamesInApi := data.ApiSchedule{}
	fmt.Println(gamesInApi)

	// Call function update games to database
	data.UpdateApiToDatabase(db, gamesInApi)

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
