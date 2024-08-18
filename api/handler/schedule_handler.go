package handler

import (
    "fmt"
    "net/http"
)

func ScheduleHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to Schedulo scheduling system!")
}