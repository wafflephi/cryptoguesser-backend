package controllers

import (
	"cryptoguess/configs"
	"cryptoguess/responses"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ResourceGetCoinsToday() gin.HandlerFunc {
	return func(c *gin.Context) {
		//* Return the daily coins for the app
		coins := configs.CurrentCoins
		var coins_list []responses.DailyCoinResponse
		for i := range coins {
			coin_name := coins[i].Name
			coin_symbol := coins[i].Symbol
			coin_price := coins[i].CurrentPrice
			coins_list = append(coins_list, responses.DailyCoinResponse{Name: coin_name, Symbol: coin_symbol, Price: coin_price})
		}

		c.JSON(http.StatusOK, coins_list)
	}
}

func ResourceGetArchiveFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		//* Returns the archival file
		//! ( For now only to admins !)

		File := struct {
			Name string `uri:"file"`
		}{}

		if err := c.BindUri(&File); err != nil {
			c.JSON(http.StatusBadRequest, responses.ResourcesResponse{Message: "Cannot bind URI", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		session := sessions.Default(c)
		if session.Get("id") == nil {
			c.JSON(http.StatusForbidden, responses.ResourcesResponse{Message: "You need to be logged in to access archive"})
			return
		}
		if session.Get("admin ") == false {
			c.JSON(http.StatusForbidden, responses.ResourcesResponse{Message: "You need to be admin to access the archive"})
			return
		}

		//? Do we need a input validator to prevent security bugs ?
		fileContent, err := ioutil.ReadFile("./archive/" + File.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ResourcesResponse{Message: "Cannot get the provided file", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.Header("Content-Disposition", "attachment; filename="+File.Name)
		c.Data(http.StatusOK, "application/octet-stream", fileContent)
	}
}
