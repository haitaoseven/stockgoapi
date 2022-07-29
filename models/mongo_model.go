package models

type StockData struct {
	Name string `json:"name" bson:"name"`
	Code string `json:"code" bson:"code"`
	Date string `json:"date" bson:"date"`
	StockFields
}

type StockFields struct {
	OpenPrice  string `json:"openPrice" bson:"openPrice"`
	TopPrice   string `json:"topPrice" bson:"topPrice"`
	LowPrice   string `json:"lowPrice" bson:"lowPrice"`
	FinalPrice string `json:"finalPrice" bson:"finalPrice"`
}
