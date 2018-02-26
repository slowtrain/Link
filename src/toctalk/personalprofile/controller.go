package personalprofile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpController(router *gin.Engine) {

	router.Handle("GET", "/personalprofile", func(context *gin.Context) {
		context.HTML(http.StatusOK, "personalprofile.html", gin.H{"title": "login"})
		//context.String(http.StatusOK, "환영합니다.")
	})

}
