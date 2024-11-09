package router

import (
	"github.com/faizallmaullana/al-quran-indonesia/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	base_url := "/api/v1/quran/resources"
	r := router.Group(base_url)

	r.GET("/surah", controller.GetAllSurah)
	r.GET("/surah/:id", controller.GetSurah)
}
