package models

import (
	"clinic-system/config"
	"fmt"
	"time"
)

type Visit struct {
	ID        int
	PatientID int
	CreatedAt time.Time
}


func CreateVisit(patientID int) (int, error) {
	var id int
	err := config.DB.QueryRow(
		"INSERT INTO visits(patient_id) VALUES($1) RETURNING id",
		patientID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("cannot create visit: %v", err)
	}
	return id, nil
}

func GetVisitsByPatient(patientID int) ([]Visit, error) {
	rows, err := config.DB.Query(
		"SELECT id, patient_id, created_at FROM visits WHERE patient_id=$1",
		patientID,
	)
	if err != nil {
		return nil, fmt.Errorf("cannot fetch visits: %v", err)
	}
	defer rows.Close()

	var visits []Visit
	for rows.Next() {
		var v Visit
		if err := rows.Scan(&v.ID, &v.PatientID, &v.CreatedAt); err != nil {
			return nil, err
		}
		visits = append(visits, v)
	}
	return visits, nil
}