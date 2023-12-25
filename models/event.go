package models

import (
	"fmt"
	"github.com/eliofery/golang-restapi/database"
	"time"
)

type Event struct {
	ID          int
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e Event) Save() error {
	op := "event.Save"

	query := "INSERT INTO events(title, description, location, dateTime, user_id) VALUES(?,?,?,?,?)"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	e.ID = int(id)

	return nil
}

func GetAllEvents() ([]Event, error) {
	op := "models.GetAllEvents"

	query := "SELECT * FROM events"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event

		err = rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		events = append(events, event)
	}

	return events, nil
}
