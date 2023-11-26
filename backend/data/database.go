package data

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/JorgeMG117/LolBets/backend/models"
)

func createDBtables(db *sql.DB) {
	fmt.Println("Creating tables...")
	/*
	   //content, err := ioutil.ReadFile("create.sql")
	   content, err := os.ReadFile("create.sql")
	   if err != nil {
	       fmt.Println("Error reading SQL file:", err)
	       return
	   }
	*/
	createTableSQL := []string{
		`CREATE TABLE User (
            Id      INT AUTO_INCREMENT PRIMARY KEY,
            Name    VARCHAR(50) UNIQUE,
            Coins   INT DEFAULT 0
        );`,
		`CREATE TABLE League (
            Id      INT AUTO_INCREMENT PRIMARY KEY,
            ApiID   VARCHAR(50) NOT NULL,
            Name    VARCHAR(50) NOT NULL,
            Region  VARCHAR(50) NOT NULL,
            Image   VARCHAR(150) NOT NULL
        );`,
		`CREATE TABLE Team (
            Name    VARCHAR(50) NOT NULL,
            Code    VARCHAR(50) PRIMARY KEY,
            Image   VARCHAR(150) NOT NULL
        );`,
		`CREATE TABLE Game (
            Id          INT AUTO_INCREMENT PRIMARY KEY,
            Team_1      VARCHAR(50) NOT NULL,
            Team_2      VARCHAR(50) NOT NULL,
            League      INT NOT NULL,
            Time        TIMESTAMP  DEFAULT CURRENT_TIMESTAMP NOT NULL,
            Bets_t1     INT DEFAULT 0 NOT NULL,
            Bets_t2     INT DEFAULT 0 NOT NULL,
            Completed   TINYINT DEFAULT 0 NOT NULL,
            BlockName   VARCHAR(20) NOT NULL,
            Strategy    VARCHAR(50) NOT NULL,
            CHECK ( Completed IN (0, 1, 2) ),
            CHECK ( Team_1 <> Team_2 ),
            FOREIGN KEY (League) REFERENCES League(Id),
            FOREIGN KEY (Team_1) REFERENCES Team(Code),
            FOREIGN KEY (Team_2) REFERENCES Team(Code),
            UNIQUE (Team_1, Time)
        );`,
		`CREATE TABLE Bet (
            Id      INT AUTO_INCREMENT PRIMARY KEY,
            GameId  INT NOT NULL,
            UserId  INT NOT NULL,
            Value   INT NOT NULL,
            Team    TINYINT NOT NULL,
            Odds    DECIMAL(3, 1) UNSIGNED NOT NULL,
            FOREIGN KEY (GameId) REFERENCES Game(Id),
            FOREIGN KEY (UserId) REFERENCES User(Id)
        );`,
	}

	for _, sqlStatement := range createTableSQL {
		_, err := db.Exec(sqlStatement)
		if err != nil {
			fmt.Println("Error creating table:", err)
			return
		}
	}
}

func dropDBtables(db *sql.DB) {
	fmt.Println("Dropping tables...")
	_, err := db.Exec("DROP TABLE IF EXISTS Bet")
	if err != nil {
		fmt.Println("Error dropping Bet table:", err)
		return
	}
	_, err = db.Exec("DROP TABLE IF EXISTS Game")
	if err != nil {
		fmt.Println("Error dropping Game table:", err)
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

func UpdateApiToDatabase(db *sql.DB, apiData ApiSchedule) {
	// Get last already completed game from DB
    // Games before the last completed game are already updated in the db
	lastGameTime, err := models.GetLastCompletedGameTime(db)
	if err != nil {
		//fmt.Println("There si no already completed last game:", err)
        fmt.Fprintf(os.Stderr, "GetLastCompletedGameTime: %s\n", err)
	}
    fmt.Println("Games from DB before last completed game:", lastGameTime)

	// Pillar todos los resultados de la api
    //TODO: Problema: Al limpiar partidos el ultimo que ya esta completado es el de la hora entonces no se limpia y se intenta volver a añadir
	gamesApi, teamsApi := CleanApiData(apiData, lastGameTime)
    fmt.Println("Games from API:", gamesApi)

	// Pillar todos los partidos incompletos de la bd
	unfinishedGames, err := models.GetUnfinishedGames(db)
	if err != nil {
        fmt.Fprintf(os.Stderr, "GetUnfinishedGames: %s\n", err)
	}
    fmt.Println("Unfinished games:", unfinishedGames)

    //TODO: maybe change it for an updateGames list that appends the ones that have to be updated
	// See what games stored in the db are completed to update them
	for i := range unfinishedGames {
        v := &unfinishedGames[i]
		key := v.League + v.Time.String()
        fmt.Println("Key:", key)
		apiGame := gamesApi[key]
        fmt.Println("Api game:", apiGame)
		if apiGame.Completed == 1 { //Game has been played and team 1 won
            // Comprobar si estan en el mismo orden
            if v.Team1 == apiGame.Team1 {
			    v.Completed = 1
            } else {
                v.Completed = 2
            }
		} else if apiGame.Completed == 2 { //Game has been played and team 2 won
            if v.Team1 == apiGame.Team1 {
			    v.Completed = 2
            } else {
                v.Completed = 1
            }
		}
		delete(gamesApi, key)
		// If the game already exists in the db, the teams are already in the db
		delete(teamsApi, v.Team1)
		delete(teamsApi, v.Team2)
	}

    fmt.Println("Unfinished games:", unfinishedGames)
	// Modificar en la bd unfinishedGames
	//go models.UpdateMultipleGames(db, unfinishedGames)
    err = models.UpdateMultipleGames(db, unfinishedGames)
	if err != nil {
		fmt.Fprintf(os.Stderr, "UpdateMultipleGames: %s\n", err)
	}

    fmt.Println("Games from API after updating unfinished games:", gamesApi)
	// Now on apiGames we have only the games that are not in the db yet
	var newGames []models.Game
	for _, game := range gamesApi {
        /*
		if game.Completed > 0 {
			delete(gamesApi, key)
			delete(teamsApi, game.Team1)
			delete(teamsApi, game.Team2)
		} else {
			newGames = append(newGames, game)
		}
        */
        // If the gama hasn't been played yet, add it to the db
        if game.Completed == 0 {
            newGames = append(newGames, game)
        }
	}

	// Add new teams to the db, they might already exist
    fmt.Println("Adding new teams, they might already exist...", teamsApi)
	for _, team := range teamsApi {
		err := models.AddTeam(db, &team)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
	}

	//go models.AddMultipleGames(db, newGames)
	/*
	   err = models.AddMultipleGames(db, newGames)
	   if err != nil {
	       fmt.Fprintf(os.Stderr, "AddMultipleGames: %s\n", err)
	   }
	*/
	fmt.Println("Adding new games...")
	for _, game := range newGames {
		fmt.Println(game)
		err := models.AddGame(db, &game)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
	}
}

func UpdateDatabase(db *sql.DB) {
	// Pillar todos los resultados de la api
	apiData := GetScheduleApi()

	UpdateApiToDatabase(db, apiData)
}

func InitializeDatabase(db *sql.DB) {
	// Drop DB tables
	dropDBtables(db)

	// Create DB tables
	createDBtables(db)

	// Pillar todos las ligas de la api e insertarlas en la bd
	leaguesAPI := GetLeaguesApi()

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
