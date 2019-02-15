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
	userInfo := UserInfo{userID: 27, userPhone: "18950295811", loginName: "yangxb", fullName: "杨夕兵"}
	c.Session.Set("UserInfo", userInfo)
	return "Login"
}

func (c *UsersController) GetTest2() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}
