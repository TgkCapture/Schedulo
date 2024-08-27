package main

import (
	"log"
	"net/http"

	"github.com/TgkCapture/Schedulo/api/handler"
	"github.com/TgkCapture/Schedulo/config"
	"github.com/TgkCapture/Schedulo/db"
	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	r.HandleFunc("/", handler.ScheduleHandler).Methods("GET")
	r.HandleFunc("/schedules", handler.AddScheduleHandler).Methods("POST")
	r.HandleFunc("/schedules", handler.GetSchedulesHandler).Methods("GET")
	r.HandleFunc("/schedules/{id}", handler.DeleteScheduleHandler).Methods("DELETE")
	r.HandleFunc("/add-schedule", handler.AddScheduleFormHandler).Methods("GET")
	r.HandleFunc("/modify-schedule/{id}", handler.ModifyScheduleFormHandler).Methods("GET") 
	r.HandleFunc("/delete-schedule/{id}", handler.DeleteScheduleFormHandler).Methods("GET")

	log.Printf("Starting ScheduloGo server on :%s\n", config.Cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.Cfg.ServerPort, r))
}