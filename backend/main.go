package main

import (
	"log"

	// 🚨 อย่าลืมเปลี่ยน "api-go" เป็นชื่อ module ของคุณ (ดูจากไฟล์ go.mod)
	"api-go/config"
	"api-go/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. เชื่อมต่อฐานข้อมูล
	config.ConnectDB()
	defer config.DB.Close()

	// 2. ตั้งค่า Router และ Middleware
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
	}))

	// 3. กำหนด Routes
	r.GET("/patients/:id", controllers.GetPatientHandler)

	// 4. เริ่มเซิร์ฟเวอร์
	log.Println("พร้อมทำงานแล้วที่ http://localhost:9999")
	r.Run(":9999")
}
