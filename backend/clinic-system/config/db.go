package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "user=postgres password=12345 dbname=clinic sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB not reachable:", err)
	}

	DB = db
	fmt.Println("Connected to DB!")

	createTables()
}

func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS patients (
			id SERIAL PRIMARY KEY,
			name TEXT,
			age INT
		);`,

		`CREATE TABLE IF NOT EXISTS visits (
			id SERIAL PRIMARY KEY,
			patient_id INT REFERENCES patients(id),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS drugs (
			id SERIAL PRIMARY KEY,
			visit_id INT REFERENCES visits(id),
			name TEXT,
			dosage TEXT,
			price NUMERIC
		);`,

		`CREATE TABLE IF NOT EXISTS lab_tests (
			id SERIAL PRIMARY KEY,
			visit_id INT REFERENCES visits(id),
			test_name TEXT,
			price NUMERIC
		);`,

		`CREATE TABLE IF NOT EXISTS notes (
			id SERIAL PRIMARY KEY,
			visit_id INT REFERENCES visits(id),
			content TEXT
		);`,
	}

	for _, q := range queries {
		_, err := DB.Exec(q)
		if err != nil {
			log.Fatal("Table error:", err)
		}
	}

	fmt.Println("Tables ready!")
}