package data

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	//"log"
	"os"
	//"time"
    "net/http"
    "io/ioutil"

	//"github.com/JorgeMG117/LolBets/backend/configs"
	//"github.com/JorgeMG117/LolBets/backend/models"
	//"github.com/joho/godotenv"
)



func fetchDataAndSaveToFile(url string, headers map[string]string, filename string) error {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set the request headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the response status code is not in the 200 range (e.g., 404 Not Found)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}
    fmt.Println(resp)

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
    fmt.Println(string(body))

	var values Welcome
	fmt.Println("Error: ", json.Unmarshal(body, &values))

    fmt.Println(values)

	// Create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the response body to the file
	_, err = file.Write(body)
	if err != nil {
		return err
	}

	fmt.Printf("Data saved to %s\n", filename)
	return nil
}

//curl 'https://esports-api.lolesports.com/persisted/gw/getLeagues?hl=es-ES' --compressed 
//-H 'TE: trailers'


func Prueba() {
    url := "https://esports-api.lolesports.com/persisted/gw/getSchedule?hl=es-ES&leagueId=98767975604431411%2C110988878756156222"
	filename := "output.json"

	headers := map[string]string{
		//"User-Agent":      "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/117.0",
		"Accept":          "*/*",
		"Accept-Language": "en-US,en;q=0.5",
		//"Accept-Encoding": "gzip, deflate, br",
		"Referer":         "https://lolesports.com/",
		"x-api-key":       "0TvQnueqKa5mxJntVWt0w4LpLfEkrV1Ta8rQBb9Z",
		"Origin":          "https://lolesports.com",
		"DNT":             "1",
		//"Connection":      "keep-alive",
		"Sec-Fetch-Dest":  "empty",
		"Sec-Fetch-Mode":  "cors",
		"Sec-Fetch-Site":  "same-site",
	}

    err := fetchDataAndSaveToFile(url, headers, filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func getApi(url string) []byte {
	headers := map[string]string{
		//"User-Agent":      "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/117.0",
		"Accept":          "*/*",
		"Accept-Language": "en-US,en;q=0.5",
		//"Accept-Encoding": "gzip, deflate, br",
		"Referer":         "https://lolesports.com/",
		"x-api-key":       "0TvQnueqKa5mxJntVWt0w4LpLfEkrV1Ta8rQBb9Z",
		"Origin":          "https://lolesports.com",
		"DNT":             "1",
		//"Connection":      "keep-alive",
		"Sec-Fetch-Dest":  "empty",
		"Sec-Fetch-Mode":  "cors",
		"Sec-Fetch-Site":  "same-site",
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Set the request headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is not in the 200 range (e.g., 404 Not Found)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Printf("HTTP request failed with status code: %d", resp.StatusCode)
	}
    fmt.Println(resp)

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
    fmt.Println(string(body))

}

// func mapLeagueApi(leagueApi) { models.League {
// func getLeaguesApi() []models.League {

//func getScheduleApi(db *sql.DB, leagues []string) map[string]game {

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

