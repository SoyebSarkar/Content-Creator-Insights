package app

import (
	"github.com/SoyebSarkar/content-creator-insight/controllers/ping"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapUrls(router *gin.Engine) {
	router.GET("/", ping.Ping)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
