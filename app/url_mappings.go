package app

import (
	"github.com/SoyebSarkar/content-creator-insight/controllers/ping"
	users "github.com/SoyebSarkar/content-creator-insight/controllers/user_controller"
	ytControllers "github.com/SoyebSarkar/content-creator-insight/controllers/youtube_controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapUrls(router *gin.Engine) {
	router.GET("/", ping.Ping)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/signup", users.CreateUser)
	router.POST("/login", users.LoginUser)

	router.GET("/yt/channel_code/:email", ytControllers.GetUserchannelCode)
	router.GET("/yt/list/videos/:channel_id", ytControllers.ListYoutubeVideos)
	router.GET("/yt/:channel_id/:video_id", ytControllers.GetVideoComment)
	router.GET("/yt/analyse/:channel_id/:video_id", ytControllers.AnalyseVideo)

}
