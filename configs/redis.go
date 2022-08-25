package configs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

var (
	ctx     = context.Background()
	rdb     *redis.Client
	counter = 0
)

type Transaction struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Action bool    `json:"action"`
	Hour   string  `json:"hour"`
}

func ConnectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func SaveTransaction(transaction Transaction) error {

	if rdb == nil {
		return errors.New("redis database not connected")
	}

	err := set(rdb, "tx"+fmt.Sprint(counter), transaction)
	if err != nil {
		return err
	}

	counter++
	return nil
}

func GetTransaction(transactionID int) (Transaction, error) {

	var transaction Transaction

	if rdb == nil {
		return Transaction{}, errors.New("redis database not connected")
	}

	err := get(rdb, "tx"+fmt.Sprint(transactionID), &transaction)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func GetAllTransactions() ([]Transaction, error) {
	//* Loops over the counter to find all transactions

	var transactions []Transaction

	if rdb == nil {
		return nil, errors.New("redis database not connected")
	}

	for i := 0; i < counter; i++ {
		transaction, err := GetTransaction(i)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func set(c *redis.Client, key string, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(ctx, key, p, 24*time.Hour).Err()
}

func get(c *redis.Client, key string, dest interface{}) error {
	p, err := c.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(p), dest)
}

func ClearAllTransactions() error {
	err := rdb.FlushAll(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}
