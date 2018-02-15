package common

import (
	"github.com/gin-gonic/gin"
)

func SetUpStatic(router *gin.Engine) {
	router.Static("/css", "./static/css")
	router.Static("/fonts", "./static/fonts")
	router.Static("/images", "./static/images")
	router.Static("/js", "./static/js")
	router.Static("/saas", "./static/sass")
}
