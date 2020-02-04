package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication entry point application
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
