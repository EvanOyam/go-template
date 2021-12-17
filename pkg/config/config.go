package config

import (
	"io/ioutil"

	"github.com/fatih/structs"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Env      string `yaml:"Env"`      // 环境：prod、dev
	BaseUrl  string `yaml:"BaseUrl"`  // base url
	Port     string `yaml:"Port"`     // 端口
	LogFile  string `yaml:"LogFile"`  // 日志文件
	MySqlUrl string `yaml:"MySqlUrl"` // 数据库连接地址
}

var conf map[string]interface{}

func InitConfig(configPath string) (err error) {
	var temp = new(Config)
	if yamlFile, err := ioutil.ReadFile(configPath); err != nil {
		return err
	} else if err = yaml.Unmarshal(yamlFile, temp); err != nil {
		return err
	}
	conf = structs.Map(temp)
	return
}

func Get(key string) interface{} {
	return conf[key]
}
