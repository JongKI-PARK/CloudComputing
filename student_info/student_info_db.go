// student_info_db.go :
// Set up database for "student_information" Micro Service
// One table in this db - students
// 		4 columns in "students" table

package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 데이터베이스 연결 설정
	db, err := sql.Open("mysql", "root:sys123457!@tcp(localhost:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 데이터베이스 생성
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS student_information")
	if err != nil {
		log.Fatal(err)
	}

	// 학생정보 데이터베이스 선택
	_, err = db.Exec("USE student_information")
	if err != nil {
		log.Fatal(err)
	}

	// 학생정보 테이블 생성
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			student_id INT AUTO_INCREMENT,
			student_name VARCHAR(255) NOT NULL,
			pwd VARCHAR(255) NOT NULL,
			department VARCHAR(255) NOT NULL,
			PRIMARY KEY (student_id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database and table created successfully!")
}
