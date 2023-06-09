// subject_info_inquiry_db.go :
// Set up database for "subject_information_inquiry" Micro Service
// One table in this db - subjects
// 		6 columns in "subjects" table

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUsername := "root"
	dbPassword := "sys123457!"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "subject_information"

	// Connect to MySQL DB
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Creating database if not exist
	createDBQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	_, err = db.Exec(createDBQuery)
	if err != nil {
		panic(err)
	}

	// Use database "subject_information"
	_, err = db.Exec(fmt.Sprintf("USE %s", dbName))
	if err != nil {
		panic(err)
	}

	// Create table subjects
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS subjects (
			subject_id INT AUTO_INCREMENT PRIMARY KEY,
			subject_name VARCHAR(255) NOT NULL,
			professor VARCHAR(255) NOT NULL,
			credits INT NOT NULL,
			department VARCHAR(255) NOT NULL,
			enrollment_limit INT NOT NULL
		)
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database and table created successfully!")
}
