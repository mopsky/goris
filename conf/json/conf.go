package json

import (
	"encoding/json"
	"errors"
	"os"
)

func LoadConfig(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("打开配置文件 " + path + " 出错：" + err.Error())
	}

	fi, _ := file.Stat()
	if fi.Size() == 0 {
		return nil, errors.New("配置文件 " + path + " 为空")
	}

	buffer := make([]byte, fi.Size())
	_, err = file.Read(buffer)
	if err != nil {
		return nil, errors.New("配置文件 " + path + " 读取失败：" + err.Error())
	}

	buffer = []byte(os.ExpandEnv(string(buffer)))
	config := make(map[string]interface{})
	err = json.Unmarshal(buffer, &config) //解析json格式数据
	if err != nil {
		return nil, errors.New("配置文件 " + path + " JSON转换失败：" + err.Error())
	}
	return config, nil
}

// 读取数据库配置
func DataBaseConf() (map[string]interface{}, error) {
	return LoadConfig("conf/json/database.json")
}

// 读取redis配置
func RedisConf() (map[string]interface{}, error) {
	return LoadConfig("conf/json/redis.json")
}

// 读取业务配置
func BusinessConf() (map[string]interface{}, error) {
	return LoadConfig("conf/json/business.json")
}
