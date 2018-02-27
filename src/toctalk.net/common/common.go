package common

import (
	"github.com/gin-gonic/gin"
)

const Layout = "slider"

func SetUpStatic(router *gin.Engine) {
	router.Static("/css", "./assets/css")
	router.Static("/fonts", "./assets/fonts")
	router.Static("/images", "./assets/images")
	router.Static("/js", "./assets/js")

	//router.Static("/lib", "./assets/lib")
	router.StaticFile("/lib/jquery-3.2.0.min.js", "./assets/lib/jquery-3.2.0.min.js")
	router.Static("/lib/bootstrap/css", "./assets/lib/bootstrap/css")
	router.Static("/lib/bootstrap/fonts", "./assets/lib/bootstrap/fonts")
	router.Static("/lib/bootstrap/js", "./assets/lib/bootstrap/js")

	router.Static("/lib/bootstrap-switch/css", "./assets/lib/bootstrap-switch/css")
	router.Static("/lib/bootstrap-switch/js", "./assets/lib/bootstrap-switch/js")

	router.Static("/lib/font-awesome/css", "./assets/lib/font-awesome/css")
	router.Static("/lib/font-awesome/fonts", "./assets/font-awesome/fonts")
	router.Static("/lib/font-awesome/less", "./assets/font-awesome/less")
	router.Static("/lib/font-awesome/scss", "./assets/font-awesome/scss")

	router.LoadHTMLGlob("./template/*/*")

}
