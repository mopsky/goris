package controllers

import (
	"fmt"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type BaseController struct {
	Session *sessions.Session
}

func (c *BaseController) BeforeActivation(b mvc.BeforeActivation) {
	b.Router()
	b.Handle("GET", "/custom", "Custom")
	fmt.Println(1)

}

func (c *BaseController) AfterActivation(a mvc.AfterActivation) {
	if a.Singleton() {
		panic("basicController should be stateless,a request-scoped,we have a 'Session' which depends on the context.")
	}
}

func (c *BaseController) Get() string {

	body := fmt.Sprintf("Hello from basicController\nTotal visits from you: %d", 22)
	return body
}

func (c *BaseController) Custom() string {
	//Type := reflect.TypeOf(c)

	//n := Type.NumMethod()
	//for i := 0; i < n; i++ {
	//	m := Type.Method(i)
	//	fmt.Println(m)
	//}

	//fmt.Println(Type.MethodByName("Yangxb"))
	return "custom"
}

func (c *BaseController) Yangxb() string {

	return "yangxb"
}