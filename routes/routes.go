package routes

import (
	"log"
	"os"
	"regexp"
	"strings"

	"server/controllers"
	"server/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// buildCORSConfigFromString builds a cors.Config from a comma-separated origins string.
// It supports simple glob-style wildcards where '*' -> '.*' and '?' -> '.'.
// Examples:
//
//	"http://example.com,https://*.example.org"
//	"http://localhost:*"
func buildCORSConfigFromString(env string) cors.Config {
	// defaults
	if strings.TrimSpace(env) == "" {
		env = "http://localhost,http://localhost:*"
	}

	parts := []string{}
	for _, p := range strings.Split(env, ",") {
		if s := strings.TrimSpace(p); s != "" {
			parts = append(parts, s)
		}
	}

	cfg := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	var exacts []string
	var regexps []*regexp.Regexp

	for _, p := range parts {
		if strings.ContainsAny(p, "*?") {
			// convert glob to regexp: escape then replace escaped '*' and '?' back to regex
			esc := regexp.QuoteMeta(p)
			esc = strings.ReplaceAll(esc, "\\*", ".*")
			re, err := regexp.Compile("^" + esc + "$")
			if err == nil {
				regexps = append(regexps, re)
			} else {
				log.Printf("Invalid CORS origin pattern '%s': %v", p, err)
			}
		} else {
			exacts = append(exacts, p)
		}
	}

	if len(regexps) > 0 {
		// Use AllowOriginFunc to evaluate both exacts and regexps
		cfg.AllowOriginFunc = func(origin string) bool {
			for _, e := range exacts {
				if origin == e {
					return true
				}
			}
			for _, r := range regexps {
				if r.MatchString(origin) {
					return true
				}
			}
			return false
		}
	} else {
		cfg.AllowOrigins = exacts
	}

	return cfg
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	cfg := buildCORSConfigFromString(os.Getenv("CORS_ALLOW_ORIGINS"))
	r.Use(cors.New(cfg))

	// images
	r.Static("/uploads", "./uploads") // Serve static files from the "uploads" directory

	// authentication routes
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)

	// Lifecycle routes
	r.GET("/lifecycles", controllers.GetLifecycles)
	r.GET("/lifecycles/:id", controllers.GetLifecyclesByID)
	// r.GET("/lifecycles/:id/phases", controllers.GetPhases)

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
		protected.GET("/reflectionAnswers", controllers.GetReflectionAnswerByUserIdAndReflectionID)
		protected.POST("/reflectionAnswers", controllers.CreateReflectionAnswer)
		protected.PUT("/reflectionAnswers/:id/edit", controllers.EditReflectionAnswer)

		// JournalAnswer routes
		protected.GET("/journalAnswers/:id", controllers.GetJournalAnswerByID)
		protected.GET("/journalAnswers", controllers.GetJournalAnswerByUserIdAndJournalID)
		protected.POST("/journalAnswers", controllers.CreateJournalAnswer)
		protected.PUT("/journalAnswers/:id/edit", controllers.EditJournalAnswer)

		// RecommendationAnswer routes
		protected.GET("/recommendationAnswers/:id", controllers.GetRecommendationAnswerByID)
		protected.GET("/recommendationAnswers", controllers.GetRecommendationAnswerByUserIdAndRecommendationID)
		protected.POST("/recommendationAnswers", controllers.CreateRecommendationAnswer)
		protected.PUT("/recommendationAnswers/:id/edit", controllers.EditRecommendationAnswer)
	}

	return r
}
