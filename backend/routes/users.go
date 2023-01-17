package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JorgeMG117/LolBets/backend/models"
)

func (s *Server) Users(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		fmt.Println("GET /users")
	case "POST":
		//Read body content
		out := make([]byte, 1024)
		bodyLen, err := r.Body.Read(out)

		if err != io.EOF {
			//log.Println(err)
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		var user models.User

		err = json.Unmarshal(out[:bodyLen], &user)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		err = models.AddUser(s.Db, &user)

		if err != nil {
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

		w.Write([]byte(`{"error":"success"}`))
	}
}
