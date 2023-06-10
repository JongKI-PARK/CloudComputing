package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type EnrollmentRequest struct {
	StudentID int `json:"student_id"`
	SubjectID int `json:"subject_id"`
}

type EnrollmentResponse struct {
	Success   string `json:"success"`
	StudentID int    `json:"student_id"`
	SubjectID int    `json:"subject_id"`
}

func main() {
	// MySQL 데이터베이스 연결 설정
	db, err := sql.Open("mysql", "root:sys123457!@tcp(localhost:3306)/order_process")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Gin 라우터 설정
	router := gin.Default()
	router.Use(cors.Default())

	// 수강신청 API 핸들러
	router.POST("/orders", func(c *gin.Context) {
		var request EnrollmentRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// enrollment_order 테이블에서 수강정원(cap)과 현재신청인원(current) 확인
		var cap, current int
		err = tx.QueryRow("SELECT enrollment_cap, enrollment_current FROM enrollment_order WHERE subject_id = ?", request.SubjectID).Scan(&cap, &current)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 수강신청 가능 여부 확인
		if current >= cap {
			tx.Rollback()
			c.JSON(http.StatusOK, gin.H{"Success": "false",
				"StudentId": request.StudentID,
				"SubjectId": request.SubjectID})
			return
		}

		// enrollment_order 테이블의 현재신청인원 업데이트
		_, err = tx.Exec("UPDATE enrollment_order SET enrollment_current = ? WHERE subject_id = ?", current+1, request.SubjectID)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Printf("point1")
			return
		}

		// enrollment_complete 테이블에 수강신청 정보 저장
		_, err = tx.Exec("INSERT INTO enrollment_complete (student_id, subject_id) VALUES (?, ?)", request.StudentID, request.SubjectID)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Printf("point2")
			return
		}

		err = tx.Commit()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Printf("point3")
			return
		}

		c.JSON(http.StatusOK, gin.H{"Success": "true",
			"StudentId": request.StudentID,
			"SubjectId": request.SubjectID})
	})

	// 서버 시작
	err = router.Run(":8085")
	if err != nil {
		log.Fatal(err)
	}
}
