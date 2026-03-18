package services

import (
	"clinic-system/config"
)

func CalculateBill(visitID int) float64 {
	var total float64

	rows, _ := config.DB.Query("SELECT price FROM drugs WHERE visit_id=$1", visitID)
	for rows.Next() {
		var p float64
		rows.Scan(&p)
		total += p
	}

	rows, _ = config.DB.Query("SELECT price FROM lab_tests WHERE visit_id=$1", visitID)
	for rows.Next() {
		var p float64
		rows.Scan(&p)
		total += p
	}

	return total
}