package routes

import (
	"os"
	"regexp"
	"strings"

	"server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// buildCORSConfigFromString builds a cors.Config from a comma-separated origins string.
// It supports simple glob-style wildcards where '*' -> '.*' and '?' -> '.'.
// Examples:
//   "http://example.com,https://*.example.org"
//   "http://localhost:*"
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
			esc = strings.ReplaceAll(esc, "\\?", ".")
			re, err := regexp.Compile("^" + esc + "$")
			if err == nil {
				regexps = append(regexps, re)
			}
			// if compile fails, skip the pattern
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

	// Questionnaire routes
	r.GET("/questionnaires", controllers.GetQuestionnaires)
	r.GET("/questionnaires/:id", controllers.GetQuestionnaireByID)
	r.DELETE("/questionnaires/:id", controllers.DeleteQuestionnaire)
	r.POST("/questionnaires", controllers.CreateQuestionnaire)

	// Answer routes
	r.GET("/questionnaires/:id/answers", controllers.GetAnswers)
	r.GET("/answers/:id", controllers.GetAnswerByID)
	r.POST("/questionnaires/:id/answers", controllers.CreateAnswer)
	r.PUT("/answers/:id/edit", controllers.EditAnswer)

	return r
}
