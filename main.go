package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"stock.api/api/controllers"
	middlewares "stock.api/middleware"
)

func main() {

	var router = gin.Default()
	router.Use(middlewares.Cors())

	var address = ":9010"
	v1 := router.Group("/api/v1/")
	{
		v1.POST("/stock/user/register", controllers.CreateUser)
		v1.POST("/stock/stock/add", controllers.CreateStock)
		v1.GET("/stock/stock/list", controllers.GetStockList)
		v1.POST("/stock/stock/addData", controllers.CreateStockData)
		v1.GET("/stock/stock/getData", controllers.GetStockInfo)

	}
	log.Fatalln(router.Run(address))

}
