package main

import (
	"toctalk/common"
	"toctalk/home"
	"toctalk/login"
	"toctalk/personalprofile"
	"toctalk/settings"
	"toctalk/userprofile"
	"toctalk/wall"

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
