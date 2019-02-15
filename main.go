package main

import (
	"github.com/goris/router"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

const (
	publicDir = "./public"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./apps/home/views", ".html").Layout("shared/layout.html"))
	//app.StaticWeb("/public", publicDir)
	//app.OnAnyErrorCode(onError)

	mvc.Configure(app.Party("/"), router.HomeRouters)
	mvc.Configure(app.Party("/admin"), router.AdminRouters)

	//homeApp := mvc.New(app.Party("/"))
	//homeApp.Register(session.Start, time.Now())
	//homeApp.Handle(new(controllers.ExampleController))

	app.Run(iris.Addr(":8080"))

}

//
//type err struct {
//	Title string
//	Code  int
//}
//
//func onError(ctx context.Context) {
//	ctx.ViewData("", err{"Error", ctx.GetStatusCode()})
//	ctx.View("shared/error.html")
//}
