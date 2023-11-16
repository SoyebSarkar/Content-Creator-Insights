package users

import (
	"fmt"
	"net/http"

	"github.com/SoyebSarkar/content-creator-insight/domain/auth"
	userdto "github.com/SoyebSarkar/content-creator-insight/domain/user"
	service "github.com/SoyebSarkar/content-creator-insight/services"
	userServices "github.com/SoyebSarkar/content-creator-insight/services"
	"github.com/gin-gonic/gin"
)

// @Summary      Creates a new user
// @Description  Creates a new user from signup
// @Tags         users
// @Produce      json
// @Param user body  userdto.User true "user details"
// @Success      200  {string}  string "pong"
// @Router      /signup [post]
func CreateUser(c *gin.Context) {
	var user userdto.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, "Json error")
		return
	}
	err := userServices.CreateUser(user)
	if err != nil {
		fmt.Println("-->", err)
		c.JSON(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]int{"status": 200})
}

// @Summary     Validate login
// @Description  validate login with credentials
// @Tags         users
// @Produce      json
// @Param login body  auth.Login true "login details"
// @Success      200  {string}  string "pong"
// @Router      /login [post]
func LoginUser(c *gin.Context) {
	var loginAuth auth.Login
	if err := c.ShouldBindJSON(&loginAuth); err != nil {
		c.JSON(http.StatusOK, "invalid json")
		return
	}
	result, err := service.LoginValidate(&loginAuth)
	if err != nil || result == false {
		fmt.Println(err)
		c.JSON(http.StatusOK, "Invalid Credentials")
		return
	}
	c.JSON(http.StatusOK, map[string]int{"status": 200})

}
