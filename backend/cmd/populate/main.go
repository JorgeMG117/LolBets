package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
    //"time"

	"github.com/JorgeMG117/LolBets/backend/models"
    "github.com/JorgeMG117/LolBets/backend/data"
    "github.com/JorgeMG117/LolBets/backend/configs"
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

func UpdateApiToDatabase(db *sql.DB, apiData data.ApiSchedule) {
    // Get last already completed game from DB
    lastGameTime, err := models.GetLastCompletedGameTime(db)
    if err != nil {
        //fmt.Println("There si no already completed last game:", err)
		fmt.Fprintf(os.Stderr, "%s\n", err)
    }

	// Pillar todos los resultados de la api
    gamesApi, teamsApi := data.CleanApiData(apiData, lastGameTime)

	// Pillar todos los partidos incompletos de la bd
	unfinishedGames, err := models.GetUnfinishedGames(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	fmt.Println(unfinishedGames)

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
        // If the game already exists in the db, the teams are already in the db
        delete(teamsApi, v.Team1)
        delete(teamsApi, v.Team2)
	}

	// Modificar en la bd unfinishedGames
    //go models.UpdateMultipleGames(db, unfinishedGames)
    err = models.UpdateMultipleGames(db, unfinishedGames)
    if err != nil {
        fmt.Fprintf(os.Stderr, "UpdateMultipleGames: %s\n", err)
    }

    // Now on apiGames we have only the games that are not in the db yet
    var newGames []models.Game
	for key, game := range gamesApi {
		if game.Completed > 0 {
			delete(gamesApi, key)
            delete(teamsApi, game.Team1)
            delete(teamsApi, game.Team2)
		} else {
            newGames = append(newGames, game)
		}
	}

    // Add new teams to the db, they might already exist
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

func UpdateDatabase2(db *sql.DB) {
	// Pillar todos los resultados de la api
    apiData := data.GetScheduleApi()

    UpdateApiToDatabase(db, apiData)
}


func UpdateDatabase(db *sql.DB) {
    // Get last already completed game from DB
    lastGameTime, err := models.GetLastCompletedGameTime(db)
    if err != nil {
        //fmt.Println("There si no already completed last game:", err)
		fmt.Fprintf(os.Stderr, "%s\n", err)
    }

	// Pillar todos los resultados de la api
    gamesApi, teamsApi := data.GetCleanScheduleApi(lastGameTime)
	fmt.Println(gamesApi)

	// Pillar todos los partidos incompletos de la bd
	unfinishedGames, err := models.GetUnfinishedGames(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	fmt.Println(unfinishedGames)

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
        // If the game already exists in the db, the teams are already in the db
        delete(teamsApi, v.Team1)
        delete(teamsApi, v.Team2)
	}

	// Modificar en la bd unfinishedGames
    //go models.UpdateMultipleGames(db, unfinishedGames)
    err = models.UpdateMultipleGames(db, unfinishedGames)
    if err != nil {
        fmt.Fprintf(os.Stderr, "UpdateMultipleGames: %s\n", err)
    }

    // Now on apiGames we have only the games that are not in the db yet
    var newGames []models.Game
	for key, game := range gamesApi {
		if game.Completed > 0 {
			delete(gamesApi, key)
            delete(teamsApi, game.Team1)
            delete(teamsApi, game.Team2)
		} else {
            newGames = append(newGames, game)
		}
	}

    // Add new teams to the db, they might already exist
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
	} else {
        printUsage()
    }

}
