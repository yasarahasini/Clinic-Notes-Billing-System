package models

import (
	"clinic-system/config"
	"fmt"
)

type Patient struct {
	ID   int
	Name string
	Age  int
}


func CreatePatient(name string, age int) (int, error) {
	var id int
	err := config.DB.QueryRow(
		"INSERT INTO patients(name, age) VALUES($1,$2) RETURNING id",
		name, age,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("cannot create patient: %v", err)
	}
	return id, nil
}

func GetPatient(id int) (*Patient, error) {
	p := &Patient{}
	err := config.DB.QueryRow(
		"SELECT id, name, age FROM patients WHERE id=$1", id,
	).Scan(&p.ID, &p.Name, &p.Age)
	if err != nil {
		return nil, fmt.Errorf("cannot get patient: %v", err)
	}
	return p, nil
}