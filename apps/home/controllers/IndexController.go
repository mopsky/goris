package controllers

import "github.com/kataras/iris/mvc"

type IndexController struct{}

func (c *IndexController) Get() mvc.Result {
	return mvc.View{Name: "index.html"}
}
