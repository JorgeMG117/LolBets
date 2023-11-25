package routes

import (
	"net/http"

	"database/sql"

	datastructures "github.com/JorgeMG117/LolBets/backend/data_structures"
)

type Server struct {
	Db     *sql.DB
    ActiveGames *datastructures.ActiveGames
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
	mux.HandleFunc("/activeBets", s.ActiveBets)
	return mux
}
