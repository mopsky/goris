package controllers

import (
	"fmt"
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

func (c *ExampleController) GetAaa() string {
	return "I am aaa"
}
