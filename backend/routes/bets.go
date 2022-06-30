package routes

import (
	"net/http"
	"log"
	"github.com/JorgeMG117/LolBets/backend/models"
	"github.com/gorilla/websocket"
    "strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//TODO
func userBetController(conn *websocket.Conn, chBets chan models.Bet) {
    defer conn.Close()
    for {
        mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
    }
}

//Usuario quiere apostar en partida identificada por el id
//Se crea una conexion websockets
//Cuando el usuario apueste se utilizara el canal asociado a dicho partido
func (s *Server) Bets(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        var gameId string 
        
        keys, exists := r.URL.Query()["game"]

        if !exists {
            w.Write([]byte("{error: param game not found}"))
            return
        }
        gameId = keys[0]
        
        gameIdInt, err := strconv.Atoi(gameId)
        
        if err != nil {
            w.Write([]byte("{error:" + err.Error() + "}"))
            return
        }

        idx := models.GetIdxOfGame(gameIdInt)

        if idx == -1 {
            w.Write([]byte("{error: Game doesnt exists}"))
            return
        }

        conn, err := upgrader.Upgrade(w, r, nil)
        
        if err != nil {
            w.Write([]byte("{error:" + err.Error() + "}"))
            return
        }

        go userBetController(conn, s.ChBets[idx])
    }
}

