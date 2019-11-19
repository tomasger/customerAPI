package main

import (
	"customerAPI"
	"customerAPI/storage"
	"log"
	"os"
	"strconv"
)

func main() {
	// TODO get env for DB type, port
	db := storage.NewMapStorage()
	app := customerAPI.NewApp(db)
	app.Init()
	portEnv := os.Getenv("CUSTOMER_API_PORT")
	if portEnv == "" {
		app.Run(8080)
	} else {
		port, err := strconv.Atoi(portEnv)
		if err != nil {
			log.Fatal("CUSTOMER_API_PORT is invalid.")
		}
		app.Run(port)
	}
}