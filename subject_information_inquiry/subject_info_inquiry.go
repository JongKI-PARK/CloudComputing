package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Subject struct {
	ID             int    `json:"subject_id"`
	Name           string `json:"subject_name"`
	Professor      string `json:"professor"`
	Credits        int    `json:"credits"`
	Department     string `json:"department"`
	EnrollmentLimt int    `json:"enrollment_limit"`
}

func main() {
	// 데이터베이스 연결 설정
	// db, err := sql.Open("mysql", "root:root#@tcp(host.docker.internal:3306)/subject_information")
	db, err := sql.Open("mysql", "root:sys123457!@tcp(localhost:3306)/subject_information")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Gin 라우터 설정
	router := gin.Default()
	router.Use(cors.Default())

	// 과목 목록 조회 핸들러
	router.GET("/subjects", func(c *gin.Context) {
		var subjects []Subject

		// 클라이언트로부터 전달된 과목 이름 검색어
		subjectName := c.Query("name")

		// 데이터베이스에서 과목 목록 조회 (검색어 필터링 포함)
		var rows *sql.Rows
		var err error
		if subjectName != "" {
			rows, err = db.Query("SELECT subject_id, subject_name, professor, credits, department, enrollment_limit FROM subjects WHERE subject_name LIKE ?", "%"+subjectName+"%")
		} else {
			rows, err = db.Query("SELECT subject_id, subject_name, professor, credits, department, enrollment_limit FROM subjects")
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		// 조회한 결과를 Subject 객체로 매핑
		for rows.Next() {
			var subject Subject
			if err := rows.Scan(&subject.ID, &subject.Name, &subject.Professor, &subject.Credits, &subject.Department, &subject.EnrollmentLimt); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			subjects = append(subjects, subject)
		}

		c.JSON(http.StatusOK, subjects)
	})

	// 과목 상세 조회 핸들러
	router.GET("/subjects/:id", func(c *gin.Context) {
		id := c.Param("id")

		var subject Subject

		// 데이터베이스에서 과목 상세 조회
		err := db.QueryRow("SELECT subject_id, subject_name, professor, credits, department, enrollment_limit FROM subjects WHERE subject_id = ?", id).
			Scan(&subject.ID, &subject.Name, &subject.Professor, &subject.Credits, &subject.Department, &subject.EnrollmentLimt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, subject)
	})

	// 서버 시작
	if err := router.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
