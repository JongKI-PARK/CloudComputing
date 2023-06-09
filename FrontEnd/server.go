package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Student struct {
	ID         int    `json:"student_id"`
	Name       string `json:"student_name"`
	Password   string `json:"pwd"`
	Department string `json:"department"`
}

type Subject struct {
	ID             int    `json:"subject_id"`
	Name           string `json:"subject_name"`
	Professor      string `json:"professor"`
	Credits        int    `json:"credits"`
	Department     string `json:"department"`
	EnrollmentLimt int    `json:"enrollment_limit"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// Gin 엔진 생성
	router := gin.Default()
	router.Use(cors.Default())

	// 정적 파일 서버 설정
	router.Static("/static", "./static")

	// HTML 템플릿 로드
	router.LoadHTMLGlob("templates/*")

	// 메인 페이지 라우팅
	router.GET("/main", func(c *gin.Context) {
		c.HTML(200, "main.html", gin.H{})
	})

	router.GET("/main/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.HTML(200, "main.html", gin.H{"ID": id})
	})

	// 서버 시작
	log.Println("Server started on http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
