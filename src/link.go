package main

import (
	"link/common"
	"link/home"
	"link/login"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	common.SetUpStatic(router)
	home.SetUpController(router)
	login.SetUpController(router)

	// router.Handle("GET", "/:name", func(context *gin.Context) {
	// 	name := context.Param("name")
	// 	family := context.Query("family")
	// 	context.String(http.StatusOK, "환영합니다  %s %s", family, name)
	// })
	router.Run(":8080")
}
