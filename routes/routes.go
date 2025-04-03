package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Questionnaire routes
	r.GET("/items", controllers.GetQuestionnaires)
	r.GET("/items/:id", controllers.GetQuestionnaireByID)
	r.POST("/items", controllers.CreateQuestionnaire)

	return r
}
