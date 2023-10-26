package main

import (
	//"time"
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
    //Launch server
    if err := server.ExecServer(); err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        os.Exit(1)
    }
}
