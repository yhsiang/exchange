package apis

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yhsiang/exchange/models"
)

func QueryOrderBook(c *gin.Context) {
	var book models.OrderBookSymbol
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderbook := &models.OrderBook{
		Symbol: book.Symbol,
	}

	db := c.MustGet("DB").(*sql.DB)

	if err := orderbook.Query(db); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orderbook)

}
