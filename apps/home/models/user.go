package models

import (
	"pfws.go/kernel/db"
	"strconv"
)

type User struct{
	db.Model
}

/** 初始化构造函数*/
func NewUser() *User {
	m := new(User)
	m.Open("user")
	return m
}

// 示例1 *****
func (m *User) GetById(iUserID int) (interface{}, interface{}) {
	where := "user_id = " + strconv.Itoa(iUserID)
	return m.Where(where).Find()
}

// 示例2 *****
func (m *User) GetParams(param []db.Where) (interface{}, interface{}) {
	return m.Where(param).Select()
}
