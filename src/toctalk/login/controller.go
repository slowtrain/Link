package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpController(router *gin.Engine) {

	router.Handle("GET", "/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{"title": "login"})
		//context.String(http.StatusOK, "환영합니다.")
	})

	router.Handle("GET", "/logincallback", func(context *gin.Context) {
		context.HTML(http.StatusOK, "logincallback.html", gin.H{"title": "login"})
		//context.String(http.StatusOK, "환영합니다.")
	})

	router.Handle("GET", "/signin", func(context *gin.Context) {
		context.HTML(http.StatusOK, "signin.html", gin.H{"title": "login"})
		//context.String(http.StatusOK, "환영합니다.")
	})

}
