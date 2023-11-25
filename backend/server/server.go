package server

import (
	"log"
	"net/http"

	//"time"

	"github.com/JorgeMG117/LolBets/backend/configs"
	"github.com/JorgeMG117/LolBets/backend/data"
    "github.com/JorgeMG117/LolBets/backend/routes"
	datastructures "github.com/JorgeMG117/LolBets/backend/data_structures"
	//"github.com/JorgeMG117/LolBets/backend/models"

	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func ExecServer(intializeDB bool) error {
	//mux := http.NewServeMux()
	//mux.Handle("/", getRoot)
    //Check if CI env variable is set
    if os.Getenv("CI") == "true" {
        fmt.Println("Running on Github Actions")
    } else {
        // Load environment variables from .env file for local development
        if err := godotenv.Load(".env"); err != nil {
            log.Fatalf("Error loading .env file for server: %s %s %s afks", err, os.Getenv("CI"), os.Getenv("DBUSER"))
        }
    }

	s := routes.Server{
		Db: configs.ConnectDB(),
	}
	defer s.Db.Close()

    if intializeDB {
        go func(){
            data.InitializeDatabase(s.Db)
            for {
                time.Sleep(time.Hour)
                data.UpdateDatabase(s.Db)
            }
        }()
    } else {
        //Launch update games program
        go func() {
            for {
                data.UpdateDatabase(s.Db)
                time.Sleep(time.Hour)
            }
        }()
    }

    fmt.Println("Waiting for database to update")
    time.Sleep(time.Second * 30)


    fmt.Println("### CREATING SERVER ###")

	//setup thigs
	serv := &http.Server{
		Addr:    ":8080",
		Handler: s.Router(),
		//ReadTimeout:    10 * time.Second,
		//WriteTimeout:   10 * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}

    ag, err := datastructures.InitializeActiveGames(s.Db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
    s.ActiveGames = ag


	log.Fatal(serv.ListenAndServe())

	return nil
}
