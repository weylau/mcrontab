package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	EtcdServerAddr           string `json:"etcd_server_addr"`
	EtcdServerConnectTimeout int    `json:"etcd_server_connect_timeout"`
	EtcdServerOptionTimeout  int    `json:"etcd_server_option_timeout"`
}

var (
	G_config *Config
)

//加载配置
func LoadConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)

	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}
	G_config = &conf
	return
}
