package models

type Schedule struct {
    ID       int64  `json:"id"`
    Title    string `json:"title"`
    TimeSlot string `json:"time_slot"`
    Channel  string `json:"channel"`
}
