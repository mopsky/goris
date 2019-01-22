package router

/**
  路由器
 */

import (
	"github.com/kataras/iris/mvc"
	AdminControllers "pfws.go/apps/admin/controllers"
	HomeControllers "pfws.go/apps/home/controllers"
)

/* Home 控制器 */
func HomeRouters(app *mvc.Application) {
	// Add the authentication(admin:password) middleware
	//app.Router.Use(middleware.BasicAuth)

	/*主路由*/
	app.Handle(new(HomeControllers.IndexController))

	/*子路由*/
	app.Party("/about").Handle(new(HomeControllers.AboutController))
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