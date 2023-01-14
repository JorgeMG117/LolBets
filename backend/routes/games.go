package routes

import (
	"fmt"
	"io"
	"net/http"

	"encoding/json"
	//"log"

	"github.com/JorgeMG117/LolBets/backend/models"
)

func (s *Server) Games(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		var league string
		var team string

		keys, exists := r.URL.Query()["league"]

		if exists {
			league = keys[0]
		} else {
			league = ""
		}

		keys, exists = r.URL.Query()["team"]

		if exists {
			team = keys[0]
		} else {
			team = ""
		}

		games, err := models.GetGames(s.Db, league, team)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		//TODO
		fmt.Println(games)
	case "POST":
		//Read body content
		out := make([]byte, 1024)
		bodyLen, err := r.Body.Read(out)

		if err != io.EOF {
			//log.Println(err)
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		var game models.Game

		err = json.Unmarshal(out[:bodyLen], &game)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		err = models.AddGame(s.Db, &game)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		w.Write([]byte(`{"error":"success"}`))
	}
}
