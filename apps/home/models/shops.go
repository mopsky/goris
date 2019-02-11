package models

import (
	"github.com/goris/kernel/db"
)

type Shops struct {
	db.Model
}

func NewShops() *User {
	m := new(User)
	m.Open("shops")
	return m
}
