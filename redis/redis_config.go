package redis

import (
	"encoding/json"
	"io/ioutil"
	"time"
	"common-go/log"
)

type RedisConfig struct {
	MaxIdle int //最大空闲连接

	IdleTimeout time.Duration //空闲连接超时时间

	Password string

	Host string //主机域名/ip
}

//实例
func Init(maxIdle int, idleTimeout time.Duration, password, host string) *RedisConfig {
	return &RedisConfig{MaxIdle: maxIdle, IdleTimeout: idleTimeout, Password: password, Host: host}
}

//通过加载json配置文件生成实例
func LoadConfigFile(uri string) []RedisConfig {

	configData, e := ioutil.ReadFile(uri)

	if e != nil {
		panic(uri + " file read fail:" + e.Error())
	}
	configs := make([]RedisConfig, 0)
	err := json.Unmarshal(configData, &configs)
	if err != nil {
		log.Logger.Error("load json redis config fail:", err)
	}
	RedisConfigs = configs

	return configs
}
