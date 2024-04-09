package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	System System `yaml:"server"`
}

var Conf Config

func init() {
	yamlFile, err := os.ReadFile("./config/settings.yaml")
	if err != nil {
		fmt.Printf("[ERROR] read settings file failed, err:%v\n", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		fmt.Printf("[ERROR] unmarshal settings file failed, err:%v\n", err)
		panic(err)
	}

	fmt.Printf("[INFO] read settings file success, conf:%#v\n", Conf)
}
