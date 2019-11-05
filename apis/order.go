package apis

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/yhsiang/exchange/models"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.OrderID = uuid.New().String()

	db := c.MustGet("DB").(*sql.DB)

	if err := order.CreateOrder(db); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"order_id": order.GetOrderId(),
	})
}
