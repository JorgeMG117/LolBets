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
