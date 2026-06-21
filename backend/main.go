package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 🚨 จุดที่ 1: เปลี่ยน import จาก lib/pq เป็น mysql
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ไม่สามารถโหลดไฟล์ .env ได้: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// 🚨 จุดที่ 2: เปลี่ยนรูปแบบ Connection String เป็นของ MySQL
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// 🚨 จุดที่ 3: สั่งเชื่อมต่อโดยใช้คำว่า "mysql" แทน "postgres"
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("เชื่อมต่อฐานข้อมูลไม่ได้: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("ฐานข้อมูลไม่ตอบสนอง: %v", err)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
	}))
	r.GET("/patients/:id", getPatientHandler)

	log.Println("พร้อมทำงานแล้วที่ http://localhost:9999")
	r.Run(":9999")
}

func getPatientHandler(c *gin.Context) {
	id := c.Param("id")
	var p Patient

	// 🚨 จุดที่ 4: เปลี่ยนจาก $1 เป็น ? (เพราะ MySQL ใช้เครื่องหมาย ? ในการรับตัวแปร)
	query := "SELECT id, first_name, last_name, age, 'condition' FROM patients WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.Age, &p.Condition)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลผู้ป่วย"})
		return
	}

	c.JSON(http.StatusOK, p)
}
