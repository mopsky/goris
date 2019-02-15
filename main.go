package main

import (
	"fmt"
	"github.com/goris/conf/yaml"
	"github.com/goris/router"
	"github.com/goris/utils/types"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

const (
	assetsDir = "assets"
)

//错误句柄
func onError(ctx iris.Context) {
	ctx.ViewData("", types.Err{"发生错误", ctx.GetStatusCode()})
	ctx.View("public/error.html")
}

//读取基本配置
func loadConfig() error {

	//读取数据库配置
	c, err := yaml.DataBaseConf()
	if err != nil {
		return err
	}

	host, port, database, user, pass := c.String("host"), c.Int("port"), c.String("database"), c.String("user"), c.String("pass")
	yaml.DB_SOURCE = user + ":" + pass + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database + "?charset=utf8"

	//读取redis配置
	c, err = yaml.RedisConf()
	if err != nil {
		return err
	}

	yaml.REDIS_SOURCE = c.String("host") + ":" + strconv.Itoa(c.Int("port"))

	//读取业务配置
	c, err = yaml.BusinessConf()
	if err != nil {
		return err
	}

	return nil
}

//主进程
func main() {
	//加载配置
	err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	app := iris.New()
	app.RegisterView(iris.HTML("apps/home/views", ".html").Layout("public/layout.html"))

	//静态文件
	app.StaticWeb("/static", assetsDir)
	app.OnAnyErrorCode(onError)

	mvc.Configure(app.Party("/"), router.HomeRouters)
	mvc.Configure(app.Party("/admin"), router.AdminRouters)

	app.Run(iris.Addr(":8080"))

}
