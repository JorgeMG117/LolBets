package main

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "github.com/joho/godotenv"

	//"github.com/JorgeMG117/LolBets/backend/models"
    "github.com/JorgeMG117/LolBets/backend/data"
    //"github.com/JorgeMG117/LolBets/backend/configs"
)

func createDBtables(db *sql.DB) {
    content, err := ioutil.ReadFile("tables.sql")
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

/*
func UpdateDatabase() {
	db := configs.ConnectDB()

	// Pillar todos los resultados de la api
	gamesAPI := getScheduleApi(db)
	//fmt.Println(gamesAPI)

	// Pillar todos los partidos incompletos de la bd
	unfinishedGames, err := models.GetUnfinishedGames(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	//fmt.Println(unfinishedGames)

	db.Close()

	// See what games stored in the db are completed to update them
	for _, v := range unfinishedGames {
		key := v.Team1 + v.Time.String()
		apiGame := gamesAPI[key]
		if apiGame.State == "completed" {
			//Change unfinishedGames
			fmt.Println("Change unfinishedGames")
		}
		delete(gamesAPI, key)
	}

	// //go Modificar en la bd unfinishedGames

	for key, game := range gamesAPI {
		if game.State == "completed" {
			delete(gamesAPI, key)
		} else {
			//models.AddGame(db, &game)
		}
	}

	// // Recorriendo partidos de la bd
	// // Encontrar el correspondiente en la llamada a la api
	// // Si APIcompleted y BDcompleted no hacemos nada
	// //
	// // Ir eliminando de la api los que vas recorrienod

	// //Quitar de la api el resto de completed
	// //Añadir lo que queda en los de la api (uncompleted a la bd)

}

func InitializeDatabase() {
	db := configs.ConnectDB()

    // Create DB tables
    createDBtables(db)

	// Pillar todos las ligas de la api
	leaguesAPI := getLeaguesApi()

	for _, league := range leaguesAPI {
		err := models.AddLeague(db, &league)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}

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

	db.Close()
}
*/

// TODO
func printUsage() {
	fmt.Println("This display the usage of the populate database program")
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	args := os.Args[1:]

	if len(args) == 0 {
		printUsage()
	} else if args[0] == "--update" {
		//UpdateDatabase()
	} else if args[0] == "--initialize" {
		InitializeDatabase()
	}

}
