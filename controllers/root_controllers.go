package controllers

import (
	"cryptoguess/configs"
	"cryptoguess/models"
	"cryptoguess/responses"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func RootVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		//* Returns API Version and the name associated with it.
		c.JSON(http.StatusOK,
			responses.VersionResponse{Version: configs.API_Version, Version_name: configs.API_Version_name},
		)
	}
}

func RootUploadResult() gin.HandlerFunc {
	return func(c *gin.Context) {

		var request models.Transaction
		//* Checks the request for its validity
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, responses.TransactionResponse{Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if validationErr := validate.Struct(&request); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.TransactionResponse{Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		//? There should be a capcha of some sorts to prevent bots, and not the cookie session
		session := sessions.Default(c)
		sessionID := session.Get("id")
		if sessionID == nil {
			c.JSON(http.StatusForbidden, responses.TransactionResponse{Message: "You need to be logged in to vote"})
			return
		}

		if session.Get("admin") == false {
			if session.Get(request.Name) == "voted" {
				c.JSON(http.StatusConflict, responses.TransactionResponse{Message: "already voted on: " + request.Name})
				return
			} else {
				session.Set(request.Name, "voted")
				session.Save()
			}
		}

		transaction := configs.Transaction{
			Name:   request.Name,
			Price:  request.Price,
			Hour:   request.Hour,
			Action: *request.Action,
		}

		validTx := false
		for i := range configs.CurrentCoins {
			if configs.CurrentCoins[i].Name == request.Name {
				validTx = true
				break
			}
		}

		if !validTx {
			c.JSON(http.StatusBadRequest, responses.TransactionResponse{Message: "We aren't accepting this coin today"})
			return
		}

		configs.SaveTransaction(transaction)

		c.JSON(http.StatusOK, responses.TransactionResponse{Message: "Transaction saved"})
	}
}
