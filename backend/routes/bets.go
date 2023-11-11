package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
    "time"

	"github.com/JorgeMG117/LolBets/backend/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func isValid(msg []byte) (models.Bet, bool) {
	var bet models.Bet
	err := json.Unmarshal(msg, &bet)
	if err != nil {
		log.Println("{error:" + err.Error() + "}")
		return bet, false
	}
    
    // Check if bet fields are correct
    if bet.Value <= 0 {
        log.Println("{\"error\": \"Amount must be greater than 0\"}")
        return bet, false
    }


	return bet, true
}

func updateGameInfo(conn *websocket.Conn, idxGame int) {
    // Lanzar gorutine que actualize la informacion de cada game para el cliente en tiempo real
    // while not out
    //      send updated game info
    //      sleep 1s


    // Necesito acceder a la informacion del game
    // Cuando me hacen el get me pasan el id del game
    // Soluciones:
    // - Hacer un getGame que dado el id me saque los valores de game del slice
	defer conn.Close()

    out := false
    for !out {
        game := models.GetGameInfoByIdx(idxGame)

        // Only return game id, and bets
        gameInfoForClient := struct {
            Id   int         `json:"id"`
            Bets1     int       `json:"bets1"`
            Bets2     int       `json:"bets2"`
        }{
            Id:   game.Id,
            Bets1: game.Bets1,
            Bets2: game.Bets2,
        }

        // Send updated game info
        err := conn.WriteJSON(gameInfoForClient)
        if err != nil {
            log.Println("{error:" + err.Error() + "}")
            return
        }
        // Sleep 1s
        time.Sleep(30 * time.Second)
    }
}

func userBetController(conn *websocket.Conn, chBets chan models.Bet) {
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("{\"error\": \"" + err.Error() + "\"}")
			err = conn.WriteMessage(mt, []byte("{\"error\": \""+err.Error()+"\"}"))
			if err != nil {
				log.Println("{error:" + err.Error() + "}")
				break
			}
		}

		log.Printf("recv: %s", message)

		if bet, ok := isValid(message); ok {

			chBets <- bet
			log.Println(`{"error":"success"}`)

			err = conn.WriteMessage(mt, []byte(`{"error":"success"}`))
			if err != nil {
				log.Println("{error:" + err.Error() + "}")
				break
			}
		} else {
			err = conn.WriteMessage(mt, []byte(`{"error":"bet is not correct"}`))
			if err != nil {
				log.Println("{error:" + err.Error() + "}")
				break
			}
		}
	}
}

// Usuario quiere apostar en partida identificada por el id
// Se crea una conexion websockets
// Cuando el usuario apueste se utilizara el canal asociado a dicho partido
func (s *Server) Bets(w http.ResponseWriter, r *http.Request) {
	var gameId string

	keys, exists := r.URL.Query()["game"]

	if !exists {
		w.Write([]byte("{\"error\": \"param game not found\"}"))
		return
	}
	gameId = keys[0]

	gameIdInt, err := strconv.Atoi(gameId)

	if err != nil {
		w.Write([]byte("{\"error\":\"" + err.Error() + "\"}"))
		return
	}

	idx := models.GetIdxOfGame(gameIdInt)

	if idx == -1 {
		w.Write([]byte("{\"error\": \"Game doesnt exists\"}"))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		w.Write([]byte("{\"error\": \"" + err.Error() + "\"}"))
		return
	}

	fmt.Println("User conected to bet: " + gameId)

	go userBetController(conn, s.ChBets[idx])
    go updateGameInfo(conn, idx)
}
