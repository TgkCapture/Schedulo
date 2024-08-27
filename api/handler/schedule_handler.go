package handler

import (
	"html/template"
	"log"
	"net/http"
)

func ScheduleHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("web/templates/index.html")
    if err != nil {
        log.Fatalf("Error Loading template: %v", err)
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        log.Fatalf("Error executing template: %v", err)
    }
}