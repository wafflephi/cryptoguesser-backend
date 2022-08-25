package controllers

import (
	"cryptoguess/configs"
	"cryptoguess/models"
	"cryptoguess/responses"
	"math/rand"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Make this only for admins, and assign session by default to
		//TODO: authenticate users or maybe show captcha?

		var auth models.Auth
		//* Checks the request for its validity
		if err := c.BindJSON(&auth); err != nil {
			c.JSON(http.StatusBadRequest, responses.AuthResponse{Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if validationErr := validate.Struct(&auth); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AuthResponse{Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		//* Gets the ID for the session
		getID := func(low, hi int) int {
			return low + rand.Intn(hi-low)
		}

		session := sessions.Default(c)
		sessionID := session.Get("id")
		if sessionID != nil {
			//* Check if the user is logged in
			c.JSON(http.StatusConflict, responses.AuthResponse{Message: "error", Data: map[string]interface{}{"data": "User already logged in"}})
			return
		}
		session.Set("id", getID(1, 10000))
		//TODO: Make this a better system !
		if auth.Username == "admin" && auth.Password == configs.EnvAdminPassword() {
			session.Set("admin", true)
		} else {
			session.Set("admin", false)
		}
		session.Save()
		c.JSON(http.StatusOK,
			responses.AuthResponse{Message: "success"},
		)
	}
}

func AuthLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.JSON(http.StatusOK, responses.AuthResponse{Message: "Logout successful"})
	}
}

func AuthTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("id")
		isAdmin := session.Get("admin")
		if sessionID == nil {
			c.JSON(http.StatusAccepted, responses.AuthResponse{Message: "User not logged in"})
			c.Abort()
		} else if isAdmin == true && sessionID != nil {
			c.JSON(http.StatusOK, responses.AuthResponse{Message: "User logged in as admin"})
		} else {
			c.JSON(http.StatusOK, responses.AuthResponse{Message: "User logged in"})
		}
	}
}
