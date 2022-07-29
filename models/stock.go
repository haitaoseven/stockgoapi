package models

type Stock struct {
	Id   int64
	Name string `xorm:"name" json:"name"`
	Code string `xorm:"code" json:"code"`
}
