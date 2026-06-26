package controllers

import (
	"net/http"

	// 🚨 อย่าลืมเปลี่ยน "api-go" เป็นชื่อ module ของคุณ (ดูจากไฟล์ go.mod)
	"api-go/config"
	"api-go/models"

	"github.com/gin-gonic/gin"
)

func GetPatientHandler(c *gin.Context) {
	id := c.Param("id")
	var p models.Patient

	// 🚨 แก้ไข 'condition' เป็น `condition` (Backtick) เพื่อให้ MySQL มองว่าเป็นชื่อคอลัมน์
	query := "SELECT id, first_name, last_name, age, `condition` FROM patients WHERE id = ?"

	// เรียกใช้ config.DB ที่เราเปิดไว้
	err := config.DB.QueryRow(query, id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.Age, &p.Condition)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลผู้ป่วย"})
		return
	}

	c.JSON(http.StatusOK, p)
}
