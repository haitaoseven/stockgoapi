package models

type UserStockData struct {
	Id        int64 `json:"id"`
	UserStock `xorm:"extends"`
	Stock     `xorm:"extends"`
}

type UserStockPaginiationData struct {
	Pagination
	Data  []*UserStockData `json:"data"`
	Count int64            `json:"count"`
}

type Pagination struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type StockDatePriceData struct {
	Date       string  `json:"date"`
	FinalPrice float64 `json:"finalPrice"`
}
