package price

import (
	"database/sql"
	"log"
	"math/rand"
	"time"
)

func UpdateStockPrices(db *sql.DB, stocks []string) {
	for _, s := range stocks {
		price := 1000 + rand.Float64()*500 
		_, err := db.Exec(`
			INSERT INTO stock_prices (stock_symbol, price_in_inr, updated_at)
			VALUES ($1, $2, NOW())
			ON CONFLICT (stock_symbol)
			DO UPDATE SET price_in_inr = EXCLUDED.price_in_inr, updated_at = NOW()`,
			s, price)
		if err != nil {
			log.Println("Error updating price for", s, ":", err)
		}
	}
	log.Println("Stock prices updated at", time.Now())
}
