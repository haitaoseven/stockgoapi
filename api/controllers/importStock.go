package controllers

type Uusller struct {
	Id   int    `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"unique" json:"name"`
	Pass string `"json:"pass"`
}
