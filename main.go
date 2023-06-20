package main

import (
	"develov_be/controller"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	server := controller.Server{}

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.InitializeServer(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	server.RunServer(":" + os.Getenv("PORT"))
}
