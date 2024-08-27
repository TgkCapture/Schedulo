package service

import (
    "github.com/TgkCapture/Schedulo/db"
    "github.com/TgkCapture/Schedulo/models"
    "fmt"
)

// AddSchedule adds a new schedule to the database.
func AddSchedule(schedule models.Schedule) error {
    query := `INSERT INTO schedules (title, time_slot, channel) VALUES (?, ?, ?)`

    _, err := db.DB.Exec(query, schedule.Title, schedule.TimeSlot, schedule.Channel)
    if err != nil {
        return fmt.Errorf("error adding schedule: %w", err)
    }
    return nil
}

// GetSchedules retrieves all schedules from the database.
func GetSchedules() ([]models.Schedule, error) {
    query := `SELECT id, title, time_slot, channel FROM schedules`

    rows, err := db.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error retrieving schedules: %w", err)
    }
    defer rows.Close()

    var schedules []models.Schedule
    for rows.Next() {
        var schedule models.Schedule
        if err := rows.Scan(&schedule.ID, &schedule.Title, &schedule.TimeSlot, &schedule.Channel); err != nil {
            return nil, fmt.Errorf("error scanning schedule: %w", err)
        }
        schedules = append(schedules, schedule)
    }

    return schedules, nil
}

// UpdateSchedule updates an existing schedule in the database.
func UpdateSchedule(schedule models.Schedule) error {
    query := `UPDATE schedules SET title = ?, time_slot = ?, channel = ? WHERE id = ?`

    _, err := db.DB.Exec(query, schedule.Title, schedule.TimeSlot, schedule.Channel, schedule.ID)
    if err != nil {
        return fmt.Errorf("error updating schedule: %w", err)
    }
    return nil
}

// DeleteSchedule deletes a schedule from the database.
func DeleteSchedule(id int64) error {
    query := `DELETE FROM schedules WHERE id = ?`

    _, err := db.DB.Exec(query, id)
    if err != nil {
        return fmt.Errorf("error deleting schedule: %w", err)
    }
    return nil
}

// GetScheduleByID retrieves a schedule by its ID.
func GetScheduleByID(id int64) (*models.Schedule, error) {
    query := `SELECT id, title, time_slot, channel FROM schedules WHERE id = $1`
    var schedule models.Schedule
    err := db.DB.QueryRow(query, id).Scan(&schedule.ID, &schedule.Title, &schedule.TimeSlot, &schedule.Channel)
    if err != nil {
        return nil, fmt.Errorf("error retrieving schedule: %w", err)
    }
    return &schedule, nil
}