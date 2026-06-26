package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB เป็นตัวแปร Global ให้แพ็กเกจอื่นเรียกใช้ได้ (ขึ้นต้นด้วยตัวพิมพ์ใหญ่)
var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ไม่สามารถโหลดไฟล์ .env ได้: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	DB, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("เชื่อมต่อฐานข้อมูลไม่ได้: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("ฐานข้อมูลไม่ตอบสนอง: %v", err)
	}

	log.Println("เชื่อมต่อฐานข้อมูล MySQL สำเร็จ!")
}
