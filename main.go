package main

import (
	"fmt"
	"github.com/goris/conf/yaml"
	"github.com/goris/router"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

const (
	assetsDir = "assets"
)

//读取基本配置
func loadConfig() {

	//读取数据库配置
	c, err := yaml.DataBaseConf()
	if err != nil {
		fmt.Println(err)
	}

	host, port, database, user, pass := c.String("host"), c.Int("port"), c.String("database"), c.String("user"), c.String("pass")
	yaml.DATASOURCE = user + ":" + pass + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database + "?charset=utf8"

	//读取redis配置

	//读取业务配置
}

//主进程
func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("apps/home/views", ".html").Layout("public/layout.html"))

	//静态文件
	app.StaticWeb("/static", assetsDir)
	//app.OnAnyErrorCode(onError)

	//加载配置
	loadConfig()

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
