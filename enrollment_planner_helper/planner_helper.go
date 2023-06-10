package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type EnrollmentPlan struct {
	ID        int `json:"id"`
	StudentID int `json:"studentId"`
	SubjectID int `json:"subjectId"`
}

type Request struct {
	StudentID int `json:"student_id"`
	SubjectID int `json:"subject_id"`
}

func main() {
	// MySQL 데이터베이스 연결 설정
	db, err := sql.Open("mysql", "root:sys123457!@tcp(localhost:3306)/enrollment_planner")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Gin 라우터 설정
	router := gin.Default()
	router.Use(cors.Default())

	// 수강신청 저장 API 핸들러
	router.POST("/planner", func(c *gin.Context) {
		var request Request
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// enrollment_plan 테이블에 수강신청 정보 저장
		_, err = db.Exec("INSERT INTO enrollment_plan (student_id, subject_id) VALUES (?, ?)", request.StudentID, request.SubjectID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true})
	})

	// 수강신청 계획 내역 조회 API 핸들러
	router.GET("/planner", func(c *gin.Context) {
		studentIDStr := c.Query("studentId")
		studentID, err := strconv.Atoi(studentIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid studentId"})
			return
		}

		// enrollment_plan 테이블에서 studentID에 해당하는 과목 ID 조회
		rows, err := db.Query("SELECT id, student_id, subject_id FROM enrollment_plan WHERE student_id = ?", studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		// 결과를 담을 슬라이스 생성
		enrollmentPlans := make([]EnrollmentPlan, 0)

		// 조회 결과를 슬라이스에 추가
		for rows.Next() {
			var plan EnrollmentPlan
			err := rows.Scan(&plan.ID, &plan.StudentID, &plan.SubjectID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			enrollmentPlans = append(enrollmentPlans, plan)
		}

		// 조회 결과를 JSON 형식으로 반환
		response, err := json.Marshal(enrollmentPlans)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, string(response))
	})

	// 서버 시작
	err = router.Run(":8084")
	if err != nil {
		log.Fatal(err)
	}
}
