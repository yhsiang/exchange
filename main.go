package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yhsiang/exchange/apis"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(apis.Database())
	r.GET("/", apis.Index)
	r.POST("/order", apis.CreateOrder)
	r.POST("/book", apis.QueryOrderBook)

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
