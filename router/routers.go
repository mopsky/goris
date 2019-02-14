package router

/**
  路由器
*/

import (
	AdminControllers "github.com/goris/apps/admin/controllers"
	HomeControllers "github.com/goris/apps/home/controllers"
	"github.com/kataras/iris/mvc"
)

/* Home 控制器 */
func HomeRouters(app *mvc.Application) {
	// Add the authentication(admin:password) middleware
	//app.Router.Use(middleware.BasicAuth)

	//当然，你可以在MVC应用程序中使用普通的中间件。
	//app.Router.Use(func(ctx iris.Context) {
	//	fmt.Println(1111,ctx,2222)
	//	ctx.Application().Logger().Infof("Path: %s", ctx.Path())
	//	ctx.Next()
	//})

	/*主路由*/
	app.Handle(new(HomeControllers.IndexController))

	/*子路由*/

	//测试控制器
	app.Party("/test").Handle(new(HomeControllers.ExampleController))

	//用户控制器
	app.Party("/users").Handle(new(HomeControllers.UsersController))
}

/* Admin 控制器 */
func AdminRouters(app *mvc.Application) {
	// Add the authentication(admin:password) middleware
	//app.Router.Use(middleware.BasicAuth)

	/*主路由*/
	app.Handle(new(AdminControllers.IndexController))

	/*子路由*/
	app.Party("/about").Handle(new(AdminControllers.AboutController))
}
