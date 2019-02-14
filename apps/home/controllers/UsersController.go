package controllers

type UsersController struct{}

func (c *UsersController) Get() string {
	return "Get"
}

func (c *UsersController) GetTest() string {
	return "GetTest"
}

func (c *UsersController) PostTest() string {
	return "PostTest"
}

func (c *UsersController) GetTest2() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}
