package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	//"time"

	"github.com/JorgeMG117/LolBets/backend/configs"
	"github.com/JorgeMG117/LolBets/backend/data"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  --update")
	fmt.Println("  --initialize")
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	args := os.Args[1:]

	db := configs.ConnectDB()
	defer db.Close()

	if len(args) == 0 {
		printUsage()
	} else if args[0] == "--update" {
		data.UpdateDatabase(db)
	} else if args[0] == "--initialize" {
		data.InitializeDatabase(db)
	} else {
		printUsage()
	}

}
