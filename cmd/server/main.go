package main

import (
	"fmt"
	"time"

	"github.com/divyanshu/stocky/internal/db"
	"github.com/divyanshu/stocky/internal/price"
	"github.com/divyanshu/stocky/internal/reward"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Stocky server...")

	database := db.Connect()
	r := gin.Default()

	// Register routes
	reward.RegisterRoutes(r, database)

	// Start a goroutine to update prices every hour
	go func() {
		stocks := []string{"RELIANCE", "TCS", "INFY", "HDFCBANK"}
		for {
			price.UpdateStockPrices(database, stocks)
			time.Sleep(1 * time.Hour)
		}
	}()

	r.Run(":8080")
}
