package server

import (
	"log"
	"net/http"

	//"time"

	"github.com/JorgeMG117/LolBets/backend/configs"
	"github.com/JorgeMG117/LolBets/backend/models"
	"github.com/JorgeMG117/LolBets/backend/routes"
    "github.com/JorgeMG117/LolBets/backend/data"

	"fmt"
	"os"
    "time"   

	"github.com/joho/godotenv"
)

func ExecServer() error {
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


    //Launch update games program
    go func() {
        for {
	        data.UpdateDatabase(s.Db)
            time.Sleep(time.Hour)
        }
    }()

    fmt.Println("Waiting for database to update")
    time.Sleep(time.Second * 30)


	//setup thigs
	serv := &http.Server{
		Addr:    ":8080",
		Handler: s.Router(),
		//ReadTimeout:    10 * time.Second,
		//WriteTimeout:   10 * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}

    chUpdateGames := make(chan int, models.MaxGames)

    err := models.InitializeGames(s.Db, chUpdateGames)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	chBets := make([]chan models.Bet, models.MaxGames)

	for i := 0; i < models.NumGames(); i++ {
		chBets[i] = make(chan models.Bet)
		go models.BetController(chBets[i], i, chUpdateGames)
	}
	s.ChBets = chBets

	log.Fatal(serv.ListenAndServe())

	return nil
}
