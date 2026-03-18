package main

import (
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"
	"clinic-system/config"
	"clinic-system/models"
)

func main() {
	config.ConnectDB()

	r := gin.Default()


	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Clinic system running!"})
	})

	
	r.POST("/patients", func(c *gin.Context) {
		var body struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		id, err := models.CreatePatient(body.Name, body.Age)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Patient registered", "patient_id": id})
	})

	
	r.POST("/visit", func(c *gin.Context) {
		var body struct {
			PatientID int    `json:"patient_id"`
			Text      string `json:"text"`
		}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		visitID, err := models.CreateVisit(body.PatientID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		lines := strings.Split(body.Text, "\n")
		var notesContent string

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			lower := strings.ToLower(line)

			if strings.Contains(lower, "mg") {
				models.AddDrug(visitID, line, "1x daily", 10)
			} else if strings.Contains(lower, "test") {
				models.AddLabTest(visitID, line, 20)
			} else {
				notesContent += line + " "
			}
		}

		if notesContent != "" {
			models.AddNote(visitID, notesContent)
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Visit saved successfully",
			"visit_id": visitID,
		})
	})

	
	r.GET("/patients", func(c *gin.Context) {
		rows, err := config.DB.Query("SELECT id, name, age FROM patients")
		if err != nil {
			c.JSON(500, gin.H{"error": "Cannot fetch patients"})
			return
		}
		defer rows.Close()

		var patients []map[string]interface{}
		for rows.Next() {
			var id, age int
			var name string
			rows.Scan(&id, &name, &age)
			patients = append(patients, map[string]interface{}{
				"id":   id,
				"name": name,
				"age":  age,
			})
		}

		c.JSON(200, patients)
	})

	r.Run(":8080")
}