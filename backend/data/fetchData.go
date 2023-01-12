package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/JorgeMG117/LolBets/backend/models"
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

type league struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
}

type match struct {
	Teams [2]team `json:"teams"`
}

type team struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Image string `json:"image"`
}

type game struct {
	Id int `json:"id"`

	StartTime time.Time `json:"startTime"`
	BlockName string    `json:"blockName"`
	State     string    `json:"state"`
	Type      string    `json:"type"`
	Match     match     `json:"match"`
	League    league    `json:"league"`
}

// Gets the schedule from the API
// Returns a map where key of games is team1:date
func getScheduleApi() map[string]game {
	// Cojer todas las ligas de nuestra bd
	// Quitar de los partidos de la api aquellos que no sean de las ligas que nos interesan
	leagues := models.GetLeaguesName()

	data, err := os.ReadFile("lvp-schedule.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	isValid := json.Valid(data)
	fmt.Println(isValid)
	fmt.Println(string(data))

	var values Data
	fmt.Println("Error: ", json.Unmarshal(data, &values))

	// Slice of games
	scheduleS := values.Data.Schedule.Events

	var scheduleM map[string]game

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

	// Pillar todos los resultados de la api
	gamesAPI := getScheduleApi()

	// Pillar todos los partidos incompletos de la bd
	unfinishedGames := models.GetUnfinishedGames()

	for _, v := range unfinishedGames {
		key := v.Team1 + v.Time
		apiGame := gamesAPI[key]
		if apiGame.State == "completed" {
			//Change unfinishedGames
		}
		delete(gamesAPI, key)
	}

	//go Modificar en la bd unfinishedGames

	for key, game := range gamesAPI {
		if game.State == "completed" {
			delete(gamesAPI, key)
		} else {
			//Añadir en la bd
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

	// url := "https://league-of-legends-esports.p.rapidapi.com/teams?id=lng-esports"

	// req, _ := http.NewRequest("GET", url, nil)

	// req.Header.Add("X-RapidAPI-Key", "a91d24051cmsh1b3a4c8bbd5a183p1ade67jsn22bb71206e2c")
	// req.Header.Add("X-RapidAPI-Host", "league-of-legends-esports.p.rapidapi.com")

	// res, _ := http.DefaultClient.Do(req)

	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	//data := []byte(`{"name":"John", "age":30, "car":null}`)
	data, err := os.ReadFile("lvp-schedule.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	isValid := json.Valid(data)
	fmt.Println(isValid)
	fmt.Println(string(data))

	var game Data
	fmt.Println("Error: ", json.Unmarshal(data, &game))

	fmt.Printf("json: %v\n", game)

	// err := os.WriteFile("p1", body, 0644)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err)
	// }
}
