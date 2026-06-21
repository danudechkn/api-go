package main

// Patient กำหนดรูปแบบข้อมูลที่แปลงไปมาระหว่าง Go กับ JSON
type Patient struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Condition string `json:"condition"`
}
