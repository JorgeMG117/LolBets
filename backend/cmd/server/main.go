package main

import (
	"time"
	"github.com/JorgeMG117/LolBets/backend/server"
    "fmt"
    "os"
)

/*
func (s *server) handleGames() http.HandlerFunc {
}

*/


/*func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}
func timeHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
}*/

func main() {

    //Sleeping for 30 seconds to setup everything else
    fmt.Println("Sleeping for 30 seconds to setup everything else")
    time.Sleep(time.Second * 30)

    initializeDB := false 

	args := os.Args[1:]

	if len(args) == 1 {
	    if args[0] == "--initialize" {
            initializeDB = true 
        }
	}
    fmt.Println("Initialize DB: ", initializeDB)

    //Launch server
    if err := server.ExecServer(initializeDB); err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        os.Exit(1)
    }
}
