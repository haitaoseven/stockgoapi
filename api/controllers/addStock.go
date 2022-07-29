package controllers

import (
	//"database/sql"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"stock.api/common"
	"stock.api/models"
)

func CreateStock(c *gin.Context) {

	var stock models.Stock
	var response common.XormResponse

	err := c.Bind(&stock)
	if err != nil {
		response.Code = http.StatusBadGateway
		response.Msg = "参数错误"
		response.Data = "error"
		c.JSON(http.StatusOK, response)
		return
	}

	affected, err := common.Dbcon.Insert(&stock)

	if err != nil || affected <= 0 {
		response.Code = http.StatusBadGateway
		response.Msg = "添加数据失败"
		response.Data = err
		c.JSON(http.StatusOK, response)
		return
	}

	sid := stock.Id
	var userStock models.UserStock
	userStock.Uid = 1
	userStock.Sid = sid
	affectedUserStock, errUserStock := common.Dbcon.Insert(&userStock)
	if errUserStock != nil || affectedUserStock <= 0 {
		response.Code = http.StatusBadGateway
		response.Msg = "添加数据失败"
		response.Data = errUserStock
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = http.StatusOK
	response.Msg = "添加数据成功"
	response.Data = userStock
	c.JSON(http.StatusOK, response)
	fmt.Println(affected)
	return
}
