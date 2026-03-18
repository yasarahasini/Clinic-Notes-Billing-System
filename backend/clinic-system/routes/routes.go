package routes

import (
	"clinic-system/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/visit", controllers.CreateVisit)
}