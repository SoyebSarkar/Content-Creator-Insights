package ytControllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SoyebSarkar/content-creator-insight/datasource/config"
	"github.com/SoyebSarkar/content-creator-insight/datasource/mysql"
	"github.com/SoyebSarkar/content-creator-insight/utils/errors"
	"github.com/gin-gonic/gin"
)

type payloadDetails struct {
	ApiKey    string `json:"apiKey"`
	ChannelId string `json:"channelID"`
}
type Comment struct {
	Comment   string  `json:"comment"`
	UserName  string  `json:"username"`
	Date      string  `json:"date"`
	LikeCount int     `json:"likeCount"`
	Possitive float32 `json:"possitive"`
	Negative  float32 `json:"negative"`
	Neutral   float32 `json:"neutral"`
}
type CommentDetails struct {
	PossitiveCount int       `json:"possitive_count"`
	NegetiveCount  int       `json:"negetive_count"`
	NeutralCount   int       `json:"neutral_count"`
	Comments       []Comment `json:"comments"`
}

const (
	queryListUserChannel    = "SELECT `yt_channel_code` FROM `user_yt_channel` WHERE email = ?"
	queryListCommentDetails = "SELECT `comment`, `username`, `positive`, `negetive`, `neutral`, `date` FROM `16_comment_stat` WHERE `channel_id` = ? AND `video_id` = ?"
)

// @Summary      list analyse detail of video comments
// @Description  Returns Sentiment analyse of all video comments
// @Tags         Yt
// @Produce      json
// @Param video_id path  string true "videoID"
// @Success      200  {string}  string "pong"
// @Router       /yt/{channel_id}/{video_id} [get]
func GetVideoComment(c *gin.Context) {
	videoId := c.Param("video_id")
	channelId := c.Param("channel_id")
	rows, err := mysql.Client.Query(queryListCommentDetails, channelId, videoId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, errors.NewInternalServerError("Database Error"))
		return
	}
	var result CommentDetails
	result.Comments = make([]Comment, 0)
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.Comment, &comment.UserName, &comment.Possitive, &comment.Negative, &comment.Neutral, &comment.Date); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, errors.NewInternalServerError("Database Error"))
			return
		}
		if comment.Possitive >= comment.Negative && comment.Possitive >= comment.Neutral {
			result.PossitiveCount++
		} else if comment.Negative >= comment.Possitive && comment.Negative >= comment.Neutral {
			result.NegetiveCount++
		} else {
			result.NeutralCount++
		}
		result.Comments = append(result.Comments, comment)
	}
	c.JSON(http.StatusOK, result)

}

// @Summary      list youtube videos
// @Description  Returns list of  Youtube Channel code of the user
// @Tags         Yt
// @Produce      json
// @Param email path  string true "email"
// @Success      200  []string  string "channelCode list"
// @Router       /yt/channel_code/{emai} [get]
func GetUserchannelCode(c *gin.Context) {

	email := c.Param("email")

	rows, err := mysql.Client.Query(queryListUserChannel, email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, errors.NewInternalServerError("Database error"))
		return
	}
	defer rows.Close()
	result := make([]string, 0)
	for rows.Next() {
		var channelCode string
		if err := rows.Scan(&channelCode); err != nil {
			fmt.Println(result)
			c.JSON(http.StatusOK, errors.NewInternalServerError("Database Error"))
			return
		}
		result = append(result, channelCode)
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}

// @Summary      list youtube videos
// @Description  Returns list of  Youtube videos with channel id
// @Tags         Yt
// @Produce      json
// @Param channel_id path  string true "channelId"
// @Success      200  {string}  string "pong"
// @Router       /yt/list/videos/{channelID} [get]
func ListYoutubeVideos(c *gin.Context) {

	channelId := c.Param("channel_id")
	data := &payloadDetails{ApiKey: config.GCYTAPI3Key, ChannelId: channelId}
	fmt.Println(data)
	pythonAPIUrl := "http://127.0.0.1:5000/list/YT"
	payload, _ := json.Marshal(data)
	resp, err := http.Post(pythonAPIUrl, "application/json", bytes.NewReader(payload))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, map[string]interface{}{"status": 400, "error": "Internal Server Error", "msg": "Python Api error"})
		return
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var respObjs []map[string]interface{}
	decoder.Decode(&respObjs)

	c.JSON(http.StatusOK, respObjs)
}

// @Summary      list analyse detail of video comments
// @Description  Returns Sentiment analyse of all video comments
// @Tags         Yt
// @Produce      json
// @Param video_id path  string true "videoID"
// @Success      200  {string}  string "pong"
// @Router       /yt/analyse/video/{video_id} [get]
func AnalyseVideo(c *gin.Context) {
	videoId := c.Param("video_id")
	channelId := c.Param("channel_id")
	pythonAPIUrl := fmt.Sprintf("http://127.0.0.1:5000/YT/analyse/%s/%s", channelId, videoId)
	resp, err := http.Get(pythonAPIUrl)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{"status": 400, "error": "Internal Server Error", "msg": "Python Api error"})
		return
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var respObjs []map[string]interface{}
	if err = decoder.Decode(&respObjs); err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusOK, map[string]interface{}{"status": 400, "error": "Internal Server Error", "msg": "Python Api error"})
		return
	}

	c.JSON(http.StatusOK, respObjs)
}
