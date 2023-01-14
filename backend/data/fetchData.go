package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JorgeMG117/LolBets/backend/configs"
	"github.com/JorgeMG117/LolBets/backend/models"
	"github.com/joho/godotenv"
)

type Data struct {
	Data schedule `json:"data"`
}

type schedule struct {
	Schedule events `json:"schedule"`
}

type events struct {
	Events []game `json:"events"`
}

type match struct {
	Teams [2]models.Team `json:"teams"`
}

type game struct {
	Id int `json:"id"`

	StartTime time.Time     `json:"startTime"`
	BlockName string        `json:"blockName"`
	State     string        `json:"state"`
	Type      string        `json:"type"`
	Match     match         `json:"match"`
	League    models.League `json:"league"`
}

func getApi(url string) []byte {
	// // url := "https://league-of-legends-esports.p.rapidapi.com/schedule"

	// req, _ := http.NewRequest("GET", url, nil)

	// req.Header.Add("X-RapidAPI-Key", os.Getenv("APIKEY"))
	// req.Header.Add("X-RapidAPI-Host", "league-of-legends-esports.p.rapidapi.com")

	// res, _ := http.DefaultClient.Do(req)

	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// isValid := json.Valid(body)
	// if !isValid {
	// 	fmt.Println("Error on the JSON returned by the API")
	// }

	// return body

	data, err := os.ReadFile("lec-schedule.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	isValid := json.Valid(data)
	fmt.Println(isValid)
	fmt.Println(string(data))

	return data
}

// Gets the schedule from the API
// Returns a map where key of games is team1:date
func getScheduleApi(db *sql.DB) map[string]game {
	// Cojer todas las ligas de nuestra bd
	// Quitar de los partidos de la api aquellos que no sean de las ligas que nos interesan
	leagues, err := models.GetLeaguesName(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	data := getApi("https://league-of-legends-esports.p.rapidapi.com/schedule")

	var values Data
	fmt.Println("Error: ", json.Unmarshal(data, &values))

	// Slice of games
	scheduleS := values.Data.Schedule.Events

	scheduleM := make(map[string]game)

	for _, v := range scheduleS {
		for _, l := range leagues {
			if v.League.Name == l {
				scheduleM[v.Match.Teams[0].Name+v.StartTime.String()] = v
				break
			}
		}
	}

	return scheduleM
}

func UpdateDatabase() {
	db := configs.ConnectDB()

	// Pillar todos los resultados de la api
	gamesAPI := getScheduleApi(db)

	// Pillar todos los partidos incompletos de la bd
	unfinishedGames := models.GetUnfinishedGames(db)

	db.Close()

	for _, v := range unfinishedGames {
		key := v.Team1 + v.Time
		apiGame := gamesAPI[key]
		if apiGame.State == "completed" {
			//Change unfinishedGames
			fmt.Println("Change unfinishedGames")
		}
		delete(gamesAPI, key)
	}

	//go Modificar en la bd unfinishedGames

	for key, game := range gamesAPI {
		if game.State == "completed" {
			delete(gamesAPI, key)
		} else {
			//Añadir en la bd
			fmt.Println("Añadir en la bd")
		}
	}

	// Recorriendo partidos de la bd
	// Encontrar el correspondiente en la llamada a la api
	// Si APIcompleted y BDcompleted no hacemos nada
	//
	// Ir eliminando de la api los que vas recorrienod

	//Quitar de la api el resto de completed
	//Añadir lo que queda en los de la api (uncompleted a la bd)

}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	UpdateDatabase()
}
