package core

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/DraftTin/gin-blog/config"
	"github.com/DraftTin/gin-blog/global"
	"gopkg.in/yaml.v2"
)

// InitConf: read the configuration from yaml file
func InitConf() {
	const ConfigFile = "settings.yaml"
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("yaml file error: %s", err))
	}
	c := &config.Config{}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yaml file loaded successfully.")
	global.Config = c
}
