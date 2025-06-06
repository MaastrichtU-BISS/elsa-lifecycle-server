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

	// images
	r.Static("/uploads", "./uploads") // Serve static files from the "uploads" directory

	// Lifecycle routes
	r.GET("/lifecycles", controllers.GetLifecycles)
	r.GET("/lifecycles/:id", controllers.GetLifecyclesByID)
	r.GET("/lifecycles/:id/phases", controllers.GetPhases)

	// Phase routes
	r.GET("/phases/:id", controllers.GetPhaseById)

	// Reflection routes
	r.GET("/reflections/:id", controllers.GetReflectionByID)

	// Journal routes
	r.GET("/journals/:id", controllers.GetJournalByID)

	// ReflectionAnswer routes
	r.GET("/reflectionAnswers/:id", controllers.GetReflectionAnswerByID)
	r.POST("/reflectionAnswers", controllers.CreateReflectionAnswer)
	r.PUT("/reflectionAnswers/:id/edit", controllers.EditReflectionAnswer)

	// JournalAnswer routes
	r.GET("/journalAnswers/:id", controllers.GetJournalAnswerByID)
	r.POST("/journalAnswers", controllers.CreateJournalAnswer)
	r.PUT("/journalAnswers/:id/edit", controllers.EditJournalAnswer)

	// Tool routes
	r.GET("/tools", controllers.GetTools)
	r.GET("/tools/:id", controllers.GetToolByID)
	r.POST("/tools", controllers.CreateTool)
	r.PUT("/tools/:id/edit", controllers.EditTool)

	// Recommendation routes
	r.GET("/recommendations/:reflectionId", controllers.GetRecommendations)

	// RecommendationAnswer routes
	r.GET("/recommendationAnswers/:id", controllers.GetRecommendationAnswerByID)
	r.POST("/recommendationAnswers", controllers.CreateRecommendationAnswer)
	r.PUT("/recommendationAnswers/:id/edit", controllers.EditRecommendationAnswer)

	return r
}
