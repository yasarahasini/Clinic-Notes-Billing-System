package models

import (
	"clinic-system/config"
	"fmt"
)

type LabTest struct {
	ID       int
	VisitID  int
	TestName string
	Price    float64
}


func AddLabTest(visitID int, testName string, price float64) error {
	_, err := config.DB.Exec(
		"INSERT INTO lab_tests(visit_id, test_name, price) VALUES($1,$2,$3)",
		visitID, testName, price,
	)
	if err != nil {
		return fmt.Errorf("cannot add lab test: %v", err)
	}
	return nil
}