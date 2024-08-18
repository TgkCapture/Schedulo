package service

import "fmt"

func AddSchedule(title, timeSlot string) {
    fmt.Printf("Adding schedule: %s at %s\n", title, timeSlot)
}
