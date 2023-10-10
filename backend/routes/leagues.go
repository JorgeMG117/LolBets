package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JorgeMG117/LolBets/backend/models"
)

func (s *Server) Leagues(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
        leagues, err := models.GetLeagues(s.Db)
        if err != nil {
            // Handle the error and return it as a JSON response
            errorResponse := map[string]string{"error": err.Error()}
            jsonResponse, _ := json.Marshal(errorResponse)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusInternalServerError)
            w.Write(jsonResponse)
            return
        }
        jsonResponse, _ := json.Marshal(leagues)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(jsonResponse)

	case "POST":
		//Read body content
		out := make([]byte, 1024)
		bodyLen, err := r.Body.Read(out)

		if err != io.EOF {
			//log.Println(err)
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		var league models.League

		err = json.Unmarshal(out[:bodyLen], &league)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		err = models.AddLeague(s.Db, &league)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		w.Write([]byte(`{"error":"success"}`))
	}
}
