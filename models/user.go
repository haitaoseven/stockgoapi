package models

import "time"

type User struct {
	Id        int       `xorm:"id" json:"id"`
	Name      string    `xorm:"unique" json:"name"`
	Pass      string    `xorm:"pass" json:"pass"`
	CreatedAt time.Time `xorm:"created_at" json:"created_at"`
}
