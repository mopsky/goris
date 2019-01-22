package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"pfws.go/router"
)

const (
	publicDir = "./public"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./apps/home/views", ".html").Layout("shared/layout.html"))
	app.StaticWeb("/public", publicDir)
	//app.OnAnyErrorCode(onError)
	//mvc.New(app.Party("/")).Handle(new(HomeControllers.BaseController))
	//mvc.New(app.Party("/admin")).Handle(new(AdminControllers.IndexController))

	mvc.Configure(app.Party("/"), router.HomeRouters)
	mvc.Configure(app.Party("/admin"), router.AdminRouters)

	app.Run(iris.Addr(":8080"))
}

func handler(ctx iris.Context){
	fmt.Println(ctx)
	ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
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
