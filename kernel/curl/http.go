package curl

import (
	"encoding/json"
	"github.com/kataras/iris/core/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Header struct {
	Name string
	Value string
}

// Get方法
func Get(url string, beJson ...interface{}) (interface{}, error) {
	if url == "" {
		return nil, errors.New("GET的URL不能为空")
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(beJson) > 0 && beJson[0].(bool) {
		var res interface{}
		err = json.Unmarshal(body, &res)
		return res, err
	}

	return string(body), nil
}

// Post方法
func Post(url string, header []Header, beJson ...interface{}) (interface{}, error) {
	if url == "" {
		return nil, errors.New("POST的URL不能为空")
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader("name=cjb"))
	if err != nil {
		return nil, err
	}

	for _, v := range header{
		req.Header.Set(v.Name, v.Value)
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(beJson) > 0 && beJson[0].(bool) {
		var res interface{}
		err = json.Unmarshal(body, &res)
		return res, err
	}

	return string(body), nil
}