package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Questionnaire routes
	r.GET("/questionnaires", controllers.GetQuestionnaires)
	r.GET("/questionnaires/:id", controllers.GetQuestionnaireByID)
	r.POST("/questionnaires", controllers.CreateQuestionnaire)

	return r
}
