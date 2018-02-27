package userprofile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpController(router *gin.Engine) {

	router.Handle("GET", "/userprofile", func(context *gin.Context) {
		context.HTML(http.StatusOK, "userprofile.html", gin.H{"title": "login"})
		//context.String(http.StatusOK, "환영합니다.")
	})

}
