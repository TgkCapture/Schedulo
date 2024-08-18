package main

import (
	"log"
	"net/http"

	"github.com/TgkCapture/Schedulo/api/handler"
	"github.com/TgkCapture/Schedulo/config"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/", handler.ScheduleHandler)

	log.Println("Starting Schedulo Server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}