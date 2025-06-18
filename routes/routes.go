package routes

import (
	"server/controllers"
	"server/utils"

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

	// authentication routes
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)

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

	// Tool routes
	r.GET("/tools", controllers.GetTools)
	r.GET("/tools/:id", controllers.GetToolByID)
	r.POST("/tools", controllers.CreateTool)
	r.PUT("/tools/:id/edit", controllers.EditTool)

	// Recommendation routes
	r.GET("/recommendations/:reflectionId", controllers.GetRecommendations)

	// Protected routes
	protected := r.Group("/")
	protected.Use(utils.AuthMiddleware())
	{
		// User routes
		protected.GET("/user", controllers.GetUser)

		// ReflectionAnswer routes
		protected.GET("/reflectionAnswers/:id", controllers.GetReflectionAnswerByID)
		protected.POST("/reflectionAnswers", controllers.CreateReflectionAnswer)
		protected.PUT("/reflectionAnswers/:id/edit", controllers.EditReflectionAnswer)

		// JournalAnswer routes
		protected.GET("/journalAnswers/:id", controllers.GetJournalAnswerByID)
		protected.POST("/journalAnswers", controllers.CreateJournalAnswer)
		protected.PUT("/journalAnswers/:id/edit", controllers.EditJournalAnswer)

		// RecommendationAnswer routes
		protected.GET("/recommendationAnswers/:id", controllers.GetRecommendationAnswerByID)
		protected.POST("/recommendationAnswers", controllers.CreateRecommendationAnswer)
		protected.PUT("/recommendationAnswers/:id/edit", controllers.EditRecommendationAnswer)
	}

	return r
}
