package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type server struct {
	Server1 string `yaml:"server1"`
	Server2 string `yaml:"server2"`
}

type config struct {
	Port   int    `yaml:"port"`
	Code   string `yaml:"code"`
	Server server `yaml:"server"`
}

func getConfig(filePath string) {
	config := config{}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("config yaml err:%v", err)
	}

	fmt.Println(string(content))

	fmt.Printf("init data: %v\n", config)

	if yaml.Unmarshal(content, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v\n", err)
	}
	fmt.Printf("File config: %v\n", config)

	fmt.Printf("port is %v\n", config.Port)
	fmt.Printf("server is %v\n", config.Server)
	fmt.Printf("server1 is %v\n", config.Server.Server1)

}

func main() {
	getConfig("./config.yaml")
}
