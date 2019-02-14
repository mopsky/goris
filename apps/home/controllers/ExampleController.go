package controllers

import (
	"context"
	"fmt"
)

type ExampleController struct {
	BaseController
	Ctx context.Context
}

func (c *ExampleController) Get() string {
	fmt.Println(c.Ctx)
	return "I am ExampleController" + c.userInfo.fullName
}

func (c *ExampleController) GetAaa() string {
	return "I am aaa"
}
