package models

import (
	"pfws.go/kernel/db"
)

type Shops struct{
	db.Model
}

func NewShops() *User {
	m := new(User)
	m.Open("shops")
	return m
}

