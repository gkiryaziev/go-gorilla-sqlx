package model

import (
	"time"
)

type User struct {
	Id         int64     `db:"id" json:"id"`
	FirstName  string    `db:"first_name" json:"first_name"`
	LastName   string    `db:"last_name" json:"last_name"`
	MiddleName string    `db:"middle_name" json:"middle_name"`
	Dob        time.Time `db:"dob" json:"dob"`
	Address    string    `db:"address" json:"address"`
	Phone      string    `db:"phone" json:"phone"`
	Login      string    `db:"login" json:"login"`
	Password   string    `db:"password" json:"password"`
}
