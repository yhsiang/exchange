package apis

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yhsiang/exchange/database"
)

func Database() gin.HandlerFunc {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
