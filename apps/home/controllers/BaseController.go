package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

// 用户结构
type UserInfo struct {
	userID    int
	userPhone string
	userLogo  string
	loginName string
	fullName  string
}

// 店铺结构
type ShopInfo struct {
	shopID   int
	shopName string
	shopLogo string
}

// 基类控制器
type BaseController struct {
	//mvc.BaseController
	Session  *sessions.Session
	userInfo UserInfo
	shopInfo ShopInfo
	userID   int
	shopID   int
}

// BeginRequest initializes the current user's Session.
func (c *BaseController) BeginRequest(ctx iris.Context) {
	fmt.Println("BeginRequest", ctx)
	c.userInfo.fullName = "样细胞"
	//c.Session = s.Manager.Start(ctx)
}

// EndRequest is here to complete the `BaseController`.
func (c *BaseController) EndRequest(ctx iris.Context) {
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
