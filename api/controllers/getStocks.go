package controllers

import (
	//"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"stock.api/common"
	"stock.api/models"
)

func GetStockList(c *gin.Context) {
	var response common.XormResponse
	uid := 1
	var userStockData = make([]*models.UserStockData, 0)
	total, err := common.Dbcon.Table("user_stock").Join("LEFT", "stock", "user_stock.sid = stock.id").Where("user_stock.uid=?", uid).Count()
	if err == nil {
		response.Code = http.StatusBadRequest
	}

	var paginationResult models.UserStockPaginiationData

	var pagination models.Pagination
	c.ShouldBindQuery(&pagination)
	pageSize := pagination.PageSize
	page := pagination.Page
	// common.Dbcon.Table("user_stock").Join("LEFT", "stock", "user_stock.sid = stock.id").Where("user_stock.uid=?", uid).Limit(, pageSize).Find(&userStockData)
	common.Dbcon.Table("user_stock").Join("LEFT", "stock", "user_stock.sid = stock.id").Where("user_stock.uid=?", uid).Limit(pageSize, pageSize*(page)).Find(&userStockData)

	response.Code = http.StatusOK
	paginationResult.Count = total
	paginationResult.Data = userStockData
	paginationResult.Pagination = pagination
	response.Data = paginationResult
	c.JSON(http.StatusOK, response)

}
