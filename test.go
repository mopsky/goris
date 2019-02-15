package main

import (
	"fmt"
	"github.com/goris/conf/yaml"
	"github.com/goris/kernel/db"
	"strconv"
)

type MyTest struct {
	a int    `json:"aaaa"`
	b string `json:"bbbb"`
}

func (m *MyTest) MyTest() {
	m.a = 1
}

func (test *MyTest) set(a int, b string) {
	test.a = a
	test.b = b
}

type MyChildTest struct {
	MyTest
}

func main() {
	//读取数据库配置
	c, err := yaml.DataBaseConf()
	if err != nil {
		fmt.Println(err)
	}

	//初始化数据库
	host, port, database, user, pass := c.String("host"), c.Int("port"), c.String("database"), c.String("user"), c.String("pass")
	yaml.DATASOURCE = user + ":" + pass + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database + "?charset=utf8"

	//	test := new(MyChildTest)
	////	test.set(1,"2")
	//	test.MyTest.MyTest()
	//	fmt.Println(test)

	//Query("SELECT login_name FROM user WHERE user_id = ? and user_id = ? and login_name = ?", args...)
	//where := []Where{
	//	{"user_id", "like", "2%"},
	//	{"login_name", "like", "%%"},
	//}
	//
	//test, err := M("user").Where(where).Where("user_id is not null").Where("user_id > 0").Limit(10).Order("   user_id   ").Select()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(test[0]["login_name"], len(test))
	//
	test2, err := db.M("user").Query("select * from user where user_id = 27")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test2)
	//
	//test3, err := M("user").Where("user_id < 100").Limit(10).Find()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(test3)
	//
	//test4, err := M("user").Where("user_id < 100").GetField("user_id, login_name")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(44444, test4)
	//
	//test5, err := M("user").Where("user_id < 100").Count()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(555555, test5)
	//
	//test6, err := M("user").Where("user_id < 100").Min("user_id")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(6666, test6)
	//
	//test7, err := M("user").Where("user_id < 100").Max("user_id")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(7777, test7)
	//
	//test8, err := M("user").Where("user_id < 100").Avg("user_id")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(8888, test8)

	//data := map[string]string{
	//	"ip": "12345",
	//	"user_id": "33",
	//	"url" : "/test",
	//	"in_param" : "aa",
	//	"out_param": "bb",
	//}
	//test9, err := M("cad_logs").Add(data)
	//fmt.Println(test9, err)

	//s := []map[string]interface{}{}
	//
	//m1 := map[string]interface{}{"name": "John", "age": 10}
	//m2 := map[string]interface{}{"name": "Alex", "age": 12}
	//
	//s = append(s, m1, m2)
	//s = append(s, m2)
	//
	//s := map[int]map[int]map[string]map[string]string{
	//	0: {
	//		0: {
	//			"test1" : {
	//				"test11": "test11",
	//			},
	//		},
	//	},
	//	1: {
	//		0: {
	//			"test2" : {
	//				"test22": "test22",
	//			},
	//		},
	//	},
	//}
	//
	//b, err := json.Marshal(s)
	//if err != nil {
	//	fmt.Println("json.Marshal failed:", err)
	//	return
	//}
	//
	//
	//fmt.Println("b:", string(b))

	//test := []byte("{\"0\":{\"0\":{\"test1\":{\"test11\":\"test11\"}}},\"1\":{\"0\":{\"test2\":{\"test22\":\"test22\"}}}}")
	//m := make(map[int]map[int]map[string]map[string]string)
	//
	//errx := json.Unmarshal(test, &m)
	//if errx != nil {
	//
	//	fmt.Println("Umarshal failed:", errx)
	//	return
	//}
	//
	//
	//fmt.Println("m:", m)

	//datas := []map[string]string{
	//	{
	//		"ip": "ip1",
	//		"user_id": "1",
	//		"url" : "/url1",
	//		"in_param" : "in_param1",
	//		"out_param": "out_param1",
	//	},
	//	{
	//		"ip": "ip2",
	//		"user_id": "2",
	//		"url" : "/url2",
	//		"in_param" : "in_param2",
	//		"out_param": "out_param2",
	//	},
	//}
	//
	//test10, err := M("cad_logs").AddAll(datas)
	//fmt.Println(test10, err)

	//data := map[string]string{
	//	"ip": "ip1",
	//	"user_id": "3",
	//	//"url" : "/url1",
	//	//"in_param" : "in_param1",
	//	//"out_param": "out_param1",
	//}
	//
	//test11, err := M("cad_logs").Where("log_id = 1698").Save(data)
	//fmt.Println(test11, err)

	//test12, err := M("cad_logs").Where("log_id = 1698").Delete()
	//fmt.Println(test12, err)

	//test13, err := M("cad_logs").Where("log_id = 1698").SetInc("user_id", -1)
	//fmt.Println(test13, err)

	//for i := 0; i <10000; i++ {
	//	m := M("cad_logs").Where("aaaaaaaaaaaaaaaaaa")
	//	fmt.Println(i, m.GetError())
	//	m.Close()
	//	//time.Sleep(time.Millisecond)
	//}

	//a := models.NewUser()
	//defer a.Close()
	////b, err := a.Where("user_id = 27").Find()
	//c, err2 := a.GetById(27)
	//fmt.Println(c, err2)

	//a, b := http.Get("")
	//for i := 0; i < 2; i++ {
	//	res, err := curl.Get("http://tool.bitefu.net/jiari/vip.php?d=201901&type=0&apikey=123456", true)
	//	//fmt.Println(res)
	//	aaa, ok := res.(map[string]interface{})["status"]
	//	fmt.Println(reflect.TypeOf(aaa), ok, err)
	//}

	//header := []curl.Header{
	//	{"Content-Type", "application/x-www-form-urlencoded; charset=UTF-8"},
	//	{"X-Requested-With", "XMLHttpRequest"},
	//}
	//fmt.Println(header)
	//res2, err2 := curl.Post("http://192.168.1.230/Shop-ajaxShopCartgory", header, true)
	//fmt.Println(res2, err2)

	//console.Log("test", 123, 321)
	//console.Log("22222222222222")
	//
	//redis := new(cache.Redis)
	//defer redis.Close()
	//err := redis.Open("192.168.1.231", 22122)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//type TestStruct struct {
	//	Yangxb string
	//	Test int
	//	Test2 float64
	//}
	//
	//err = redis.Set("yangxb", 1111, 30)
	//fmt.Println(err)
	//
	//res, err2 := redis.Keys("yangxb12", true)
	//fmt.Println(res, err2)

	//t := interface{}(1)
	//fmt.Println(t.(int))
	//
	//x := interface{}(&struct { a int; b string}{a: 1, b: "c"})
	//fmt.Println(x.(*struct{a int; b string}))
	//
	//type y struct { a int; b string}
	//fmt.Println(x.(*y))
}
