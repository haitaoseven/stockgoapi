package controllers

import (
	//"database/sql"

	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"stock.api/common"
	"stock.api/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	var response common.XormResponse

	err := c.Bind(&user)
	user.CreatedAt = time.Now()
	if err != nil {
		response.Code = http.StatusBadGateway
		response.Msg = "参数错误"
		response.Data = "error"
		c.JSON(http.StatusOK, response)
		return
	}

	affected, err := common.Dbcon.Insert(&user)

	if err != nil || affected <= 0 {
		response.Code = http.StatusBadGateway
		response.Msg = "添加数据失败"
		response.Data = err
		c.JSON(http.StatusOK, response)
		return
	}

	response.Code = http.StatusOK
	response.Msg = "添加数据成功"
	response.Data = user
	c.JSON(http.StatusOK, response)
	fmt.Println(affected)
	return
}
