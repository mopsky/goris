package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"time"
)

// 用户结构
type UserInfo struct {
	UserID    int
	UserPhone string
	UserLogo  string
	LoginName string
	FullName  string
}

// 店铺结构
type ShopInfo struct {
	ShopID   int
	ShopName string
	ShopLogo string
}

// 基类控制器
type BaseController struct {
	Session   *sessions.Session
	StartTime time.Time
	ctx       iris.Context
	UserInfo  UserInfo
	ShopInfo  ShopInfo
	UserID    int
	ShopID    int
}

// BeginRequest initializes the current user's Session.
func (c *BaseController) BeginRequest(ctx iris.Context) {
	c.ctx = ctx
}

// EndRequest is here to complete the `BaseController`.
func (c *BaseController) EndRequest(ctx iris.Context) {
	//
}

// 登录判断
func (c *BaseController) CheckLogin() {
	//
}

//func (c *BaseController) Post() {
//	fmt.Println("Post")
//}
//func (c *BaseController) Put() {
//	fmt.Println("Put")
//}
//func (c *BaseController) Delete() {
//	fmt.Println("Delete")
//}
//func (c *BaseController) Connect() {
//	fmt.Println("Connect")
//}
//func (c *BaseController) Head() {
//	fmt.Println("Head")
//}
//func (c *BaseController) Patch() {
//	fmt.Println("Patch")
//}
//func (c *BaseController) Options() {
//	fmt.Println("Options")
//}
//func (c *BaseController) Trace() {
//	fmt.Println("Trace")
//}
//func (c *BaseController) All() string {
//	fmt.Println("All")
//	return "I am All"
//}
//func (c *BaseController) Any() string {
//	return "I am BaseController22" + c.userInfo.fullName
//}
