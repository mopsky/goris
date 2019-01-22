package controllers

import "github.com/kataras/iris/mvc"

type IndexController struct{}

func (c *IndexController) Get() mvc.Result {
	return mvc.View{Name: "index.html"}
}

func (c *IndexController) GetAbout() mvc.Result {
	return mvc.View{
		Name: "about.html",
		Data: map[string]interface{}{
			"Title":   "About Admin",
			"Message": "Your application description page."},
	}
}

func (c *IndexController) GetContact() mvc.Result {
	return mvc.View{
		Name: "contact.html",
		Data: map[string]interface{}{
			"Title":   "Contact Page",
			"Message": "Your application description page."},
	}
}

func (c *IndexController) GetTest() mvc.Result {
	return mvc.View{
		Name: "contact.html",
		Data: map[string]interface{}{
			"Title":   "xxxx Page",
			"Message": "Your application description page."},
	}
}
