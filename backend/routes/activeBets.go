package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
    "strconv"

	"github.com/JorgeMG117/LolBets/backend/models"
)

//Return a json that contains a bet and its game
type respose struct {
    Bet models.Bet `json:"bet"`
    Team1 string `json:"team1"`
    Team2 string `json:"team2"`
    League string `json:"league"`
    Completed int `json:"completed"`
}


func (s *Server) ActiveBets(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
        keys, exists := r.URL.Query()["userId"]

        if !exists {
            w.Write([]byte("{\"error\": \"param userId not found\"}"))
            return
        }
        userId, err := strconv.Atoi(keys[0])
        if err != nil {
            w.Write([]byte("{\"error\":\"" + err.Error() + "\"}"))
            return
        }

        // Tiene que devolver las active bets de un usuario
        bets := s.ActiveGames.GetActiveBets(userId)
        //get active games
        games := s.ActiveGames.GetGames("", "")
        //Get the games that appear in the bets
        var response []respose

        for _, bet := range bets {
            for _, game := range games {
                if bet.GameId == game.Id {
                    response = append(response, respose{bet, game.Team1, game.Team2, game.League, game.Completed})
                }
            }
        }

        //TODO Get user recent already completed bets
        rows, err := models.GetBetsOfUser(s.Db, userId)
        if err != nil {
            w.Write([]byte("{\"error\":\"" + err.Error() + "\"}"))
            return
        }
        for rows.Next() {
            fmt.Println("Next")
            //Add to response
            var res respose
            err = rows.Scan(&res.Bet.GameId, &res.Bet.UserId, &res.Bet.Value, &res.Bet.Team, &res.Bet.Odds, &res.Team1, &res.Team2, &res.League, &res.Completed)

            if err != nil {
                w.Write([]byte("{\"error\":\"" + err.Error() + "\"}"))
                return
            }
            response = append(response, res)
        }

        //Return
        jsonResponse, _ := json.Marshal(response)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(jsonResponse)
	}
}
