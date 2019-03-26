package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	HttpServerPort         int `json:"http_server_port"`
	HttpServerReadTimeout  int `json:"http_server_read_timeout"`
	HttpServerWriteTimeout int `json:"http_server_write_timeout"`
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
