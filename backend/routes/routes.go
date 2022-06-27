package routes

import (
	"net/http"

	"database/sql"
)

type Server struct {
    Db      *sql.DB
    //router
}

func (s *Server) Router() http.Handler {
    //th := timeHandler{format: "a"}
    mux := http.NewServeMux() 
    mux.HandleFunc("/games", s.Games)
    return mux
}
