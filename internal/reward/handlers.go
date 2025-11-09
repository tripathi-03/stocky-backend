package reward

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	r.POST("/reward", func(c *gin.Context) {
		var req Reward
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		err := InsertReward(db, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record reward"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Reward recorded successfully",
			"reward":  req,
		})
	})

	r.GET("/today-stocks/:userId", func(c *gin.Context) {
		userID, _ := strconv.Atoi(c.Param("userId"))
		rewards, err := GetTodayRewards(db, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching rewards"})
			return
		}
		c.JSON(http.StatusOK, rewards)
	})

	r.GET("/stats/:userId", func(c *gin.Context) {
		userID, _ := strconv.Atoi(c.Param("userId"))
		todayShares, totalValue, err := GetStats(db, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching stats"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"todayShares":             todayShares,
			"currentPortfolioValueINR": totalValue,
		})
	})

	r.GET("/historical-inr/:userId", func(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userId"))
	data, err := GetHistoricalINR(db, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching historical INR data"})
		return
	}
	c.JSON(http.StatusOK, data)
})


}
