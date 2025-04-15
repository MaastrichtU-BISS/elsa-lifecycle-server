package routes

import (
	"server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS middleware configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Take from env variable
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Allow HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allow headers
		AllowCredentials: true,                                                // Allow cookies if needed
	}))

	// Questionnaire routes
	r.GET("/questionnaires", controllers.GetQuestionnaires)
	r.GET("/questionnaires/:id", controllers.GetQuestionnaireByID)
	r.POST("/questionnaires", controllers.CreateQuestionnaire)

	// Answer routes
	r.GET("/questionnaires/:id/answers", controllers.GetAnswers)
	r.GET("/answers/:id", controllers.GetAnswerByID)
	r.POST("/questionnaires/:id/answers", controllers.CreateAnswer)
	r.PUT("/answers/:id/edit", controllers.EditAnswer)

	return r
}
