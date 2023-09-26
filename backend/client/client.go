package client

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/JorgeMG117/LolBets/backend/models"
	"github.com/gorilla/websocket"
)

// Send the message of the bet
func makeBet(conn *websocket.Conn, bet models.Bet) {

	b, _ := json.Marshal(&bet)

    err := conn.WriteMessage(websocket.TextMessage, []byte(string(b)))
	if err != nil {
		log.Println("write:", err)
		return
	}
}

func Client(chStop chan bool, chBet chan models.Bet) {
    gameId := "1"
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/bets", RawQuery: "game=" + gameId}
	log.Printf("connecting to %s", u.String())

	c, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Printf("handshake failed with status %d", resp.StatusCode)
		log.Fatal("dial:", err)
	}

	//When the program closes close the connection
	defer c.Close()

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("{error:" + err.Error() + "}")

			}
			log.Printf("recv: %s", message)

		}
	}()


    for {
        select {
        case <-chStop:
            return
        case b := <-chBet:
            makeBet(c, b)
        //case return games
        }

    }



	/*
	    messageOut := make(chan string)
	    interrupt := make(chan os.Signal, 1)
	    signal.Notify(interrupt, os.Interrupt)
	    u := url.URL{Scheme: "wss", Host: "marketdata.tradermade.com", Path: "/feedadv",}
	    log.Printf("connecting to %s", u.String())
	    c, resp, err := websocket.DefaultDialer.Dial(u.String(), nil);
	    if err != nil {
	        log.Printf("handshake failed with status %d", resp.StatusCode)
	        log.Fatal("dial:", err)
	    }

	    //When the program closes close the connection
	    defer c.Close()
	    done := make(chan struct{})
	    go func() {
	    defer close(done)
	    for {
	        _, message, err := c.ReadMessage()
	        if err != nil {
	            log.Println("read:", err)
	            return
	        }
	        log.Printf("recv: %s", message)
	        if string(message) == "Connected"{
	            log.Printf("Send Sub Details: %s", message)
	            messageOut <- "{"userKey":"YOUR_API_KEY", "symbol":"EURUSD"}"
	        }
	    }

	    }()

	    ticker := time.NewTicker(time.Second)
	    defer ticker.Stop()
	  for {
	    select {
	    case <-done:
	      return
	    case m := <-messageOut:
	      log.Printf("Send Message %s", m)
	      err := c.WriteMessage(websocket.TextMessage, []byte(m))
	      if err != nil {
	        log.Println("write:", err)
	        return
	      }
	    case t := <-ticker.C:
	      err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
	      if err != nil {
	        log.Println("write:", err)
	        return
	      }
	    case <-interrupt:
	      log.Println("interrupt")
	      // Cleanly close the connection by sending a close message and then
	      // waiting (with timeout) for the server to close the connection.
	      err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	      if err != nil {
	        log.Println("write close:", err)
	        return
	      }
	      select {
	      case <-done:
	      case <-time.After(time.Second):
	      }
	      return
	    }
	  }
	*/
}
