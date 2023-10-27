package ytControllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/SoyebSarkar/content-creator-insight/datasource/config"
	"github.com/gin-gonic/gin"
)

type payloadDetails struct {
	ApiKey    string `json:"apiKey"`
	ChannelId string `json:"channelID"`
}

// @Summary      list youtube videos
// @Description  Returns list of  Youtube videos with channel id
// @Tags         Yt
// @Produce      json
// @Param channelID path  string true "channelId"
// @Success      200  {string}  string "pong"
// @Router       /yt/list/videos/{channelID} [get]
func ListYoutubeVideos(c *gin.Context) {

	channelId := c.Param("channelID")
	data := &payloadDetails{ApiKey: config.GCYTAPI3Key, ChannelId: channelId}
	pythonAPIUrl := "http://127.0.0.1:5000/list/YT"
	payload, _ := json.Marshal(data)
	resp, _ := http.Post(pythonAPIUrl, "application/json", bytes.NewReader(payload))
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var respObjs []map[string]interface{}
	decoder.Decode(&respObjs)
	c.JSON(http.StatusOK, respObjs)
}
