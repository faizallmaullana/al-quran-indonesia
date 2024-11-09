package middleware

import (
	"time"

	"github.com/faizallmaullana/al-quran-indonesia/router"
	"github.com/gin-gonic/gin"
)

func ActivateMiddleware(r *gin.Engine) {
	router.SetupRouter(r)
	r.Use(CORSMiddleware())
	r.Use(RateLimiterMiddleware(50, 1*time.Second, 1*time.Hour))
	r.Use(SecurityHeader())
}
