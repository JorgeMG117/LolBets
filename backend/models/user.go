package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Coins int    `json:"coins"`
}

func AddUser(db *sql.DB, newUser *User) error {
	result, err := db.Exec("INSERT INTO User(Name, Coins) VALUES (?, ?)", newUser.Name, newUser.Coins)
	if val, _ := result.RowsAffected(); val != 1 {
		fmt.Println("No se ha insertado nada o se ha insertado mas de un valor")
		fmt.Println(newUser)
	}
	return err
}
