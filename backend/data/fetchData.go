package data

import (
	//"database/sql"
	"encoding/json"
	"fmt"

	//"log"
	//"os"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
    "os"

	//"github.com/JorgeMG117/LolBets/backend/configs"
	"github.com/JorgeMG117/LolBets/backend/models"
	//"github.com/joho/godotenv"
)


// Reads an api_schedule.json file and returns the games in it
func ReadApiSchedule(filename string) ApiSchedule {
    // Open our jsonFile
    jsonFile, err := os.Open(filename)
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened api_schedule.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened jsonFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    var apiSchedule ApiSchedule
    err = json.Unmarshal(byteValue, &apiSchedule)
    if err != nil {
        fmt.Println(err)
    }

    return apiSchedule
}

func getApi(url string) []byte {
	headers := map[string]string{
		//"User-Agent":      "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/117.0",
		"Accept":          "*/*",
		"Accept-Language": "en-US,en;q=0.5",
		//"Accept-Encoding": "gzip, deflate, br",
		"Referer":   "https://lolesports.com/",
		"x-api-key": "0TvQnueqKa5mxJntVWt0w4LpLfEkrV1Ta8rQBb9Z",
		"Origin":    "https://lolesports.com",
		"DNT":       "1",
		//"Connection":      "keep-alive",
		"Sec-Fetch-Dest": "empty",
		"Sec-Fetch-Mode": "cors",
		"Sec-Fetch-Site": "same-site",
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
	//fmt.Println(resp)

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	//fmt.Println(string(body))

	return body
}

func GetScheduleApi() ApiSchedule {
	result := getApi("https://esports-api.lolesports.com/persisted/gw/getSchedule?hl=en-US&leagueId=98767975604431411%2C110988878756156222")

	var values ApiSchedule
	fmt.Println("Error: ", json.Unmarshal(result, &values))

	return values
}

func CleanApiData(apiData ApiSchedule, timeFromWhich time.Time) (map[string]models.Game, map[string]models.Team) {
	gamesApi := make(map[string]models.Game)
	teamsApi := make(map[string]models.Team)

	for _, event := range apiData.Data.Schedule.Events {
		// Check if the game is before the timeFromWhich, that means we already have it in the database with all the info
		if event.StartTime.Before(timeFromWhich) {
			continue
		}

		// Check if there is 2 teams
		teams := event.Match.Teams
		if len(teams) != 2 {
			fmt.Println("Error: ", "There is not 2 teams in the game")
		}

		// Check if either of 2 teams name is TBD
		if teams[0].Name == "TBD" || teams[1].Name == "TBD" {
			continue
		}

		// Safe every team just in case the game can't be created
		teamsApi[teams[0].Name] =
			models.Team{
				Name:  teams[0].Name,
				Image: teams[0].Image,
				Code:  teams[0].Code,
			}
		teamsApi[teams[1].Name] =
			models.Team{
				Name:  teams[1].Name,
				Image: teams[1].Image,
				Code:  teams[1].Code,
			}

		// Check what team has won
		gameResult := 0
		completed := event.State == "completed"
		if completed && *teams[0].Result.Outcome == "win" { //First team won
			gameResult = 1
		} else if completed { //Second team won
			gameResult = 2
		}

		// TODO: Cambiar primary key de game a team1:date en db
		gamesApi[teams[0].Name+event.StartTime.String()] =
			models.Game{
				Time:      event.StartTime,
				Team1:     teams[0].Name,
				Team2:     teams[1].Name,
				League:    event.League.Name,
				BlockName: string(event.BlockName),
				Strategy:  "best of " + strconv.FormatInt(event.Match.Strategy.Count, 10),
				Completed: gameResult,
			}
	}

	return gamesApi, teamsApi
}

func GetLeaguesApi() []models.League {
	result := getApi("https://esports-api.lolesports.com/persisted/gw/getLeagues?hl=en-US")

	var values ApiLeague
	fmt.Println("Error: ", json.Unmarshal(result, &values))

	//fmt.Println(values)

	var leagues []models.League

	for _, l := range values.Data.Leagues {
		league := models.League{
			ApiID:  l.ID,
			Name:   l.Name,
			Region: l.Region,
			Image:  l.Image,
		}
		leagues = append(leagues, league)
	}

	return leagues
}
