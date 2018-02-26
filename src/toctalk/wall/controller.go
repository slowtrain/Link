package wall

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpController(router *gin.Engine) {

	router.Handle("GET", "/wall", func(context *gin.Context) {
		context.HTML(http.StatusOK, "wall.html", gin.H{"title": "login"})
		//context.String(http.StatusOK, "환영합니다.")
	})

	router.Handle("GET", "/wallsidebar", func(context *gin.Context) {
		context.HTML(http.StatusOK, "wallsidebar.html", gin.H{"title": "login"})
		//context.String(http.StatusOK, "환영합니다.")
	})

}
