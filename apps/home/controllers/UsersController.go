package controllers

type UsersController struct {
	BaseController
}

func (c *UsersController) Get() string {
	return "Get"
}

func (c *UsersController) GetTest() string {
	return "GetTest"
}

func (c *UsersController) PostTest() string {
	return "PostTest"
}

func (c *UsersController) GetLogin() string {
	userInfo := UserInfo{UserID: 27, UserPhone: "18950295811", LoginName: "yangxb", FullName: "杨夕兵"}
	c.Session.Set("UserInfo", userInfo)
	return "Login"
}

func (c *UsersController) GetTest2() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}
