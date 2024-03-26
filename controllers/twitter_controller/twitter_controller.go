package twittercontroller

import (
	"fmt"
	"net/http"

	"github.com/SoyebSarkar/content-creator-insight/datasource/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitterv2"
)

func Twtlogin(c *gin.Context) {
	provid := twitterv2.New(config.TwitterClientID, config.TwitterClientSecret, config.TwtCallBack)

	store := sessions.NewCookieStore([]byte("sec-cookei"))
	// sessions.coo
	store.Options.MaxAge = 30 * 100 // Set the MaxAge of the session
	gothic.Store = store            // Set the store for gothic
	q := c.Request.URL.Query()
	q.Add("provider", "twitterv2")
	c.Request.URL.RawQuery = q.Encode()
	goth.UseProviders(provid)

	url, err := gothic.GetAuthURL(c.Writer, c.Request)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, url)
}

func HandleTwtAuth(c *gin.Context) {
	provid := twitterv2.New(config.TwitterClientID, config.TwitterClientSecret, config.TwtCallBack)

	store := sessions.NewCookieStore([]byte("sec-cookei"))
	// sessions.coo
	store.Options.MaxAge = 30 * 100 // Set the MaxAge of the session
	gothic.Store = store            // Set the store for gothic
	q := c.Request.URL.Query()
	q.Add("provider", "twitterv2")
	c.Request.URL.RawQuery = q.Encode()
	goth.UseProviders(provid)
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, user)
}
