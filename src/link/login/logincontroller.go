package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpController(router *gin.Engine) {
	router.Handle("GET", "/login/:name", func(context *gin.Context) {
		name := context.Param("name")
		family := context.Query("family")
		context.String(http.StatusOK, "로그인하셨습니다  %s %s", family, name)
	})
}
