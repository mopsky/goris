package controllers

import (
	"github.com/kataras/iris/mvc"
)

type AboutController struct{}

var aboutView = mvc.View{
	Name: "about.html",
	Data: map[string]interface{}{
		"Title":   "About admin",
		"Message": "Your application description page..",
	},
}

func (c *AboutController) Get() mvc.View {
	return aboutView
}


func (c *AboutController) GetTest() string {
	return "GetTest"
}

func (c *AboutController) PostTest() string {
	return "PostTest"
}

func (c *AboutController) GetTest2() interface{}  {
	return map[string]string{"message": "Hello Iris!"}
}