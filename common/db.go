package common

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

var Dbcon *xorm.Engine

func init() {

	strStr := "stock:stock123@tcp(localhost:3306)/stock?charset=utf8" // /golang database name

	var err error
	Dbcon, err = xorm.NewEngine("mysql", strStr)

	if err != nil {
		fmt.Println("数据库连接失败:", err)
	}

}
