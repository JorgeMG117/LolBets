package test

import (
	"fmt"
	//"net/http"
	"log"
	"testing"
    //"time"

	"github.com/joho/godotenv"

	"github.com/JorgeMG117/LolBets/backend/configs"
	"github.com/JorgeMG117/LolBets/backend/data"
	"github.com/JorgeMG117/LolBets/backend/models"
)

/*
This tests if the data we get from the api and is correctly added to the database
*/

// Add games to the database
//func initializeDbGames(db, games) {


// Opciones para testear:
// 1. Comprobar que cambian el numero de filas en la bd que se esperan
// 2. Comprobar que los valores de la bd son los que se esperan

func TestDbEmpty(t *testing.T) {
	t.Skip("SKIPPED soloArranqueYparadaTest1")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Create connection to DB
	db := configs.ConnectDB()
	defer db.Close()

	// Initialize database
	data.InitializeDatabase(db)

	// Create returned games from the Api
    // Read file with games
    gamesInApi := data.ReadApiSchedule("schedule_test.json")
	//fmt.Println(gamesInApi)

	// Call function update games to database
	data.UpdateApiToDatabase(db, gamesInApi)

	// Get games from the database
    gamesInDb, err := models.GetGamesDb(db, "", "")
    if err != nil {
        t.Fatalf("could not get games from database: %v", err)
    }
    fmt.Println("Games in database:")
    fmt.Println(gamesInDb)

	// Check if database has the correct values
    /*
    [{0 Team BDS Golden Guardians Worlds Qualifying Series 2023-10-09 04:00:00 +0000 UTC 0 0 0 Finals best of 5} {0 Movistar R7 PSG Talon Worlds 2023-10-10 07:00:00 +0000 UTC 0 0 0 Play In Groups best of 3} {0 LOUD GAM Esports Worlds 2023-10-10 10:00:00 +0000 UTC 0 0 0 Play In Groups best of 3} {0 DetonatioN FocusMe CTBC Flying Oyster Worlds 2023-10-11 07:00:00 +0000 UTC 0 0 0 Play In Groups best of 3}]
    */
    //string to time
    //time.Parse("2006-01-02 15:04:05", "2023-10-09 04:00:00")
    /*
    expectedGames := []models.Game{
        {0, "Team BDS", "Golden Guardians", "Worlds Qualifying Series", time.Parse("2006-01-02 15:04:05", "2023-10-09 04:00:00"), 0, 0, 0, "Finals", "best of 5"},
    }
    fmt.Println(expectedGames)
    */

}

func TestUpdateOneAddOne(t *testing.T) {} 

/*
func TestFetchGophers(t *testing.T) {
	_, err := http.NewRequest("GET", "/gophers", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
}
*/
