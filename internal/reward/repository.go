package reward

import (
	"database/sql"
)

func InsertReward(db *sql.DB, r Reward) error {
	_, err := db.Exec(`
		INSERT INTO rewards (user_id, stock_symbol, shares, rewarded_at)
		VALUES ($1, $2, $3, NOW())`,
		r.UserID, r.StockSymbol, r.Shares)
	return err
}

func GetTodayRewards(db *sql.DB, userID int) ([]Reward, error) {
	rows, err := db.Query(`
		SELECT id, user_id, stock_symbol, shares, rewarded_at
		FROM rewards
		WHERE user_id = $1 AND DATE(rewarded_at) = CURRENT_DATE`,
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rewards []Reward
	for rows.Next() {
		var r Reward
		rows.Scan(&r.ID, &r.UserID, &r.StockSymbol, &r.Shares, &r.RewardedAt)
		rewards = append(rewards, r)
	}
	return rewards, nil
}

func GetStats(db *sql.DB, userID int) (map[string]float64, float64, error) {
	// Get today’s total shares grouped by stock
	todayShares := make(map[string]float64)

	rows, err := db.Query(`
		SELECT stock_symbol, SUM(shares)
		FROM rewards
		WHERE user_id = $1 AND DATE(rewarded_at) = CURRENT_DATE
		GROUP BY stock_symbol`, userID)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var symbol string
		var total float64
		rows.Scan(&symbol, &total)
		todayShares[symbol] = total
	}

	// Get total portfolio value
	var totalINR float64
	err = db.QueryRow(`
		SELECT COALESCE(SUM(r.shares * sp.price_in_inr), 0)
		FROM rewards r
		JOIN stock_prices sp ON r.stock_symbol = sp.stock_symbol
		WHERE user_id = $1`, userID).Scan(&totalINR)

	return todayShares, totalINR, err
}

func GetHistoricalINR(db *sql.DB, userID int) ([]map[string]interface{}, error) {
	rows, err := db.Query(`
		SELECT DATE(r.rewarded_at) AS date,
		       SUM(r.shares * sp.price_in_inr) AS total_inr
		FROM rewards r
		JOIN stock_prices sp ON sp.stock_symbol = r.stock_symbol
		WHERE r.user_id = $1
		  AND DATE(r.rewarded_at) < CURRENT_DATE
		GROUP BY DATE(r.rewarded_at)
		ORDER BY DATE(r.rewarded_at) DESC`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var date string
		var total float64
		rows.Scan(&date, &total)
		results = append(results, map[string]interface{}{
			"date":             date,
			"totalValueInINR":  total,
		})
	}

	return results, nil
}
