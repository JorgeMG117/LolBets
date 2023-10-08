package server

import (
	"log"
	"net/http"

	//"time"

	"github.com/JorgeMG117/LolBets/backend/configs"
	"github.com/JorgeMG117/LolBets/backend/models"
	"github.com/JorgeMG117/LolBets/backend/routes"

	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ExecServer() error {
	//mux := http.NewServeMux()
	//mux.Handle("/", getRoot)

    if os.Getenv("GITHUB_ACTIONS") == "true" {
    } else {
        // Load environment variables from .env file for local development
        if err := godotenv.Load(".env"); err != nil {
            log.Fatalf("Error loading .env file: %s", err)
        }
    }

	s := routes.Server{
		Db: configs.ConnectDB(),
	}
	defer s.Db.Close()

	//setup thigs
	serv := &http.Server{
		Addr:    ":8080",
		Handler: s.Router(),
		//ReadTimeout:    10 * time.Second,
		//WriteTimeout:   10 * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}

    err := models.InitializeGames(s.Db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	chBets := make([]chan models.Bet, models.MaxGames)

	for i := 0; i < models.NumGames(); i++ {
		chBets[i] = make(chan models.Bet)
		go models.BetController(chBets[i], i)
	}
	s.ChBets = chBets

	log.Fatal(serv.ListenAndServe())

	return nil
}
