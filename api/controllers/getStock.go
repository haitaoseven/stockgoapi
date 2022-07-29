package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"stock.api/common"
	"stock.api/models"
)

func GetStockInfo(c *gin.Context) {
	var response common.XormResponse
	code, ok := c.GetQuery("code")

	startDate, ok1 := c.GetQuery("startDate")
	endDate, ok2 := c.GetQuery("endDate")

	if ok && ok1 && ok2 {

		filter := bson.M{"code": bson.M{"$eq": code}, "date": bson.M{"$lte": endDate, "$gte": startDate}}

		var stockInfo = make([]*models.StockDatePriceData, 0)

		cursor, err := common.StockMogoCollection.Find(c, filter)
		if err != nil {
			panic(err)
		}
		defer cursor.Close(c)
		for cursor.Next(c) {
			var episode models.StockDatePriceData
			if err = cursor.Decode(&episode); err != nil {
				log.Fatal(err)
			}
			stockInfo = append(stockInfo, &episode)
		}
		response.Code = http.StatusOK
		response.Data = stockInfo
		c.JSON(http.StatusOK, response)

	}

}
