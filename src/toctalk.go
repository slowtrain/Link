package main

import (
	"toctalk.net/common"
	"toctalk.net/home"
	"toctalk.net/login"
	"toctalk.net/personalprofile"
	"toctalk.net/settings"
	"toctalk.net/userprofile"
	"toctalk.net/wall"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	common.SetUpStatic(router)
	home.SetUpController(router)
	login.SetUpController(router)
	personalprofile.SetUpController(router)
	settings.SetUpController(router)
	userprofile.SetUpController(router)
	wall.SetUpController(router)

	router.Run(":8080")
}
