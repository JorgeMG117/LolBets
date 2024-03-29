package models

import (
	"database/sql"
	"fmt"
)

type Team struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Image string `json:"image"`
}

func AddTeam(db *sql.DB, newTeam *Team) error {
	result, err := db.Exec("INSERT INTO Team(Name, Code, Image) VALUES (?, ?, ?)", newTeam.Name, newTeam.Code, newTeam.Image)

	if err != nil {
		return err
	}

	if val, _ := result.RowsAffected(); val != 1 {
		fmt.Println("No se ha insertado nada o se ha insertado mas de un valor")
		fmt.Println(newTeam)
	}
	return err
}

func GetTeams(db *sql.DB, teamsReq []string) ([]Team, error) {
    sql := "SELECT Name, Code, Image FROM Team"

    if len(teamsReq) > 0 {
        sql += " WHERE Name IN ("
        for i := 0; i < len(teamsReq); i++ {
            sql += "'" + teamsReq[i] + "'"
            if i < len(teamsReq)-1 {
                sql += ","
            }
        }
        sql += ")"
    }

    
    rows, err := db.Query(sql)

    if err != nil {
        return nil, err
    }

    var teams []Team

    for rows.Next() {
        var team Team
        err = rows.Scan(&team.Name, &team.Code, &team.Image)
        if err != nil {
            return nil, err
        }
        teams = append(teams, team)
    }
        
	return teams, nil
}
