package reward

import "time"

type Reward struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userId"`
	StockSymbol string    `json:"stockSymbol"`
	Shares      float64   `json:"shares"`
	RewardedAt  time.Time `json:"rewardedAt"`
}

type StockPrice struct {
	StockSymbol string  `json:"stockSymbol"`
	PriceINR    float64 `json:"priceInInr"`
}
