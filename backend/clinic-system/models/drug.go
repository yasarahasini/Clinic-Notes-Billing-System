package models

import (
	"clinic-system/config"
	"fmt"
)

type Drug struct {
	ID      int
	VisitID int
	Name    string
	Dosage  string
	Price   float64
}


func AddDrug(visitID int, name string, dosage string, price float64) error {
	_, err := config.DB.Exec(
		"INSERT INTO drugs(visit_id, name, dosage, price) VALUES($1,$2,$3,$4)",
		visitID, name, dosage, price,
	)
	if err != nil {
		return fmt.Errorf("cannot add drug: %v", err)
	}
	return nil
}