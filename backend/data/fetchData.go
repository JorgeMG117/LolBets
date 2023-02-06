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

type LeaguesData struct {
	LeaguesData leagues `json:"data"`
}

type leagues struct {
	Leagues []models.League `json:"leagues"`
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

	data, err := os.ReadFile(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	isValid := json.Valid(data)
	if !isValid {
		fmt.Println("Data not valid")
	}
	// fmt.Println(isValid)
	// fmt.Println(string(data))

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
	//fmt.Println(leagues)

	data := getApi("lec-schedule.json")

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

func getLeaguesApi() []models.League {
	data := getApi("leagues.json")

	var values LeaguesData
	fmt.Println("Error: ", json.Unmarshal(data, &values))

	return values.LeaguesData.Leagues
}

// https://league-of-legends-esports.p.rapidapi.com/teams
// There is no teams api, so I'm gonna get the schedule of each league on the database and add the teams of first week of competition
// Actualmente esta hecho para probar otras cosas
// TODO
func getTeamsApi(leagues []models.League) []models.Team {
	teams := make([]models.Team, 0)
	data := getApi("lec-schedule.json")

	var values Data
	fmt.Println("Error: ", json.Unmarshal(data, &values))

	// Slice of games
	schedule := values.Data.Schedule.Events

	i := 0
	for _, game := range schedule {
		if game.BlockName == "Week 1" {
			teams = append(teams, game.Match.Teams[0])
			teams = append(teams, game.Match.Teams[1])
			i = i + 1
			if i == 5 {
				break
			}
		}
	}
	return teams
}

func UpdateDatabase() {
	db := configs.ConnectDB()

	// Pillar todos los resultados de la api
	gamesAPI := getScheduleApi(db)
	//fmt.Println(gamesAPI)

	// Pillar todos los partidos incompletos de la bd
	unfinishedGames, err := models.GetUnfinishedGames(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	//fmt.Println(unfinishedGames)

	db.Close()

	// See what games stored in the db are completed to update them
	for _, v := range unfinishedGames {
		key := v.Team1 + v.Time.String()
		apiGame := gamesAPI[key]
		if apiGame.State == "completed" {
			//Change unfinishedGames
			fmt.Println("Change unfinishedGames")
		}
		delete(gamesAPI, key)
	}

	// //go Modificar en la bd unfinishedGames

	for key, game := range gamesAPI {
		if game.State == "completed" {
			delete(gamesAPI, key)
		} else {
			models.AddGame(db, &game)
		}
	}

	// // Recorriendo partidos de la bd
	// // Encontrar el correspondiente en la llamada a la api
	// // Si APIcompleted y BDcompleted no hacemos nada
	// //
	// // Ir eliminando de la api los que vas recorrienod

	// //Quitar de la api el resto de completed
	// //Añadir lo que queda en los de la api (uncompleted a la bd)

}

func InitializeDatabase() {
	db := configs.ConnectDB()

	// Pillar todos las ligas de la api
	leaguesAPI := getLeaguesApi()

	for _, league := range leaguesAPI {
		err := models.AddLeague(db, &league)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}

	// Pillar todos los equipos
	teamsAPI := getTeamsApi(leaguesAPI)
	fmt.Println(teamsAPI)

	for _, team := range teamsAPI {
		err := models.AddTeam(db, &team)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}

	db.Close()
}

// TODO
func printUsage() {
	fmt.Println("This display the usage of the populate database program")
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	args := os.Args[1:]

	if len(args) == 0 {
		printUsage()
	} else if args[0] == "--update" {
		UpdateDatabase()
	} else if args[0] == "--initialize" {
		InitializeDatabase()
	}

}
