package controllers

import (
	"clinic-system/config"
	"clinic-system/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateVisit(c *gin.Context) {
	var body struct {
		Text string `json:"text"`
	}

	c.BindJSON(&body)

	parsed := services.ParseText(body.Text)

	
	var visitID int
	config.DB.QueryRow("INSERT INTO visits DEFAULT VALUES RETURNING id").Scan(&visitID)


	for _, d := range parsed.Drugs {
		config.DB.Exec("INSERT INTO drugs (visit_id, name, dosage, price) VALUES ($1,$2,$3,$4)",
			visitID, d["name"], d["dosage"], 10)
	}

	
	for _, l := range parsed.Labs {
		config.DB.Exec("INSERT INTO lab_tests (visit_id, test_name, price) VALUES ($1,$2,$3)",
			visitID, l, 20)
	}


	config.DB.Exec("INSERT INTO notes (visit_id, content) VALUES ($1,$2)",
		visitID, parsed.Notes)

	c.JSON(http.StatusOK, gin.H{
		"visit_id": visitID,
		"parsed":   parsed,
	})
}