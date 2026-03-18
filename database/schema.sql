package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	// 1. Connect to default postgres DB
	defaultConn := "user=postgres password=12345 dbname=clinic sslmode=disable"
	db, err := sql.Open("postgres", defaultConn)
	if err != nil {
		log.Fatal("Error connecting to default DB:", err)
	}

	// 2. Create 'clinic' database if it doesn't exist
	_, err = db.Exec(`CREATE DATABASE clinic`)
	if err != nil {
		fmt.Println("Database 'clinic' might already exist:", err)
	} else {
		fmt.Println("Database 'clinic' created successfully!")
	}
	db.Close()

	// 3. Connect to clinic database
	connStr := "user=postgres password=12345 dbname=clinic sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to clinic DB:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot reach clinic DB:", err)
	}

	DB = db
	fmt.Println("Connected to PostgreSQL clinic database!")

	// 4. Create tables
	createTables()
}