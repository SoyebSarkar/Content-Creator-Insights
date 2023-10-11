package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks             godoc
// @Summary      Get ping response
// @Description  Returns pong , just to check server is running
// @Tags         Ping
// @Produce      json
// @Success      200  {string}  string "pong"
// @Router       / [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
