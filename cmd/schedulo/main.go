package main

import (
	"log"
	"net/http"

	"github.com/TgkCapture/Schedulo/db"
	"github.com/TgkCapture/Schedulo/api/handler"
	"github.com/TgkCapture/Schedulo/config"
)

func main() {
	config.LoadConfig()

	err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize the databe: %v\n", err)
	}
	defer db.CloseDB()

	err = db.CreateTables()
	if err != nil {
		log.Fatalf("Could not create tables: %v\n", err)
	}

	http.HandleFunc("/", handler.ScheduleHandler)

	log.Printf("Starting ScheduloGo server on :%s\n", config.Cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.Cfg.ServerPort, nil))
}