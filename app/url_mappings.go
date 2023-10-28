package app

import (
	"github.com/SoyebSarkar/content-creator-insight/controllers/ping"
	ytControllers "github.com/SoyebSarkar/content-creator-insight/controllers/youtube_controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapUrls(router *gin.Engine) {
	router.GET("/", ping.Ping)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/yt/list/videos/:channel_id", ytControllers.ListYoutubeVideos)
	router.GET("/yt/analyse/video/:video_id", ytControllers.AnalyseVideo)

}
