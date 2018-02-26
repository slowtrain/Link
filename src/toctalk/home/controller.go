package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpController(router *gin.Engine) {

	router.Handle("GET", "/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"title": "Main website"})
		//context.String(http.StatusOK, "환영합니다.")
	})
}
