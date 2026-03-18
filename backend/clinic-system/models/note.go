package models

import (
	"clinic-system/config"
	"fmt"
)

type Note struct {
	ID      int
	VisitID int
	Content string
}


func AddNote(visitID int, content string) error {
	_, err := config.DB.Exec(
		"INSERT INTO notes(visit_id, content) VALUES($1,$2)",
		visitID, content,
	)
	if err != nil {
		return fmt.Errorf("cannot add note: %v", err)
	}
	return nil
}

func GetNotesByVisit(visitID int) ([]Note, error) {
	rows, err := config.DB.Query(
		"SELECT id, visit_id, content FROM notes WHERE visit_id=$1",
		visitID,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot fetch notes: %v", err)
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.VisitID, &n.Content); err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}

	return notes, nil
}