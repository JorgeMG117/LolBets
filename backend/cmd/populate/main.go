package main

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "github.com/joho/godotenv"
    "time"

	"github.com/JorgeMG117/LolBets/backend/models"
    "github.com/JorgeMG117/LolBets/backend/data"
    "github.com/JorgeMG117/LolBets/backend/configs"
)

func createDBtables(db *sql.DB) {
    content, err := ioutil.ReadFile("create.sql")
    if err != nil {
        fmt.Println("Error reading SQL file:", err)
        return
    }

    _, err = db.Exec(string(content))
    if err != nil {
        fmt.Println("Error creating tables:", err)
        return
    }
}

func dropDBtables(db *sql.DB) {
    _, err := db.Exec("DROP TABLE IF EXISTS Game")
    if err != nil {
        fmt.Println("Error dropping Game table:", err)
        return
    }

    _, err = db.Exec("DROP TABLE IF EXISTS Bet")
    if err != nil {
        fmt.Println("Error dropping Bet table:", err)
        return
    }
    _, err = db.Exec("DROP TABLE IF EXISTS Team")
    if err != nil {
        fmt.Println("Error dropping Team table:", err)
        return
    }

    _, err = db.Exec("DROP TABLE IF EXISTS League")
    if err != nil {
        fmt.Println("Error dropping League table:", err)
        return
    }

    _, err = db.Exec("DROP TABLE IF EXISTS User")
    if err != nil {
        fmt.Println("Error dropping User table:", err)
        return
    }

}

func UpdateDatabase(db *sql.DB) {
    // Get last already completed game from DB
    lastGameTime, err := models.GetLastCompletedGameTime(db)
    if lastGameTime == nil {
        fmt.Println("There si no already completed last game:", err)
    }

	// Pillar todos los resultados de la api
    gamesApi, _ := data.GetScheduleApi(time.Now())
    //data.GetScheduleApi(time.Now())
	//fmt.Println(gamesAPI)

	// Pillar todos los partidos incompletos de la bd
	unfinishedGames, err := models.GetUnfinishedGames(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	//fmt.Println(unfinishedGames)

	// See what games stored in the db are completed to update them
	for _, v := range unfinishedGames {
		key := v.Team1 + v.Time.String()
		apiGame := gamesApi[key]
		if apiGame.Completed == 1 {//Game has been played and team 1 won
            v.Completed = 1
		} else if apiGame.Completed == 2 {//Game has been played and team 2 won
            v.Completed = 2
        }
		delete(gamesApi, key)
	}

	// Modificar en la bd unfinishedGames
    go models.UpdateMultipleGames(db, unfinishedGames)

    // Now on apiGames we have only the games that are not in the db yet
    var newGames []models.Game
	for key, game := range gamesApi {
		if game.Completed > 0 {
			delete(gamesApi, key)
		} else {
            newGames = append(newGames, game)
		}
	}
    go models.AddMultipleGames(db, newGames)
}


func InitializeDatabase(db *sql.DB) {
    // Drop DB tables
    dropDBtables(db)

    // Create DB tables
    createDBtables(db)

	// Pillar todos las ligas de la api e insertarlas en la bd
	leaguesAPI := data.GetLeaguesApi()

	for _, league := range leaguesAPI {
		err := models.AddLeague(db, &league)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}

    /*
	// Pillar todos los equipos
	teamsAPI := getTeamsApi(leaguesAPI)
	fmt.Println(teamsAPI)

	for _, team := range teamsAPI {
		err := models.AddTeam(db, &team)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}
    */
}

// TODO
func printUsage() {
    fmt.Println("Usage:")
    fmt.Println("  --update")
    fmt.Println("  --initialize")
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	args := os.Args[1:]

	db := configs.ConnectDB()
    defer db.Close()

	if len(args) == 0 {
		printUsage()
	} else if args[0] == "--update" {
		UpdateDatabase(db)
	} else if args[0] == "--initialize" {
		InitializeDatabase(db)
	}

}
