package models

import (
	"database/sql"
	"fmt"
)

type User struct {
    Id    int    `json:"id"`
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

func GetUserByGoogleId(db *sql.DB, googleId string) (*User, error) {
    var user User
    err := db.QueryRow("SELECT * FROM User WHERE Name = ?", googleId).Scan(&user.Id, &user.Name, &user.Coins)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
