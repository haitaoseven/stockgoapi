package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"stock.api/common"
	"stock.api/service"
)

func CreateStockData(c *gin.Context) {
	code := "600367"
	startDate := "20220601"
	endDate := "20220602"
	var response common.XormResponse

	contents := service.GetOnlineStockDataAndSaveToMongo(code, startDate, endDate)

	pdata, errors := service.ParseNeteaseData(contents)
	if errors != nil {
		fmt.Printf("ffffffff")

		fmt.Println(errors)
	}

	docs := []interface{}{}

	for _, v := range *pdata {
		fmt.Println(v)
		// open := common.Float64(v.Open)
		docs = append(docs,
			bson.D{{"Name", "aaa"}, {"teasfxt", "No bytefst insert a document, in MongoDB"}},
		)
	}
	// docsa := []interface{}{
	// 	bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}},
	// 	bson.D{{"title", "Showcasing a Blossoming Binary"}, {"text", "Binary data, safely stored with GridFS. Bucket the data"}},
	// }

	//common.StockMogoCollection.InsertMany(c, docs)

	// fmt.Printf("%s", contents)
	response.Code = http.StatusOK
	response.Data = pdata
	c.JSON(http.StatusOK, response)
	// var stockData models.StockData
	// stockData.Name = "重庆啤酒"
	// stockData.FinalPrice = "10"
	// stockData.Code = "600001"
	// stockData.Date = "2022-07-01"
	// common.StockMogoCollection.InsertOne(c, stockData)
}
