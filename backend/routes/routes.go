package routes

import (
	"net/http"

	"database/sql"

	"github.com/JorgeMG117/LolBets/backend/models"
)

type Server struct {
	Db     *sql.DB
	ChBets []chan models.Bet
	//router
}

func (s *Server) Router() http.Handler {
	//th := timeHandler{format: "a"}
	mux := http.NewServeMux()
	mux.HandleFunc("/games", s.Games)
	mux.HandleFunc("/bets", s.Bets)
	mux.HandleFunc("/leagues", s.Leagues)
	mux.HandleFunc("/teams", s.Teams)
	mux.HandleFunc("/users", s.Users)
	return mux
}
