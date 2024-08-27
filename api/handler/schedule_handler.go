package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/TgkCapture/Schedulo/models"
	"github.com/TgkCapture/Schedulo/service"
	"github.com/gorilla/mux"
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

// AddScheduleHandler handles adding a new schedule.
func AddScheduleHandler(w http.ResponseWriter, r *http.Request) {
    var schedule models.Schedule
    if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := service.AddSchedule(schedule); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Schedule added successfully"})
}

// GetSchedulesHandler handles fetching all schedules.
func GetSchedulesHandler(w http.ResponseWriter, r *http.Request) {
    schedules, err := service.GetSchedules()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(schedules)
}

// DeleteScheduleHandler handles deleting a schedule by ID.
func DeleteScheduleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err != nil {
        http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
        return
    }

    if err := service.DeleteSchedule(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Schedule deleted successfully"})
}