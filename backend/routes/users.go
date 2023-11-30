package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JorgeMG117/LolBets/backend/models"
)

const (
    //Initial coins for a new user
    INITIAL_COINS = 100
)

type requestBody struct {
	GoogleId string `json:"googleId"`
}

func (s *Server) Users(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		fmt.Println("GET /users")
	case "POST":
		out := make([]byte, 1024)
		bodyLen, err := r.Body.Read(out)

		if err != io.EOF {
			//log.Println(err)
			w.Write([]byte("{error:" + err.Error() + "}"))
			return
		}

        //The user send the google id, a string, that is what we will use as name
        //Extract the google id from the body
        //Respond in well formated json
        var googleId requestBody 
        err = json.Unmarshal(out[:bodyLen], &googleId)

        if err != nil {
            w.Write([]byte("{\"error\":" + err.Error() + "}"))
            return
        }


        //Check if the user is already in the database
        user, err := models.GetUserByGoogleId(s.Db, googleId.GoogleId)

        if err != nil {
            //If the user is not in the database, add it
            newUser := models.User{Id: 0, Name: googleId.GoogleId, Coins: INITIAL_COINS}
            err = models.AddUser(s.Db, &newUser)
            if err != nil {
                w.Write([]byte("{\"error\":" + err.Error() + "}"))
                return
            }
            //Respond with the user
            user, err = models.GetUserByGoogleId(s.Db, googleId.GoogleId)

            if err != nil {
                w.Write([]byte("{\"error\":" + err.Error() + "}"))
                return
            }   

            if user == nil {
                w.Write([]byte("{\"error\":\"User not found\"}"))
                return
            }

        } 


        userJson, err := json.Marshal(user)

        if err != nil {
            w.Write([]byte("{\"error\":" + err.Error() + "}"))
            return
        }


        w.Write(userJson)


        /*
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
        */
	}
}
