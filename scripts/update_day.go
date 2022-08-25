package scripts

import (
	"cryptoguess/configs"
	"log"
	"math/rand"
	"time"
)

func UpdateToday() {
	err := configs.ArchiveToday()
	if err != nil {
		log.Panic("Cannot archive transactions", err)
	}
	currentTime := time.Now()
	configs.Today = currentTime

	generateTodaysCoins()
	log.Println("Coins Generated for today: ", configs.CurrentCoins)
}

func generateTodaysCoins() {

	market_coins, err := configs.GetMarketCoins()

	if err != nil {
		// TODO: Improve this error handling, because it can occur in
		// TODO: Different scenarios
		log.Panic("Could not get the market coins")
	}

	rand.Seed(time.Now().UnixNano())
	var todays_coins []configs.Coin

	for i := 0; i < 4; i++ {
		number := rand.Intn(100)
		todays_coins = append(todays_coins, market_coins[number])
	}

	configs.CurrentCoins = todays_coins
}

func UpdateCoinPrices() {
	//* A simple update script for updating prices
	updatedCoins, err := configs.GetSpecificCoins(configs.CurrentCoins)

	if err != nil {
		//TODO: Same as in update_day.go:27-28
		log.Panic("Could not get the price of market coins")
	}
	configs.CurrentCoins = updatedCoins
}
