package main

import (
	"database/sql"
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
			c.JSON(http.StatusBadRequest, gin.H{"success": "false", "error": err.Error()})
			return
		}

		// 이미 수강신청이 존재하는지 확인
		var existingPlan EnrollmentPlan
		err = db.QueryRow("SELECT student_id, subject_id FROM enrollment_plan WHERE student_id = ? AND subject_id = ?", request.StudentID, request.SubjectID).Scan(&existingPlan.StudentID, &existingPlan.SubjectID)
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": "false", "error": "Enrollment plan already exists", "studentId": request.StudentID, "subjectId": request.SubjectID})
			return
		} else if err != sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"success": "false", "error": err.Error()})
			return
		}

		// enrollment_plan 테이블에 수강신청 정보 저장
		result, err := db.Exec("INSERT INTO enrollment_plan (student_id, subject_id) VALUES (?, ?)", request.StudentID, request.SubjectID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": "false", "error": err.Error(), "studentId": request.StudentID, "subjectId": request.SubjectID})
			return
		}

		// 새로 추가된 수강신청의 ID 가져오기
		id, err := result.LastInsertId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": "false", "error": err.Error(), "studentId": request.StudentID, "subjectId": request.SubjectID})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "true", "error": "", "id": id, "studentId": request.StudentID, "subjectId": request.SubjectID})
	})

	// 수강신청 계획 내역 조회 API 핸들러
	// 수강신청 계획 내역 조회 API 핸들러
	router.GET("/planner/:id", func(c *gin.Context) {
		studentIDStr := c.Param("id")
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
		c.JSON(http.StatusOK, gin.H{"enrollmentPlans": enrollmentPlans})
	})

	// 서버 시작
	err = router.Run(":8084")
	if err != nil {
		log.Fatal(err)
	}
}
