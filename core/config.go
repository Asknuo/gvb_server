package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/globel"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func InitConfig() {
	c := &config.Config{}
	// 1. 读取配置文件
	yamlFile, err := ioutil.ReadFile("settings.yaml")
	if err != nil {
		fmt.Println("yamlFile.Get err   ", err)
	}
	// 2. 解析配置文件
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	globel.Config = c
}
