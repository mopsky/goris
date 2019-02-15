package types

import (
	"strings"
)

//自定义类型库
type T map[interface{}]interface{}

//错误
type Err struct {
	Title string
	Code  int
}

type TMap struct {
	Value interface{}
}

func (t *TMap) Get(key string) interface{} {
	keys := strings.Split(key, ".")
	value := t.Value
	for _, v := range keys {
		switch value.(type) {
		case map[interface{}]interface{}:
			value = value.(map[interface{}]interface{})[v]
		case map[string]interface{}:
			value = value.(map[string]interface{})[v]
		}
	}
	return value
}

func (t *TMap) Int(key string) int {
	value := t.Get(key)
	return value.(int)
}

func (t *TMap) String(key string) string {
	value := t.Get(key)
	return value.(string)
}
