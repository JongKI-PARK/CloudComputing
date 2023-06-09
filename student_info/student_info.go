package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID         int    `json:"student_id"`
	Name       string `json:"student_name"`
	Password   string `json:"password"`
	Department string `json:"department"`
}

func main() {
	// 데이터베이스 연결 설정
	db, err := sql.Open("mysql", "root:sys123457!@tcp(localhost:3306)/student_information")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Gin 라우터 설정
	router := gin.Default()
	router.Use(cors.Default())

	// 학생 정보 조회 핸들러
	router.GET("/students", func(c *gin.Context) {
		var students []Student

		// 데이터베이스에서 학생 정보 조회
		rows, err := db.Query("SELECT student_id, student_name, pwd, department FROM students")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		// 조회한 결과를 Student 객체로 매핑
		for rows.Next() {
			var student Student
			if err := rows.Scan(&student.ID, &student.Name, &student.Password, &student.Department); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			students = append(students, student)
		}

		c.JSON(http.StatusOK, students)
	})

	// 학생 상세 조회 핸들러
	router.GET("/students/:id", func(c *gin.Context) {
		id := c.Param("id")

		var student Student

		// 데이터베이스에서 학생 상세 조회
		err := db.QueryRow("SELECT student_id, student_name, pwd, department FROM students WHERE student_id = ?", id).
			Scan(&student.ID, &student.Name, &student.Password, &student.Department)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, student)
	})

	// 로그인 핸들러
	router.POST("/login", func(c *gin.Context) {
		var loginRequest struct {
			Userid   int    `json:"Userid"`
			Password string `json:"Password"`
			Success  string `json:"Success"`
			Username string `json:"Username"`
		}

		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 데이터베이스에서 사용자명과 비밀번호 확인
		var count int
		var username string
		err := db.QueryRow("SELECT COUNT(*), student_name FROM students WHERE student_id = ? AND pwd = ? GROUP BY student_id, pwd",
			loginRequest.Userid, loginRequest.Password).Scan(&count, &username)

		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusOK, gin.H{"Userid": loginRequest.Userid,
					"Password": nil,
					"Success":  "false",
					"Username": nil})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		if count > 0 {
			// 로그인 성공
			c.JSON(http.StatusOK, gin.H{"Userid": loginRequest.Userid,
				"Password": nil,
				"Success":  "true",
				"Username": username})
		} else {
			// 로그인 실패
			c.JSON(http.StatusOK, gin.H{"Userid": loginRequest.Userid,
				"Password": nil,
				"Success":  "false",
				"Username": nil})
		}

		fmt.Println("받은 데이터 : ", loginRequest)
	})

	// 서버 시작
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
