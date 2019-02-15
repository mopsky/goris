package controllers

import (
	"fmt"
	"github.com/goris/kernel/db"
	"github.com/kataras/iris/mvc"
	"time"
)

type ExampleController struct {
	BaseController
}

func (c *ExampleController) Get() string {
	//return "I am ExampleController" + c.userInfo.fullName
	//fmt.Println(c.ctx, "xxxxxxxxxxxxxxxxxxxxxxxxx")
	visits := c.Session.Increment("visits", 1)
	fmt.Println(c.Session.Get("UserInfo"))
	// write the current, updated visits.
	since := time.Now().Sub(c.StartTime).Seconds()
	return fmt.Sprintf("%d visit(s) from my current session in %0.1f seconds of server's up-time",
		visits, since)
}

func (c *ExampleController) GetView() mvc.Result {
	users, err := db.M("user").Where("user_id < 100").Limit(2).Select()
	fmt.Println(users, err)

	return mvc.View{
		Name: "test/test.html",
		Data: map[string]interface{}{
			"Title":   "Title",
			"Message": "Message",
			"Map":     map[int]string{1: "a", 2: "b"},
			"Users":   users,
		},
	}
}
