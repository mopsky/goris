package controllers

type ExampleController struct {}

func (c *ExampleController) Get() string {
	return "I am ExampleController"
}