package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	MapUrls(router)
	fmt.Println("About to start the Appplication")
	router.Run(":8080")
}
