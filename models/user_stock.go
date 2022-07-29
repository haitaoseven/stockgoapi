package models

type UserStock struct {
	Id  int64 `xorm:"id" json:"id"`
	Uid int64 `xorm:"uid" json:"uid"`
	Sid int64 `xorm:"sid" json:"sid"`
}
