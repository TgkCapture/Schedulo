package handler

import (
    "github.com/TgkCapture/Schedulo/service"
    "html/template"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// AddScheduleFormHandler renders the add schedule form.
func AddScheduleFormHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("web/templates/add_schedule.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

// ModifyScheduleFormHandler renders the modify schedule form.
func ModifyScheduleFormHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err != nil {
        http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
        return
    }

    schedule, err := service.GetScheduleByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles("web/templates/modify_schedule.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, schedule)
}

// DeleteScheduleFormHandler renders the delete schedule form.
func DeleteScheduleFormHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err != nil {
        http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
        return
    }

    schedule, err := service.GetScheduleByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles("web/templates/delete_schedule.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, schedule)
}
