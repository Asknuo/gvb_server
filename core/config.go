package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const ConfigFile = "settings.yaml"

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
	global.Config = c
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	global.Log.Info("修改成功")
	return nil
}
