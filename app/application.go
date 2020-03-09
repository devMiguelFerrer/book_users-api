package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication ...
func StartApplication() {
	urlMapping()
	router.Run(":5000")
}
