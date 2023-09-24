package main

import (
	"archive-bot/cmd"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.SetupBot()
	if err != nil {
		log.Fatal(err)
	}
}
