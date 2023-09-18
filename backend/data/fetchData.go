package data

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	//"log"
	//"os"
	"time"
    "net/http"
    "io/ioutil"

	//"github.com/JorgeMG117/LolBets/backend/configs"
	"github.com/JorgeMG117/LolBets/backend/models"
	//"github.com/joho/godotenv"
)


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
    //fmt.Println(resp)

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
    //fmt.Println(string(body))

    return body
}


// Gets the schedule from the API
// timeFromWhich is the time from which we want to get the games
// Returns a map where key of games is team1:date
// Return a map of games and a map of teams
func GetScheduleApi(timeFromWhich time.Time) (map[string]models.Game, map[string]models.Team) {
    result := getApi("https://esports-api.lolesports.com/persisted/gw/getSchedule?hl=en-US&leagueId=98767975604431411%2C110988878756156222")

	var values ApiSchedule
	fmt.Println("Error: ", json.Unmarshal(result, &values))

    //fmt.Println(values)

    gamesApi := make(map[string]models.Game)
    teamsApi := make(map[string]models.Team)

    for _, event := range values.Data.Schedule.Events {
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
            models.Team {
                Name: teams[0].Name,
                Image: teams[0].Image,
                Code: teams[0].Code,
            }
        teamsApi[teams[1].Name] = 
            models.Team {
                Name: teams[1].Name,
                Image: teams[1].Image,
                Code: teams[1].Code,
            }


        // Check what team has won
        gameResult := 0
        completed := event.State == "completed"
        if completed && *teams[0].Result.Outcome == "win" {//First team won
            gameResult = 1
        } else if completed {//Second team won
            gameResult = 2
        }


        gamesApi[teams[0].Name+event.StartTime.String()] = 
            models.Game {
                Time: event.StartTime,
                Team1: teams[0].Name,
                Team2: teams[1].Name,
                League: event.League.Name,
                BlockName: "best of " + string(event.Match.Strategy.Count),
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
        league := models.League {
            ApiID: l.ID,
            Name: l.Name,
            Region: l.Region,
            Image: l.Image,
        }
        leagues = append(leagues, league)
	}

    return leagues
}

// func mapLeagueApi(leagueApi) { models.League {
// func getLeaguesApi() []models.League {

//func getScheduleApi(db *sql.DB, leagues []string) map[string]game {


/*
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
*/

/*
func getLeaguesApi() []models.League {
	data := getApi("leagues.json")

	var values LeaguesData
	fmt.Println("Error: ", json.Unmarshal(data, &values))

	return values.LeaguesData.Leagues
}
*/

/*
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
*/

