package controllers

type Uuhhser struct {
	Id   int    `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"unique" json:"name"`
	Pass string `"json:"pass"`
}
