package yaml

import (
	"errors"
	"github.com/goris/utils/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadConfig(path string) (*types.TMap, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("打开配置文件 " + path + " 出错：" + err.Error())
	}

	//config := make(CT)
	var config types.TMap
	err = yaml.Unmarshal(yamlFile, &config.Value)
	if err != nil {
		return nil, errors.New("读取配置文件 " + path + " 出错：" + err.Error())
	}

	return &config, nil

}

// 读取数据库配置
func DataBaseConf() (*types.TMap, error) {
	return LoadConfig("conf/yaml/database.yaml")
}

// 读取redis配置
func RedisConf() (*types.TMap, error) {
	return LoadConfig("conf/yaml/redis.yaml")
}

// 读取业务配置
func BusinessConf() (*types.TMap, error) {
	return LoadConfig("conf/yaml/business.yaml")
}
