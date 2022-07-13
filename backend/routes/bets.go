package routes

import (
	"net/http"
	"log"
	"github.com/JorgeMG117/LolBets/backend/models"
	"github.com/gorilla/websocket"
    "fmt"
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
        //Check if bets if bet is correct
        correct := true

        if correct {
            val, err := strconv.Atoi(string(message))
            if err != nil {
                log.Println("write:", err)
                break
            }
            chBets <- models.Bet{val, false} 
        } else {
            response := "There was an error in the bet" 
            err = conn.WriteMessage(mt, []byte(response))
            if err != nil {
                log.Println("write:", err)
                break
            }
        }
    }
}

//Usuario quiere apostar en partida identificada por el id
//Se crea una conexion websockets
//Cuando el usuario apueste se utilizara el canal asociado a dicho partido
func (s *Server) Bets(w http.ResponseWriter, r *http.Request) {
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

    fmt.Println("User conected to bet: " + gameId)

    go userBetController(conn, s.ChBets[idx])
}

